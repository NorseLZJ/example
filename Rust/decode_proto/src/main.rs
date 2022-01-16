use std::fmt::format;
use std::fs;
use std::ops::Index;
use std::process::exit;
use regex::Regex;


fn main() {
    let text = fs::read_to_string("netcmd.proto").unwrap();

    //let text = text.replace("\n", "");
    //println!("{}", text);

    let re = Regex::new(r"(?s)message(.*?)}").unwrap();
    let refs = Regex::new(r"(?s)[uirs](.*?);").unwrap();

    for cap in re.captures_iter(&text) {
        let mut cmsg = &cap[0];
        let cmsg = cmsg.replace("\n\r", "").replace("\n", "").replace("\n\t", "").replace("\t", "");
        let tmsg = &cmsg;
        tmsg.replace(" ", "");

        let (idx1, idx2) = if let (Some(idx1), Some(idx2)) = (tmsg.find("{"), tmsg.find("}")) {
            (idx1, idx2)
        } else {
            (0, 0)
        };

        if idx2 - idx1 <= 2 {
            continue;
        }

        let mut cmsg = cmsg.replace("message", "");
        if let Some(idx1) = cmsg.find("{") {
            let msg_name = cmsg.split_off(idx1);
            println!("{:?}", msg_name);
        } else {};


        //println!("{:?}", cmsg);
        //let fes:Vec<_>=cmsg.split("\n").collect();
    }

    exit(0);
}
