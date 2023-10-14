import yaml
import socket
import os
import sys
import concurrent.futures
import random
import ping3


def parse_and_check(domain):
    try:
        ip = socket.gethostbyname(domain)
        response = ping3.ping(ip, timeout=2)
        if response is not None:
            return f"{domain} 是通畅的 {ip}"
        else:
            return None
    except socket.gaierror as e:
        print(f"err:{e}")
        return None


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("输入文件!!!")
        exit(1)
    domains = []
    with open(sys.argv[1], encoding="UTF-8") as f:
        y = yaml.safe_load(f)
        for v in y["proxies"]:
            domains.append(v["server"])

    with concurrent.futures.ThreadPoolExecutor(max_workers=8) as executor:
        futures = [executor.submit(parse_and_check, domain) for domain in domains]
        for future, domain in zip(concurrent.futures.as_completed(futures), domains):
            ret = future.result()
            if ret:
                print(ret)
