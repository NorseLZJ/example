use protobuf::{Enum, Message, MessageField};
use std::convert::TryInto;
use std::io::{Read, Write};
use std::net::TcpStream;
use std::sync::atomic::{AtomicU32, Ordering};
use std::sync::{Arc, Mutex};
use std::thread::*;
use std::thread::{self, sleep};
use std::time::Duration;

use crate::pb;

const MAGIC: u16 = 0xC738;
const VERSION: u16 = 1;
const CUSTOM_HEADER_LENGTH: u32 = 32;
const CKSUM_ENCRYPT_KEY: u32 = 0xc9bea7e5;
const CONTENT_ENCRYPT_KEY: u8 = 0xBE;

// #[derive(Debug, Clone, Copy)]
pub struct Robot {
    uid: u64,
    phone_number: String,
    conn: Arc<Mutex<Option<TcpStream>>>,
    seq: Arc<AtomicU32>,
}

struct ProtocolHead {
    body_length: u32,
    magic: u16,
    version: u16,
    uid: u64,
    cmd_id: u32,
    seq: u32,
    cksum: u32,
    flag: u32,
}

struct MsgPacket {
    head: ProtocolHead,
}

impl Robot {
    pub fn new(uid: u64, phone_number: &str) -> Self {
        Robot {
            uid,
            phone_number: phone_number.to_string(),
            conn: Arc::new(Mutex::new(None)),
            seq: Arc::new(AtomicU32::from(1)),
            // req_index: 0,
        }
    }

    pub fn set_conn(&self, conn: TcpStream) {
        let mut conn_lock = self.conn.lock().unwrap();
        *conn_lock = Some(conn);
    }

    fn retry(host: &str) -> Option<TcpStream> {
        match TcpStream::connect(host) {
            Ok(conn) => Some(conn),
            Err(_) => None,
        }
    }

    pub fn init_connect(&mut self, server_flag: usize) {
        let server_hosts = vec![
            "172.23.176.238:9000",  // wsl
            "127.0.0.1:19003",      // NorseLZJ
        ];

        let host = match server_hosts.get(server_flag) {
            Some(&h) => h,
            None => panic!("Server not configured"),
        };
        println!("{}", host);

        let mut count = 0;
        let mut flag = false;
        while count <= 10 {
            if let Some(conn) = Robot::retry(host) {
                self.set_conn(conn);
                println!("连上了");
                flag = true;
                break;
            }
            sleep(Duration::from_secs(1));
            println!(
                "num:{} {} {} retry {}",
                count, self.uid, self.phone_number, host
            );
            count += 1;
        }

        if !flag {
            println!(
                "{} {} Failed to connect to the server {}",
                self.uid, self.phone_number, host
            );
            return;
        }

        //let _ = thread::spawn(move || Robot::read_msg(&self));
        //let self_clone = self.clone(); // 克隆 self，使闭包拥有自己的拷贝
        //let _ = thread::spawn(move || self_clone.read_msg());

        //self.read_msg();
        let mut msg = pb::Login::LoginReq::new();
        let mut lpd = pb::UserData::LoginPlatformData::new();
        lpd.uid = self.uid;
        lpd.acount = self.phone_number.clone();
        msg.login_platform_data = MessageField::from_option(Some(lpd));
        let out_bytes: Vec<u8> = msg.write_to_bytes().unwrap();
        self.uid = 0;
        self.send_msg(out_bytes.clone(), pb::CmdId::CmdId::C_Login);
        println!("发送登录消息了");
        loop {
            sleep(Duration::from_secs(1));
            if self.uid != 0 {
                break;
            }
        }
        println!("登录成功了 ,然后阻塞住");

        loop {
            sleep(Duration::from_secs(1));
        }

        // loop {
        //     if self.req_index < REQ_LIST.len() && self.uid != 0 {
        //         let v = &REQ_LIST[self.req_index];
        //         self.send_msg(&v.msg, v.cmd_id);
        //         self.req_index += 1;
        //     } else {
        //         self.req_index = 0;
        //     }
        //     thread::sleep(Duration::from_secs(5));
        // }
    }

