use std::collections::{HashMap};
use std::fs;

use serde::{Deserialize};
use serde::__private::Option;

#[derive(Deserialize, Debug)]
pub struct Group {
    pub start: String,
    pub stop: String,
}

#[derive(Deserialize, Debug)]
pub struct Config {
    pub listen: u16,
    pub options: HashMap<String, Group>,
}

impl Config {
    pub fn new(path: &str) -> Config {
        let text = fs::read_to_string(&path).unwrap();
        let config: Config = serde_json::from_str(&text).unwrap();
        return config;
    }
}

