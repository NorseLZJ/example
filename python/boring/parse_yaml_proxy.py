import yaml
import socket
import os
import concurrent.futures
import numpy as np


def parse_and_check(domain):
    idx = np.random.random_integers(low=1, high=99999999, size=1)[0]

    def clean_file():
        os.remove(f"{idx}.txt")

    try:
        ip = socket.gethostbyname(domain)
        response = os.system(f"ping -c 1 {ip} >{idx}.txt")
        if response == 0:
            clean_file()
            return f"{domain} 是通畅的 {ip}"
        else:
            clean_file()
            return f"{domain} 是不通畅的 {ip}"
    except socket.gaierror:
        clean_file()
        return f"{domain} 是无效的域名 {ip}"


if __name__ == "__main__":
    domains = []
    with open("1670388083880.yml", encoding="UTF-8") as f:
        y = yaml.safe_load(f)
        for v in y["proxies"]:
            domains.append(v["server"])

    with concurrent.futures.ProcessPoolExecutor(max_workers=20) as executor:
        for number, ret in zip(domains, executor.map(parse_and_check, domains)):
            print("%s" % (ret))
