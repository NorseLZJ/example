import aiohttp
import asyncio


async def download_chunk(url, start, end, session):
    headers = {"Range": f"bytes={start}-{end}"}
    async with session.get(url, headers=headers) as response:
        return await response.read()


async def download_file(url, session):
    async with session.head(url) as response:
        size = int(response.headers["Content-Length"])
    chunk_size = size // 4
    tasks = []
    for i in range(4):
        start = i * chunk_size
        end = start + chunk_size - 1
        if i == 3:
            end = size - 1
        tasks.append(download_chunk(url, start, end, session))
    chunks = await asyncio.gather(*tasks)
    return b"".join(chunks)


"""
需要替换,一定的时间之后就会失效
"""
url = "https://36-134-210-67.pd1.123pan.cn:30443/download-cdn.123pan.cn/123-402/3c5b8451/1630998-0/3c5b845145ef9b6368b48d2796113592/c-m5?v=5&t=1695732060&s=1695732060842387033b799b611e4558cc6bd3a0a8&r=VR1FN4&bzc=1&bzs=1812752236&filename=Fedora-Workstation-Live-x86_64-38-1.6.iso&x-mf-biz-cid=823eb311-aae8-4a20-ac42-f9b45e7abcea-a0d664&auto_redirect=0&xmfcid=07c4840b-5783-478d-9347-153316b4ae3d-0-abf611255"


async def main():
    async with aiohttp.ClientSession() as session:
        data = await download_file(url, session)
        """
        需要知道文件名
        """
        with open("a.iso", "wb") as f:
            f.write(data)


asyncio.run(main())
