use std::path::Path;
use std::{env, fs};

use protobuf_codegen::Codegen;

fn main() {
    let args: Vec<String> = env::args().collect();
    if args.iter().any(|arg| arg == "gen") {
        generate_protos();
        if let Err(err) = convert_camelcase_to_snakecase("robot/pb/") {
            eprintln!("Error: {}", err);
        }
    }
}

fn find_proto_files(dir: &Path, suffix: &str) -> Result<Vec<String>, std::io::Error> {
    let file_paths: Vec<String> = fs::read_dir(dir)?
        .filter_map(|entry| {
            entry.ok().and_then(|e| {
                if e.path().is_file() && e.path().extension() == Some(suffix.as_ref()) {
                    Some(e.path().to_string_lossy().into_owned())
                } else {
                    None
                }
            })
        })
        .collect();

    Ok(file_paths)
}

fn camelcase_to_snakecase(input: &str) -> String {
    let mut result = String::new();
    let mut last_char_was_uppercase = false;

    for c in input.chars() {
        if c.is_uppercase() {
            if last_char_was_uppercase {
                result.push('_');
            }
            result.push(c.to_lowercase().next().unwrap());
            last_char_was_uppercase = true;
        } else {
            result.push(c);
            last_char_was_uppercase = false;
        }
    }

    result
}

fn convert_camelcase_to_snakecase(dir_path: &str) -> Result<(), std::io::Error> {
    let entries = fs::read_dir(dir_path)?;

    for entry in entries {
        if let Ok(entry) = entry {
            let file_path = entry.path();
            if file_path.is_file() {
                if let Some(file_name) = file_path.file_name() {
                    let original_name = file_name.to_string_lossy();
                    let new_name = camelcase_to_snakecase(&original_name);
                    let new_path = file_path.with_file_name(new_name);
                    fs::rename(&file_path, &new_path)?;
                    //println!("Renamed: {} -> {}", original_name, new_name);
                }
            }
        }
    }

    Ok(())
}

fn generate_protos() {
    let cur_dir = env::current_dir().unwrap();
    let src_dir = cur_dir.join("src").join("protos");
    let pb_dir = cur_dir.join("robot").join("pb");
    println!("{:?}", &pb_dir);
    let files = find_proto_files(&src_dir, "proto").unwrap();
    Codegen::new()
        .pure()
        .out_dir(pb_dir.clone())
        .inputs(files)
        .include(src_dir)
        .run_from_script();
}

// let mut msg = pb::Login::LoginReq::new();
// let mut lpd = pb::UserData::LoginPlatformData::new();
// lpd.uid = 1111;
// lpd.acount = String::from("Norse");
// msg.login_platform_data = MessageField::from_option(Some(lpd));
// println!("{}", msg);
// let out_bytes: Vec<u8> = msg.write_to_bytes().unwrap();
// println!("{:?}", out_bytes);
// println!("{}", out_bytes.len());
