use std::env;
use std::net::SocketAddr;

use futures::StreamExt;
use tokio::io;
use tokio_util::codec::{BytesCodec, FramedRead, FramedWrite};

#[tokio::main]
async fn main() {
    let mut args = env::args().skip(1).collect::<Vec<_>>();
    let tcp = match args.iter().position(|a| a == "--udp") {
        Some(i) => {
            args.remove(i);
            false
        }
        None => true,
    };
    println!("is_tcp: {}", tcp);
    println!("{:?}", args);

    let d_addr = "172.23.176.238:7711".to_string(); // 设置默认地址
    let addr = args.first().map_or_else(|| d_addr.clone(), |a| a.clone());
    println!("{}", addr);
    let addr = addr.parse::<SocketAddr>().unwrap(); // 这里不再使用?，直接unwrap
    println!("{}", addr);

    let stdin = FramedRead::new(io::stdin(), BytesCodec::new());
    let stdin = stdin.map(|i| i.map(|bytes| bytes.freeze()));
    let stdout = FramedWrite::new(io::stdout(), BytesCodec::new());
    tcp::connect(&addr, stdin, stdout).await.unwrap();
}

mod tcp {
    use bytes::Bytes;
    use futures::{future, Sink, SinkExt, Stream, StreamExt};
    use std::{error::Error, io, net::SocketAddr};
    use tokio::net::TcpStream;
    use tokio_util::codec::{BytesCodec, FramedRead, FramedWrite};

    pub async fn connect(
        addr: &SocketAddr,
        mut stdin: impl Stream<Item=Result<Bytes, io::Error>> + Unpin,
        mut stdout: impl Sink<Bytes, Error=io::Error> + Unpin,
    ) -> Result<(), Box<dyn Error>> {
        let mut stream = TcpStream::connect(addr).await?;
        let (r, w) = stream.split();
        let mut sink = FramedWrite::new(w, BytesCodec::new());
        let mut stream = FramedRead::new(r, BytesCodec::new())
            .filter_map(|i| match i {
                Ok(i) => future::ready(Some(i.freeze())),
                Err(e) => {
                    println!("failed to read from socket; error={}", e);
                    future::ready(None)
                }
            })
            .map(Ok);

        match future::join(sink.send_all(&mut stdin), stdout.send_all(&mut stream)).await {
            (Err(e), _) | (_, Err(e)) => Err(e.into()),
            _ => Ok(()),
        }
    }
}
