use regex::Regex;
use std::env;
use std::fs::File;
use std::io::Write;

mod txt;
use txt::txt::src_txt;

fn usage(args: &str) {
    eprintln!(
        "Usage: {} rpc_server [map,commerce,dynamic,masterd] <MsgName>",
        args
    );
    std::process::exit(1);
}

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() < 3 {
        usage(&args[0])
    }

    let arg1 = &args[1];
    let arg2 = &args[2];
    let re = Regex::new(r"LoadMap+").unwrap();
    let mut replaced = re.replace_all(src_txt(), arg2);
    let r2 = replaced.clone();

    match arg1.as_str() {
        "map" => {}
        "commerce" => {
            let re = Regex::new(r"SERVER_SCHCEME_MAP").unwrap();
            replaced = re.replace_all(&r2, "SERVER_SCHCEME_COMMERCE");
        }
        "dynamic" => {
            let re = Regex::new(r"SERVER_SCHCEME_MAP").unwrap();
            replaced = re.replace_all(&r2, "SERVER_SCHCEME_DYNAMIC");
        }
        "masterd" => {
            let re = Regex::new(r"SERVER_SCHCEME_MAP").unwrap();
            replaced = re.replace_all(&r2, "SERVER_SCHCEME_MASTERD");
        }
        _ => usage(&args[0]),
    }

    //println!("Replaced text: {}", replaced);
    let mut file = File::create(format!("{}.go", arg2)).unwrap();
    file.write_all(replaced.as_bytes()).unwrap();
}
