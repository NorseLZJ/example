import socket
import struct


def send_data(sock, data):
    # 计算数据长度
    length = len(data)
    length_data = struct.pack("!I", length)

    # 发送长度信息
    sock.sendall(length_data)

    # 发送数据
    sock.sendall(data.encode())


def main():
    # 创建TCP套接字
    client_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

    # 连接服务器
    client_socket.connect(("localhost", 8888))
    print("Connected to server")

    # 发送多个小数据包
    send_data(client_socket, "Hello")
    send_data(client_socket, ",")
    send_data(client_socket, "world!")

    # 关闭套接字
    client_socket.close()


if __name__ == "__main__":
    main()
