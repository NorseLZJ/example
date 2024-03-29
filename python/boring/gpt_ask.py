import os

import httpx
from openai import OpenAI

# openai_key = os.getenv("OPENAI_KEY")
openai_key = os.getenv("OPENAI_KEY")
if openai_key is None:
    print("openai key not found")
    exit(0)

client = OpenAI(
    api_key=openai_key,
    # TODO 此处的代理，如果脚本在国外或者外网的机器上执行，是不需要的
    http_client=httpx.Client(proxy="http://127.0.0.1:7890"),
    timeout=4,
)

if __name__ == "__main__":
    try:
        completion = client.chat.completions.create(
            model="gpt-3.5-turbo",
            messages=[
                {"role": "system",
                 "content": "You are a poetic assistant, skilled in explaining complex programming concepts with creative flair."},
                {"role": "user", "content": "Compose a poem that explains the concept of recursion in programming."}
            ]
        )
        print(completion.choices[0].message)
    except Exception as e:
        print(e)
