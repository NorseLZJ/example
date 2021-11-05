//#![deny(warnings)]
#![warn(rust_2018_idioms)]

//use futures_util::TryStreamExt;
use hyper::service::{make_service_fn, service_fn};
use hyper::{Body, Method, Request, Response, Server, StatusCode};

use std::collections::HashMap;
//use std::net::{SocketAddr, SocketAddrV4};
use url::form_urlencoded;
use accept_cmd::*;
use std::process::Command;
//use proc_macro::TokenTree::Group;

static MISSING: &[u8] = b"Missing field";

#[macro_use]
extern crate lazy_static;
lazy_static! {
    static ref CONFIG: Config = Config::new("config.json");
}

/// This is our service handler. It receives a Request, routes on its
/// path, and returns a Future of a Response.
async fn echo(req: Request<Body>) -> Result<Response<Body>, hyper::Error> {
    let query = if let Some(q) = req.uri().query() {
        q
    } else {
        return Ok(Response::builder()
            .status(StatusCode::UNPROCESSABLE_ENTITY)
            .body(MISSING.into())
            .unwrap());
    };

    match (req.method(), req.uri().path()) {
        // Serve some instructions at /
        (&Method::GET, "/cmd") => {
            let params = form_urlencoded::parse(query.as_bytes())
                .into_owned()
                .collect::<HashMap<String, String>>();

            let (group, opt) = if let (Some(group), Some(opt)) = (params.get("group"), params.get("opt")) {
                (group, opt)
            } else {
                return Ok(Response::new(Body::from("params not enough!")));
            };

            println!("group:{},opt:{}", group, opt);
            let g = if let Some(g) = CONFIG.options.get(group) {
                g
            } else {
                return Ok(Response::new(Body::from("params group not exist!")));
            };

            //println!("start:{},stop{}", g.start, g.stop);
            if opt == "start" {
                Command::new("cmd").args(&["/C", &g.start])
                    .output()
                    .expect("start exec err -->>");
            } else if opt == "stop" {
                Command::new("cmd").args(["/C", &g.stop])
                    .output()
                    .expect("stop exec err -->>");
            } else {
                return Ok(Response::new(Body::from("opt not exist!")));
            }

            return Ok(Response::new(Body::from("cmd success")));
        }

        // Return the 404 Not Found for other routes.
        _ => {
            let mut not_found = Response::default();
            *not_found.status_mut() = StatusCode::NOT_FOUND;
            Ok(not_found)
        }
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error + Send + Sync>> {
    let addr = ([0, 0, 0, 0], CONFIG.listen).into();

    let service = make_service_fn(|_| async {
        Ok::<_, hyper::Error>(service_fn(echo))
    });

    let server = Server::bind(&addr).serve(service);

    println!("Listening on http://{}", addr);

    server.await?;

    Ok(())
}