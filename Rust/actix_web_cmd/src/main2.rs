use actix_web::{web, App, HttpResponse, HttpServer, Result};
use serde::{Deserialize, Serialize};
use std::collections::HashMap;
use std::fs;
use std::sync::{Arc, Mutex};
use tracing::{info, Level};

#[derive(Debug, Serialize, Deserialize, Clone)]
struct Config {
    listen: String,
    options: HashMap<String, HashMap<String, String>>,
}

async fn exec_cmd(
    info: web::Query<HashMap<String, String>>,
    data: web::Data<Arc<Mutex<Config>>>,
) -> Result<HttpResponse> {
    let group = info.get("group").unwrap_or(&String::new()).to_owned();
    let opt = info.get("opt").unwrap_or(&String::new()).to_owned();

    if group.is_empty() || opt.is_empty() {
        return Ok(HttpResponse::Ok().json("cmd error param check!"));
    }

    info!("Received request with group: {} and opt: {}", &group, &opt);

    let config = data.lock().unwrap(); // Get the lock to access the configuration

    if let Some(cmd) = config
        .options
        .get(&group)
        .and_then(|groups| groups.get(&opt))
    {
        if !cmd.is_empty() {
            #[cfg(target_os = "windows")]
            {
                let child = std::process::Command::new("cmd")
                    .arg("/C")
                    .arg(format!("{} {}", cmd, opt))
                    .spawn();
                match child {
                    Ok(_) => {
                        info!("success group: {} and opt: {}", group, opt);
                        return Ok(HttpResponse::Ok().json("cmd success"));
                    }
                    Err(e) => {
                        info!("failed group: {} and opt: {}, err: {}", group, opt, e);
                        return Ok(HttpResponse::Ok().json("cmd error param check!"));
                    }
                }
            }

            #[cfg(not(target_os = "windows"))]
            {
                let mut child_cmd = Command::new("sh");
                child_cmd.arg("-c").arg(format!("{} {}", cmd, opt));

                let child = child_cmd.spawn();
                match child {
                    Ok(_) => {
                        info!("success group: {} and opt: {}", group, opt);
                        Ok(HttpResponse::Ok().json("cmd success"))
                    }
                    Err(e) => {
                        info!("failed group: {} and opt: {}, err: {}", group, opt, e);
                        Ok(HttpResponse::Ok().json("cmd error param check!"))
                    }
                }
            }
        }
    }

    Ok(HttpResponse::Ok().json("cmd error param check!"))
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    // Initialize the logging system
    tracing_subscriber::fmt()
        .with_max_level(Level::DEBUG)
        .init();

    let config_str = fs::read_to_string("./config.json").expect("Failed to read config file");
    let config: Config = serde_json::from_str(&config_str).expect("Failed to parse config");

    let config_data = web::Data::new(Arc::new(Mutex::new(config.clone())));

    println!("-----------------------------------------------");
    println!("listen: {}", config.listen);
    for (group, opts) in &config.options {
        println!("{}", group);
        for (opt, sh) in opts {
            println!("\t{}:{}", opt, sh);
        }
    }
    println!("-----------------------------------------------");

    let host = format!("0.0.0.0:{}", config.listen);
    HttpServer::new(move || {
        App::new()
            .app_data(config_data.clone())
            .service(web::resource("/cmd").route(web::get().to(exec_cmd)))
    })
    .bind(host)?
    .run()
    .await
}
