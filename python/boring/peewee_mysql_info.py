"""
使用需要安装的库
pip install peewee pymysql
"""

from peewee import *

db = MySQLDatabase(
    "information_schema",
    user="liuzijian",
    password="123456",
    host="127.0.0.1",
    port=13306,
)


class TABLES(Model):
    TABLE_SCHEMA = CharField(null=False, primary_key=True)
    TABLE_NAME = CharField(null=False)

    class Meta:
        database = db  # 指定模型使用的数据库


class COLUMNS(Model):
    TABLE_SCHEMA = CharField(null=False, primary_key=True)
    TABLE_NAME = CharField()
    COLUMN_NAME = CharField()
    DATA_TYPE = CharField()

    class Meta:
        database = db  # 指定模型使用的数据库


db.connect()

tbs = (
    TABLES.select()
    .where(TABLES.TABLE_SCHEMA == "aiweb")
    .order_by(TABLES.TABLE_NAME.desc())
    .limit(3)
)
for v in tbs:
    print(f"{v.TABLE_NAME}")

v_list = (
    COLUMNS.select()
    .where(
        COLUMNS.TABLE_SCHEMA == "aiweb",
        COLUMNS.TABLE_NAME == "auth_user",
    )
    .order_by(COLUMNS.COLUMN_NAME.asc())
)
print("*" * 100)
for v in v_list:
    print(f"{v.TABLE_NAME} {v.COLUMN_NAME}")
