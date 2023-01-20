"""
123pan.com 登录,获取文件列表示例
"""

import requests
import json
import os
import hashlib

_url = "https://123pan.com/"
_passport = ""
_password = ""


def file_info(self, path: str) -> tuple:
    ss = path.split("\\")
    filename = ss[len(ss) - 1]
    stat = os.stat(path)
    with open(path, "rb") as fp:
        data = fp.read()
        md5 = hashlib.md5(data).hexdigest()
        return stat.st_size, filename, md5


class PanClient(object):
    def __init__(self):
        self.token = ""
        self.expire = ""
        self.nick_name = ""
        self.space_used = 0
        self.space_permanent = 0
        self.sign()
        self.user_info()
        # upload 最后好像有个签名问题，先不管了
        # self.upload_request()
        self.file_list()

    def __str__(self) -> str:
        return format("nick_name:%s\nspaceUsed:%d\nspacePermanent:%d" % (self.nick_name, self.space_used, self.space_permanent))

    def headers(self, token: bool = True):
        ret = {
            "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
            # "Host": "123pan.com:8080",
            "App-Version": "1.2",
            "platform": "web",
        }
        if token:
            ret["Authorization"] = format("Bearer %s" % (self.token))
        return ret

    def sign(self):
        url = format("%s/api/user/sign_in" % (_url))
        headers = self.headers(token=False)
        resp = requests.post(url, headers=headers, data={"passport": _passport, "password": _password, "remember": True})
        if resp.status_code == 200:
            jsoner = json.loads(resp.text)
            if jsoner["message"] == "success":
                self.expire = jsoner["data"]["expire"]
                self.token = jsoner["data"]["token"]
                print("登录成功...")
            else:
                print("登录失败")
                exit(1)

    def user_info(self):
        url = format("%s/api/user/info" % (_url))
        headers = self.headers()
        resp = requests.get(url, headers=headers)
        if resp.status_code == 200:
            jsoner = json.loads(resp.text)
            self.nick_name = jsoner["data"]["Nickname"]
            self.space_used = jsoner["data"]["SpaceUsed"]
            self.space_permanent = jsoner["data"]["SpacePermanent"]

    def upload_request(self):
        url = format("%s/api/file/upload_request" % (_url))
        headers = self.headers()
        path = "D:\\document\\alter_add_file.sql"
        size, filename, md5 = file_info(path)
        data = {
            "driveId": 0,
            "etag": md5,
            "fileName": filename,
            "parentFileId": 0,
            "size": size,
            "type": 0,
            "duplicate": 0,
        }
        resp = requests.post(url, headers=headers, data=data)
        if resp.status_code == 200:
            # print(resp.text)
            jsoner = json.loads(resp.text)
            self.s3_repare_upload_parts_batch(
                bucket=jsoner["data"]["Bucket"],
                key=jsoner["data"]["Key"],
                uploadId=jsoner["data"]["UploadId"],
            )

    def s3_repare_upload_parts_batch(self, bucket: str, key: str, uploadId: str):
        url = format("%s/api/file/s3_repare_upload_parts_batch" % (_url))
        headers = self.headers()
        data = {
            "bucket": bucket,
            "key": key,
            "partNumberEnd": 2,
            "partNumberStart": 1,
            "uploadId": uploadId,
        }
        resp = requests.post(url, headers=headers, data=data)
        if resp.status_code == 200:
            print(resp.text)
            jsoner = json.loads(resp.text)
            up_urls = jsoner["data"]["presignedUrls"]["1"]

    def cjdd_up(self, cjdd_url: str):
        pass

    def file_list(self, driveId=0, limit=1000, next=0, parentFileId=0, Page=1):
        url = format("%s/api/file/list/new" % (_url))
        headers = self.headers()
        data = {
            "driveId": 0,
            "limit": 1000,
            "next": 0,
            "orderBy": "fileId",
            "orderDirection": "desc",
            "parentFileId": 0,
            "trashed": False,
            "Page": 1,
        }
        resp = requests.get(url, headers=headers, data=data)
        if resp.status_code == 200:
            # print(resp.text)
            jsoner = json.loads(resp.text)
            for v in jsoner["data"]["InfoList"]:
                print(v["FileName"])


if __name__ == "__main__":
    c = PanClient()
    # print(c)
