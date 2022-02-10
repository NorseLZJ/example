use actix_web::{middleware, web, App, HttpRequest, HttpServer, Responder};
use serde_derive::Deserialize;

#[derive(Deserialize)]
pub struct ReqParam {
    group: String,
}

async fn index(req: HttpRequest) -> &'static str {
    println!("REQ: {:?}", req);
    "Hello world!"
}

async fn stop(web::Query(info): web::Query<ReqParam>) -> impl Responder {
    format!("Stop Success {}!", info.group)
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
