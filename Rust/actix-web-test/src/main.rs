use actix_web::{middleware, web, App, HttpRequest, HttpServer, Responder};
use serde_derive::Deserialize;
use std::fs;

#[derive(Deserialize)]
pub struct ReqParam {
    group: String,
}

fn dir_exist(group: &str) -> (bool, String, usize) {
    let c_path = std::env::current_dir().unwrap();
    for entry in fs::read_dir(c_path).unwrap() {
        let entry = entry.unwrap();
        let path = entry.path().display().to_string();
        let md = std::fs::metadata(&path).unwrap();
        if md.is_file() {
            continue;
        }

        let idx = path.find(group);
        match idx {
            Some(idx) => {
                println!("path:{} idx: {}", path, idx);
                return (true, path.clone(), idx);
            }
            None => {
                //println!("None...");
            }
        }
    }
    (false, String::from(""), 0)
}

async fn index(req: HttpRequest) -> &'static str {
    println!("REQ: {:?}", req);
    "Hello world!"
}

async fn stop(web::Query(info): web::Query<ReqParam>) -> impl Responder {
    let (ret, path, _) = dir_exist(&info.group);
    if ret {
        return format!("Stop Success {}!", path);
    }
    format!("Unknow group {}!", info.group)
}

async fn start(web::Query(info): web::Query<ReqParam>) -> impl Responder {
    format!("Start Success {}!", info.group)
}

async fn update(web::Query(info): web::Query<ReqParam>) -> impl Responder {
    format!("Update Success {}!", info.group)
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    std::env::set_var("RUST_LOG", "actix_web=info");
    env_logger::init();

    HttpServer::new(|| {
        App::new()
            // enable logger
            .wrap(middleware::Logger::default())
            .service(web::resource("/index.html").to(|| async { "Hello world!" }))
            .service(web::resource("/").to(index))
            .service(web::resource("/stop").to(stop))
            .service(web::resource("/start").to(start))
            .service(web::resource("/update").to(update))
    })
    .bind("127.0.0.1:8080")?
    .run()
    .await
}
