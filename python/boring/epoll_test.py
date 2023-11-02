import concurrent.futures
import socket
import random
import string
import time


MAX_WORKS = 10


def generate_random_string(length):
    letters = string.ascii_letters + string.digits
    return "".join(random.choice(letters) for _ in range(length))


def send_data(client_socket: socket.socket):
    count = 0
    while True:
        message = generate_random_string(1000)
        client_socket.send(message.encode())

        response = client_socket.recv(1024)
        # print("Received response:", response.decode())

        time.sleep(0.3)  # 睡眠300毫秒
        count += 1
        if count >= 100:
            break

    client_socket.close()


def main():
    server_address = ("172.28.241.123", 8080)

    with concurrent.futures.ThreadPoolExecutor(max_workers=MAX_WORKS) as executor:
        for _ in range(MAX_WORKS):
            client_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
            client_socket.connect(server_address)
            executor.submit(send_data, client_socket)


if __name__ == "__main__":
    main()
