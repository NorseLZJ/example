use serde::{Deserialize, Serialize};
use std::collections::HashMap;
use std::fs;

#[derive(Serialize, Deserialize, Debug)]
struct Config {
    listen: String,
    options: HashMap<String, HashMap<String, String>>,
}

fn main() {
    let config_str = fs::read_to_string("./config.json").expect("Failed to read config file");
    let config: Config = serde_json::from_str(&config_str).expect("Failed to parse config");

    println!("{:?}", config.listen);

    let cmd = config
        .options
        .get("dopai")
        .and_then(|opts| opts.get("start"))
        .map(|cmd| cmd.as_str());

    if let Some(cmd) = cmd {
        println!("{}", cmd);
    }
}
