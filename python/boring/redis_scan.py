"""
删除限制
"""
import redis as rdb

# 测试redis
r = rdb.Redis(host="127.0.0.1", port=6379, password="", db=0)

try:
    r.ping()
    print("连接正常")
except Exception as e:
    print("检查参数代理")
    exit(1)

decode_type = "UTF-8"

if __name__ == "__main__":
    kkeys = []
    cursor, keys = r.scan()
    for key in keys:
        kkeys.append(key.decode(decode_type))
    while cursor:
        cursor, keys = r.scan(cursor=cursor)
        for key in keys:
            kkeys.append(key.decode(decode_type))

    for k in kkeys:
        print("%-10s %s" % (r.type(k).decode(decode_type), k))
    print(len(kkeys))
