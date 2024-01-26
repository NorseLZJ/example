use std::net::{Ipv4Addr, SocketAddr, UdpSocket};
use tokio::time::Duration;

#[tokio::main]
async fn main() {
    let multicast_addr: Ipv4Addr = "239.0.0.0".parse().unwrap();
    let multicast_port: u16 = 6000;
    let multicast_socket_addr = SocketAddr::new(multicast_addr.into(), multicast_port);

    let server_socket = UdpSocket::bind("0.0.0.0:0").unwrap();
    server_socket
        .join_multicast_v4(&multicast_addr, &"0.0.0.0".parse().unwrap())
        .unwrap();

    println!("Server listening on {:?}", multicast_socket_addr);

    let mut buf = [0u8; 1024];

    loop {
        if let Ok((size, _)) = server_socket.recv_from(&mut buf) {
            let received_data = String::from_utf8_lossy(&buf[..size]);
            println!("Received data: {}", received_data);
        }
        tokio::time::sleep(Duration::from_secs(1)).await;
    }
}
