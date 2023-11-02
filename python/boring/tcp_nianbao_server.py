import socket
import struct


def receive_data(conn):
    # 接收长度信息
    length_data = conn.recv(4)
    if len(length_data) <= 0:
        return None
    length = struct.unpack("!I", length_data)[0]

    # 接收数据
    received_data = b""
    while len(received_data) < length:
        chunk = conn.recv(length - len(received_data))
        if not chunk:
            break
        received_data += chunk

    return received_data.decode()


def main():
    # 创建TCP套接字
    server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server_socket.bind(("localhost", 8888))
    server_socket.listen(1)

    print("Server is listening...")

    # 接受客户端连接
    client_socket, client_address = server_socket.accept()
    print("Connected with client:", client_address)

    # 接收多个数据包
    while True:
        received_data = receive_data(client_socket)
        if not received_data:
            break
        print("Received data:", received_data)

    # 关闭套接字
    client_socket.close()
    server_socket.close()


if __name__ == "__main__":
    main()
