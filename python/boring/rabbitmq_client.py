import pika
import signal
import sys
import threading
import concurrent.futures

MAX_WORKS = 4


# 定义回调函数来处理接收到的消息
def callback(ch, method, properties, body):
    thread_id = threading.get_ident()
    print(
        "Thread ID: %s, Received message from queue '%s': %r"
        % (thread_id, method.routing_key, body)
    )


# 定义信号处理函数
def signal_handler(sig, frame):
    print("Exiting...")
    sys.exit(0)


def run_task(index: int):
    # 连接 RabbitMQ 服务器
    connection = pika.BlockingConnection(
        pika.ConnectionParameters(
            host="172.28.241.123",
            port=5672,
            virtual_host="/123pan",
            credentials=pika.PlainCredentials("123pan", "123"),
        )
    )

    ch = connection.channel()
    ch.queue_declare(queue="queue1")
    ch.basic_consume(queue="queue1", on_message_callback=callback, auto_ack=True)
    print("Waiting for messages. To exit press Ctrl+C", index)
    ch.start_consuming()


def main():
    with concurrent.futures.ThreadPoolExecutor(max_workers=MAX_WORKS) as executor:
        for index in range(MAX_WORKS):
            executor.submit(run_task, index)

        signal.signal(signal.SIGINT, signal_handler)
        executor.shutdown()


if __name__ == "__main__":
    main()