    fn read_msg(&self) {
        loop {
            println!("循环读消息");
            let mut head_buff = vec![0; CUSTOM_HEADER_LENGTH as usize];

            let mut conn_lock = self.conn.lock().unwrap();
            // 读head
            if let Some(ref mut conn) = *conn_lock {
                let _ = conn.read_exact(&mut head_buff);
            } else {
                println!("连接错误，需要断开");
                break;
            }

            if let Ok(head) = self.parse_message_header(&head_buff) {
                let mut buff = vec![0; head.head.body_length as usize];
                // 读body
                if let Some(ref mut conn) = *conn_lock {
                    let _ = conn.read_exact(&mut buff);
                } else {
                    println!("连接错误，需要断开");
                    break;
                }
                let buff = self.encrypt_decrypt_content(&buff);
                if !buff.is_empty() {
                    if self.calc_cksum(&buff) != head.head.cksum {
                        continue;
                    }
                }
                match pb::CmdId::CmdId::from_i32(head.head.cmd_id as i32) {
                    Some(cmdid) => {
                        if cmdid == pb::CmdId::CmdId::C_Login {
                            println!("登录成功特殊处理")
                        } else {
                            println!("其它消息，打印下就行")
                        }
                    }
                    None => println!("None"),
                }
            } else {
                continue;
            }
        }
        eprintln!("user exit: {}", self.uid);
    }

    fn calc_cksum(&self, data: &[u8]) -> u32 {
        let mut cksum = 0;

        for i in (0..data.len()).step_by(4) {
            let end = std::cmp::min(i + 4, data.len());
            let word = &data[i..end];

            for (k, &b) in word.iter().enumerate() {
                cksum += (b as u32) << (8 * k as u32);
            }
        }
        !cksum ^ CKSUM_ENCRYPT_KEY
    }

    fn encrypt_decrypt_content(&self, data: &[u8]) -> Vec<u8> {
        data.iter().map(|&b| b ^ CONTENT_ENCRYPT_KEY).collect()
    }

    fn send_msg(&self, msg: Vec<u8>, cmd_id: pb::CmdId::CmdId) {
        let data_len = msg.len() as u32;
        let mut buffer = vec![0; (CUSTOM_HEADER_LENGTH + data_len) as usize];

        let cmdid = cmd_id as u32;
        buffer[0..4].copy_from_slice(&data_len.to_be_bytes());
        buffer[4..6].copy_from_slice(&MAGIC.to_be_bytes());
        buffer[6..8].copy_from_slice(&VERSION.to_be_bytes());
        buffer[8..16].copy_from_slice(&self.uid.to_be_bytes());
        buffer[16..20].copy_from_slice(&cmdid.to_be_bytes());
        buffer[20..24].copy_from_slice(&self.seq.fetch_add(1, Ordering::SeqCst).to_be_bytes());
        buffer[24..28].copy_from_slice(&self.calc_cksum(&msg).to_be_bytes());
        buffer[28..32].copy_from_slice(&0u32.to_be_bytes());

        let encrypted_data = self.encrypt_decrypt_content(&msg);
        buffer[CUSTOM_HEADER_LENGTH as usize..].copy_from_slice(&encrypted_data);

        let mut conn_lock = self.conn.lock().unwrap();
        if let Some(ref mut conn) = *conn_lock {
            let _ = conn.write_all(&buffer);
        } else {
            eprintln!("不存在的连接, 需要断开");
        }
    }

    fn parse_message_header(&self, head_buffer: &[u8]) -> Result<MsgPacket> {
        let body_length = u32::from_be_bytes(head_buffer[0..4].try_into().unwrap());
        let magic = u16::from_be_bytes(head_buffer[4..6].try_into().unwrap());
        let version = u16::from_be_bytes(head_buffer[6..8].try_into().unwrap());
        let uid = u64::from_be_bytes(head_buffer[8..16].try_into().unwrap());
        let cmd_id = u32::from_be_bytes(head_buffer[16..20].try_into().unwrap());
        let seq = u32::from_be_bytes(head_buffer[20..24].try_into().unwrap());
        let cksum = u32::from_be_bytes(head_buffer[24..28].try_into().unwrap());
        Ok(MsgPacket {
            head: ProtocolHead {
                body_length,
                magic,
                version,
                uid,
                cmd_id,
                seq,
                cksum,
                flag: 0,
            },
        })
    }
}
