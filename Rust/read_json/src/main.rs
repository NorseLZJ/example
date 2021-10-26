use std::fs;
use std::path;

mod config;

fn main()
{
    let cfg: config::Config = config::new_config("");
    let data: config::Data = config::new_data("E:\\lzj\\Rust\\hello\\src\\data.json");

    for i in data.frames {
        println!("{}{:?}", i.0.as_str(), i.1.frame)
    }
}
