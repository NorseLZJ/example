import requests
import json


def get_bilibili_video_download_url(url):
    # 获取视频aid和cid
    response = requests.get(url)
    print(response.status_code)
    html = response.text
    aid_start = html.find("aid=") + 4
    aid_end = html.find("&", aid_start)
    aid = html[aid_start:aid_end]

    cid_start = html.find("cid=") + 4
    cid_end = html.find("&", cid_start)
    cid = html[cid_start:cid_end]

    # 构造API地址
    api_url = f"https://api.bilibili.com/x/player/playurl?aid={aid}&cid={cid}&qn=80&otype=json"

    # 发送API请求
    api_response = requests.get(api_url)
    api_data = json.loads(api_response.text)

    # 解析API返回的数据，获取视频下载地址
    video_url = api_data["data"]["durl"][0]["url"]

    return video_url


# 测试代码
video_url = get_bilibili_video_download_url("https://www.bilibili.com/video/BV12z4y1V7xB")
print(video_url)
