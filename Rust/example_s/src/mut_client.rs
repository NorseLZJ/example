use std::{
    net::{Ipv4Addr, SocketAddr, UdpSocket},
    process::exit,
};
use tokio::time::{sleep, Duration};

#[tokio::main]
async fn main() {
    let multicast_addr: Ipv4Addr = "239.0.0.0".parse().unwrap();
    let multicast_port: u16 = 6000;
    let multicast_socket_addr = SocketAddr::new(multicast_addr.into(), multicast_port);

    let client_socket = UdpSocket::bind("0.0.0.0:0").unwrap();
    if let Err(err) = client_socket.join_multicast_v4(&multicast_addr, &"0.0.0.0".parse().unwrap())
    {
        eprintln!("Failed to join multicast group: {}", err);
        exit(1);
        // 处理错误的逻辑
    }

    println!(
        "Client joined the multicast group {:?}",
        multicast_socket_addr
    );

    loop {
        tokio::select! {
            _ = send_data(&client_socket, &multicast_socket_addr) => {
                // Sent data
            }
            result = receive_data(&client_socket) => {
                if let Ok(received_data) = result {
                    println!("Received data: {}", received_data);
                }
            }
        }

        // Sleep for 5 seconds
        sleep(Duration::from_secs(5)).await;
    }
}

async fn send_data(socket: &UdpSocket, addr: &SocketAddr) {
    let data = "Hello from client!";
    println!("send data....");
    socket.send_to(data.as_bytes(), addr).unwrap();
}

async fn receive_data(socket: &UdpSocket) -> Result<String, std::io::Error> {
    let mut buf = [0u8; 1024];
    let (size, _) = socket.recv_from(&mut buf).expect("sss");
    let received_data = String::from_utf8_lossy(&buf[..size]).to_string();
    Ok(received_data)
}
