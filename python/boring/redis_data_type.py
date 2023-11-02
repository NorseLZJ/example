"""
redis 数据类型测试 
string
hash
set 
sort set
list
"""

import redis

# 连接Redis服务器
r = redis.Redis(host="172.28.241.123", port=6379)
try:
    r.ping()
except Exception as e:
    print(e)
    exit(1)

# 测试String类型
r.set("my_key", "Hello Redis!")
print(r.get("my_key"))

# 测试Hash类型
r.hset("my_hash", "field1", "value1")
r.hset("my_hash", "field2", "value2")
print(r.hgetall("my_hash"))

# 测试List类型
r.lpush("my_list", "item1")
r.lpush("my_list", "item2")
print(r.lrange("my_list", 0, -1))

# 测试Set类型
r.sadd("my_set", "item1")
r.sadd("my_set", "item2")
print(r.smembers("my_set"))

"""
应用场景 白名单
"""


# 测试Sorted Set类型
r.zadd("my_sorted_set", {"item1": 1.99, "item2": 1.88})
print(r.zrangebyscore("my_sorted_set", 1.01, 1.92))

"""
应用场景 排行榜
"""
