import requests
import json
import os
import hashlib
import base64

# _url = "http://127.0.0.1:8200/"
# _url = "https://123pan.com:8443/"
_url = "https://123pan.com/"

passport = ""
# password = "111111"
password = ""


def Exit(msg: str, val: int = 0):
    print(msg)
    if val != 0:
        exit(val)
    exit(1)


def file_info(path: str) -> tuple:
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
        self.platform = "actix-web"
        self.app_version = 111
        # self.platform = "web"
        # self.platform = "pc"
        self.sign()
        self.user_info()
        # self.upload_request()
        self.file_list()
        # self.download_file(file_id=341)
        # self.download_file_info("1812752236-0", "Book1.pdf", "f4c5ffdd8a2bb11e685521060c12e50e", 77734)

    def download_file_info(self, s3keyflag: str, file_name: str, etag: str, size: int):
        """
        需要数据库查询 s3key_flag,参数才行
        """
        url = format("%s/api/file/download_info" % (_url))
        headers = self.headers(pc=True)
        data = {
            "file_id": 1,
            "s3key_flag": s3keyflag,
            "drive_id": 0,
            "file_name": file_name,
            "size": size,
            "etag": etag,
        }
        resp = requests.post(url, headers=headers, data=data)
        if resp.status_code != 200:
            print("网络问题,检查网络连接,或尝试修复dns...稍后重试")
            exit(1)
        # print(resp.text)
        jsoner = json.loads(resp.text)
        if jsoner["code"] == 1:
            print(jsoner["message"])
            exit(0)
        address = jsoner["data"]["DownloadUrl"]
        print(address)

    def download_file(self, file_name: str, file_id: int, size: int, etag: str):
        # mp4结尾
        if file_name.endswith("mp4") is False:
            return
        url = format("%s/api/file/download?file_id=%d" % (_url, file_id))
        headers = self.headers(pc=True)
        resp = requests.get(url, headers=headers)
        if resp.status_code != 200:
            print("网络问题,检查网络连接,或尝试修复dns...稍后重试")
            exit(1)
        # print(resp.text)
        jsoner = json.loads(resp.text)
        if jsoner["code"] == 1:
            print(jsoner["message"])
            exit(0)
        address = jsoner["data"]["url"]
        resp = requests.get(address, allow_redirects=False)
        if resp.status_code != 302:
            Exit("调度节点失败了")
        address = resp.headers.get("Location")
        print("%s\n%dM\n%d\n%s\n%s" % (file_name, int(size / 1024 / 1024), size, address, etag))

    def __str__(self) -> str:
        return format("nick_name:%s\nspaceUsed:%d\nspacePermanent:%d" % (self.nick_name, self.space_used, self.space_permanent))

    def headers(self, token: bool = True, pc: bool = False):
        ret = {
            "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
            # "Host": "123pan.com:8080",
            "App-Version": "1.2",
            "platform": self.platform,
        }
        if token:  # 第一次登录没有token
            ret["Authorization"] = format("Bearer %s" % (self.token))

        if pc is True:
            ret["App-Version"] = "1.0.115"
            ret["platform"] = "pc]"
        return ret

    def sign(self):
        url = format("%s/api/user/sign_in" % (_url))
        headers = self.headers(token=False)
        resp = requests.post(url, headers=headers, data={"passport": passport, "password": password, "remember": True})
        if resp.status_code == 200:
            jsoner = json.loads(resp.text)
            if jsoner["message"] == "success":
                self.expire = jsoner["data"]["expire"]
                self.token = jsoner["data"]["token"]
                print("登录成功...")
            else:
                Exit("登陆失败")

    def user_info(self):
        url = format("%s/api/user/info" % (_url))
        headers = self.headers()
        resp = requests.get(url, headers=headers)
        if resp.status_code == 200:
            jsoner = json.loads(resp.text)
            self.nick_name = jsoner["data"]["Nickname"]
            self.space_used = jsoner["data"]["SpaceUsed"]
            self.space_permanent = jsoner["data"]["SpacePermanent"]
        else:
            Exit("网络不通，检查网络连接!")

    def upload_request(self):
        url = format("%s/api/file/upload_request" % (_url))
        headers = self.headers()
        path = "D:\\document\\kong.md"
        file_name = "kong.md"
        size, filename, md5 = file_info(path)
        data = {
            "driveId": 0,
            "etag": md5,
            "fileName": filename,
            "parentFileId": 1829700,
            "size": size,
            "type": 0,
            "duplicate": 0,
        }
        # print(data)
        resp = requests.post(url, headers=headers, data=data)
        if resp.status_code == 200:
            jsoner = json.loads(resp.text)
            if jsoner["code"] != 0:
                print(jsoner)
                return
            bucket = jsoner["data"]["Bucket"]
            key = jsoner["data"]["Key"]
            uploadId = jsoner["data"]["UploadId"]
            print("-" * 50)
            print("bucket:%s key:%s md5:%s size:%d" % (bucket, key, md5, size))
            print("-" * 50)
            file_id = jsoner["data"]["FileId"]
            self.s3_list_upload_parts(bucket, key, uploadId)
            self.s3_repare_upload_parts_batch(bucket, key, uploadId)
            self.s3_complete_multipart_upload(bucket, key, uploadId)
            # self.upload_complete(file_id)
            self.download_info2(md5, file_id, file_name, size)

    def download_info2(self, etag, file_id, file_name, size):
        url = format("%s/api/file/download_info" % (_url))
        headers = self.headers()
        data = {
            "driveId": 0,
            "etag": etag,
            "fileId": file_id,
            "s3keyFlag": "1746916-0",
            "type": 0,
            "fileName": file_name,
            "size": size,
        }
        resp = requests.post(url, headers=headers, data=data)
        if resp.status_code != 200:
            print("download_info 出错, %s " % resp.text)
            exit(1)
        # print(resp.text)
        jsoner = json.loads(resp.text)
        v = jsoner["data"]["DownloadUrl"].split("params=")[1]
        # print("\nDownloadUrl :%s\n" % v)
        ret = base64.decodebytes(v.encode())
        print(ret)

    def upload_complete(self, file_id: int):
        url = format("%s/api/file/upload_complete" % (_url))
        headers = self.headers()
        data = {"fileId": file_id}
        resp = requests.post(url, headers=headers, data=data)
        if resp.status_code != 200:
            print("upload_complete 出错, %s " % resp.text)
            exit(1)

    def s3_complete_multipart_upload(self, bucket: str, key: str, uploadId: str):
        url = format("%s/api/file/s3_complete_multipart_upload" % (_url))
        headers = self.headers()
        data = {
            "bucket": bucket,
            "key": key,
            "uploadId": uploadId,
        }
        resp = requests.post(url, headers=headers, data=data)
        if resp.status_code != 200:
            print("s3_complete_multipart_upload 出错, %s " % resp.text)
            exit(1)

    def s3_list_upload_parts(self, bucket: str, key: str, uploadId: str):
        url = format("%s/api/file/s3_list_upload_parts" % (_url))
        headers = self.headers()
        data = {
            "bucket": bucket,
            "key": key,
            "uploadId": uploadId,
        }
        resp = requests.post(url, headers=headers, data=data)
        if resp.status_code != 200:
            print("s3_list_upload_parts 出错, %s " % resp.text)
            exit(1)

    def s3_repare_upload_parts_batch(self, bucket: str, key: str, uploadId: str):
        url = format("%s/api/file/s3_repare_upload_parts_batch" % (_url))
        headers = self.headers()
        data = {
            "bucket": bucket,
            "key": key,
            "partNumberEnd": 2,  # 没看start，end有啥规律，先填一下看看
            "partNumberStart": 1,
            "uploadId": uploadId,
        }
        resp = requests.post(url, headers=headers, data=data)
        if resp.status_code == 200:
            # print(resp.text)
            jsoner = json.loads(resp.text)
            up_urls = jsoner["data"]["presignedUrls"]["1"]
            self.cjdd_up(cjdd_url=up_urls)

    def cjdd_up(self, cjdd_url: str):
        # path = "D:\\迅雷下载\\kali-linux-2022.3-live-amd64.iso"
        path = "D:\\迅雷下载\\go1.18.7.windows-amd64.msi"
        with open(path, "rb") as f:
            resp = requests.put(cjdd_url, data=f.read())
            if resp.status_code == 200:
                print("上传完成了")

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
                if v["Type"] == 1:  # 文件夹
                    continue
                self.download_file(v["FileName"], v["FileId"], v["Size"], v["Etag"])


if __name__ == "__main__":
    c = PanClient()
    # print(c)
