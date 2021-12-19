use std::fs;
use std::ops::Index;
use std::process::exit;
use calamine::{Reader, open_workbook, Xlsx};

fn main() {
    if std::env::args().len() < 2 {
        println!("要输入一个文件呀大哥");
        exit(1);
    }
    let path = std::env::args().nth(1).expect("要输入一个文件呀大哥");
    run(path);
}

fn run(path: String) {
    let mut workbook: Xlsx<_> = open_workbook(path).expect("文件路径不对吧大哥!");

    let s_str = workbook.sheet_names()[0].to_string();
    let s: Vec<_> = s_str.split("|").collect();
    if s.len() < 2 {
        panic!("sheet name must \"chinese|english\"")
    }
    let out_xml = format!("服务器导出配置/{}DB.xml", s[1]);
    let prefix = s[1];
    println!("outFile:{}", out_xml);

    if let Some(Ok(r)) = workbook.worksheet_range_at(0) {
        let height = r.height();
        let width = r.width();
        println!("[{},{}]", height, width);

        let (fields, register) =
            (
                r.index(2),
                r.index(3),
            );
        //coll_server_idx(register);
        println!("{:?}\n{:?}\n{}", fields, register, fields.len() == register.len());
        let mut server_index = Vec::new();
        for i in 0..register.len() {
            let temp = register[i].to_string();
            if temp.contains("server") {
                server_index.push(i);
            }
        }

        println!("serverIndex{:?}", server_index.to_vec());
        let mut write_cent = String::from("");
        for idx in 4..height {
            let cent = r.index(idx).to_vec();
            //println!("{:?}", cent.to_vec());
            //let centVec = cent.to_vec();

            let mut write_line = String::from("");
            for i in server_index.to_vec() {
                let c_name = fields[i].to_string();
                let c_val = cent[i].to_string();
                //println!("{} -> {}", c_name, c_val);
                write_line = format!("{} {}=\"{}\"", write_line, c_name, c_val);
            }
            let w_line = format!("<{}{}/>", prefix, write_line);

            if idx == 4 {
                write_cent = w_line;
            } else {
                write_cent = format!("{}\n{}", write_cent, w_line);
            }
            //println!("w_line -> :{}", write_line);
        }

        //println!("w_cent-> :{}", write_cent);

        let write_cent = format!(
            "<?xml version=\"1.0\" encoding=\"utf-8\"?>
<{0}DB>
<{0}s>
{1}
</{0}s>
</{0}DB>"
            , prefix, write_cent);

        //println!("{}", write_cent);
        fs::write(out_xml, write_cent).unwrap();
    }
}