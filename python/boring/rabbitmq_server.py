import time
import pika
import string
import random
import sys
import signal


"""
启动rabbit-mq 的docker命令
    docker run -d --name rabbitmq1 -p 5672:5672 -p 15672:15672 -e RABBITMQ_DEFAULT_VHOST=my_vhost -e RABBITMQ_DEFAULT_USER=root -e RABBITMQ_DEFAULT_PASS=123456 rabbitmq:3-management
"""
# 连接 RabbitMQ 服务器
conn = pika.BlockingConnection(
    pika.ConnectionParameters(
        host="172.28.241.123",
        port=5672,
        virtual_host="/123pan",
        credentials=pika.PlainCredentials("123pan", "123"),
    )
)

channel = conn.channel()

# 创建一个名为 "hello" 的队列
channel.queue_declare(queue="queue1")


def generate_random_string(length):
    letters = string.ascii_letters + string.digits
    return "".join(random.choice(letters) for _ in range(length))


# 发送消息到队列
def send_message():
    while True:
        channel.basic_publish(
            exchange="", routing_key="queue1", body=generate_random_string(10)
        )
        time.sleep(1)


# 定义信号处理函数
def signal_handler(signal, frame):
    print("Exiting...")
    conn.close()
    sys.exit(0)


# 注册信号处理函数
signal.signal(signal.SIGINT, signal_handler)

# 启动发送消息的函数
send_message()
