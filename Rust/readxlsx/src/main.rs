use readxlsx::static_variable::*;
use std::process::exit;

fn main() {
    if std::env::args().len() < 2 {
        println!("要输入一个文件呀大哥");
        exit(1);
    }
    let path = std::env::args().nth(1).expect("要输入一个文件呀大哥");

    let mut cc = CppConfig::new(&path);
    cc.run();
}
