import requests
import json
import os
import hashlib
import base64

if __name__ == '__main__':
    url= "https://www.123pan.com/b/api/file/list/new?20312273=1702268655-4637517-2907810043&driveId=0&limit=100&next=0&orderBy=file_id&orderDirection=desc&parentFileId=2261251&trashed=false&SearchData=&Page=1&OnlyLookAbnormalFile=0&event=homeListFile&operateType=4&inDirectSpace=false"
    headers= {
        "App-Version":"3",
        "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDEzOTc3MzQsImlhdCI6MTcwMDc5MjkzNCwiaWQiOjE4MTI3NTIyMzYsIm1haWwiOiIiLCJuaWNrbmFtZSI6IuWRgOWYv-WRgCIsInN1cHBlciI6ZmFsc2UsInVzZXJuYW1lIjoxODg0MTY4NTA1NCwidiI6MH0.iwYebYUd3YvShJDh-dZJmqfJmfi9Jf8r4b1mW-aot2k",
        "Loginuuid":"c3c9311120f1888bc40ce6f9a3762591c692f339dc869e2f4ca6a57f9b7bfa93c0bad0618f1e0ee1d7a758f1c6dd5267",
        "Referer":"https://www.123pan.com/2261250/2261251",
        "Platform":"web"
    }
    resp = requests.get(url, headers=headers)
    if resp.status_code != 200:
        print("网络问题,检查网络连接,或尝试修复dns...稍后重试")
        exit(1)
    jsoner = json.loads(resp.text)
    if jsoner["code"] == 1:
        print(jsoner["message"])
        exit(0)
    # print(jsoner)
    for v in jsoner["data"]["InfoList"]:
        print(v["FileName"])