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
url = "https://59-47-225-195.d.cjjd15.com:30443/download-cdn.123pan.cn/123-241/f5ae483c/1812752236-0/f5ae483cccc0f14ec15120652eb84cf1?v=2&t=1680944884&s=1f0ed3221b9fce60eb7dbe3c6c06934d&filename=Encounter.of.the.Spooky.Kind.II.1990.WEB-DL.4K.H265.AAC-CTRLTV.mp4&d=f6acb5e3"


async def main():
    async with aiohttp.ClientSession() as session:
        data = await download_file(url, session)
        """
        需要知道文件名
        """
        with open("a.mp4", "wb") as f:
            f.write(data)


asyncio.run(main())
