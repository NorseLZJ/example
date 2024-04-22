"""
使用需要安装的库
pip install peewee pymysql
"""

from peewee import *

db = MySQLDatabase(
    "dopai", user="liuzijian", password="123456", host="127.0.0.1", port=13306
)


class dopai_master_server_state(Model):
    key = CharField(unique=True)
    id = IntegerField()
    type = CharField()
    host = CharField()
    port = IntegerField()
    create_time = IntegerField()
    tcp_port = IntegerField()

    class Meta:
        database = db  # 指定模型使用的数据库


db.connect()
# db.create_tables([DopaiMasterServerState])

# val = dopai_master_server_state.get(dopai_master_server_state.key == "commerce:10001")
values = dopai_master_server_state.select()
for v in values:
    print(f"{v.id} {v.key} {v.type}")
