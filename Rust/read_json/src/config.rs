use std::collections::HashMap;
use std::fs;

use serde::{Serialize, Deserialize};

// rep config
#[derive(Deserialize, Debug)]
pub struct Config {
    frames: Vec<String>,
    outdir: String,
}

pub fn new_config(path: &str) -> Config {
    Config {
        frames: vec![],
        outdir: String::from(""),
    }
}

pub fn new_data(path: &str) -> Data {
    let text = fs::read_to_string(&path).unwrap();
    let data: Data = serde_json::from_str(&text).unwrap();
    return data;
}

// data file struct
#[derive(Deserialize, Debug)]
pub struct Data {
    pub frames: HashMap<String, Bmp>,
    pub meta: Meta,
}

#[derive(Deserialize, Debug)]
pub struct Meta {
    pub image: String,
    pub prefix: String,
}

#[derive(Deserialize, Debug)]
pub struct Bmp {
    pub frame: Frame,
    pub sourceSize: SourceSize,
    pub spriteSourceSize: SpriteSourceSize,
}

#[derive(Deserialize, Debug)]
pub struct Frame {
    pub h: i32,
    pub idx: i32,
    pub w: i32,
    pub x: i32,
    pub y: i32,
}

#[derive(Deserialize, Debug)]
pub struct SourceSize {
    pub h: i32,
    pub w: i32,
}

#[derive(Deserialize, Debug)]
pub struct SpriteSourceSize {
    pub x: i32,
    pub y: i32,
}

fn decode_data(data: &Data) {}