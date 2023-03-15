"""
check share link

使用线程池处理问题的一个办法 
"""
import requests
import json
import concurrent.futures
import os


def headers():
    ret = {
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36",
    }
    return ret


class ShareLink(object):
    def __init__(self, share_key: str, pwd: str):
        # print(share_key)
        self.share_key = share_key
        self.pwd = pwd
        self.run(0)

    def run(self, parent_file_id: int):
        url = "https://www.123pan.com/a/api/share/get"
        data = {
            "limit": 100,
            "next": 1,
            "orderBy": "share_id",
            "orderDirection": "desc",
            "shareKey": self.share_key,
            "SharePwd": self.pwd,
            "ParentFileId": parent_file_id,
            "Page": 1,
        }
        resp = requests.get(url, headers=headers(), data=data)
        if resp.status_code == 200:
            # print(resp.text)
            jsoner = json.loads(resp.text)
            infolist = jsoner["data"]["InfoList"]
            for v in infolist:
                if v["Type"] == 1:
                    self.run(v["FileId"])
                else:
                    # print("%-50s%-30s" % (v["FileName"], v["Etag"]))
                    pass


def create_share_link(path: str):
    context = ""
    with open(path, "r") as f:
        context = f.read()
    jsoner = json.loads(context)
    for s in jsoner:
        ShareLink(share_key=s["share_key"], pwd=s["share_pwd"])
    return True


def create_share_link2(v):
    print(os.getpid(), "-", v["share_key"])
    if v["share_key"] == "":
        return False
    ShareLink(share_key=v["share_key"], pwd=v["share_pwd"])
    return True


if __name__ == "__main__":
    """
    files = os.listdir("data")
    fs = []
    for f in files:
        if f.endswith("json"):
            fs.append(format("%s\data\%s" % (os.getcwd(), f)))

    # create_share_link(fs[0])
    with concurrent.futures.ProcessPoolExecutor(max_workers=4) as executor:
        for number, prime in zip(fs, executor.map(create_share_link, fs)):
            print("%d is prime: %s" % (number, prime))
    """
    ss = [
        {"share_key": "a09KVv-IYQD3", "share_pwd": ""},
        {"share_key": "a09KVv-MYQD3", "share_pwd": ""},
        {"share_key": "a09KVv-wYQD3", "share_pwd": ""},
        {"share_key": "a09KVv-FYQD3", "share_pwd": "HeiY"},
        {"share_key": "a09KVv-YYQD3", "share_pwd": "wHdi"},
        {"share_key": "", "share_pwd": ""},
    ]
    with concurrent.futures.ProcessPoolExecutor(max_workers=4) as executor:
        executor.map(create_share_link2, ss)
        # print(number, "-", prime)
