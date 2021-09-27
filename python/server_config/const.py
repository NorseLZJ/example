# -*- coding: utf-8 -*-

# define key

GameStartTime = 1
GameRegion = 2  # 区id
GameGroup = 3  # 服id
Latest = 4  # 合服标记
LatestArea = 5  # 合服后所有服的id
ServerID = 6  # 服务器id
ServerHost = 7  # 服务器外网ip
IntranetHost = 8  # 服务器内网ip

# mysql data
DataHost = 20
DataPort = 21
DataUser = 22
DataPasswd = 23

# mysqll log
LogHost = 30
LogPort = 31
LogUser = 32
LogPasswd = 33

# sql path
SqlPath = 40  # cur dir + dbservver\\sql\\

# db name
GuildDb = 44
DbLogDb = 45
NewGameDbDb = 46

GsNum = 60  # GameServer 数量
GateNum = 62  # gameGate 数量
WsGateNum = 63  # wsGate 数量

ServerPath = 64  # 服务器所在目录  绝对路径
MonitorPath = 65  # 监控脚本所在目录

GsSvn = 66
GateSvn = 67
WsGateSvn = 68

# connType 1:data 2:log
ConnTypeData = 1
ConnTypeLog = 2

centerDbName = "centerdb"

sql_db_exist = """
SELECT *
FROM information_schema.SCHEMATA
WHERE SCHEMA_NAME = '%s';
"""

sql_use_create = """
USE %s;
%s
"""

# 后面加上建表语句
sql_create_center_db = """
CREATE DATABASE IF NOT EXISTS %s;
USE %s;
%s
"""

sql_serverInfo = """
INSERT INTO %s.serverinfo
(id,
 origin_id,
 db_host,
 db_port,
 db_name,
 db_username,
 db_password,
 intranet_host,
 server_host)
VALUES (%d, %d, '%s', '%s', '%s', '%s', '%s', '%s', '%s');
"""

# 盘符 目录 远程链接　本地目录
svn_checkout_cmd = """
@echo off && %s: && cd %s && svn checkout %s %s -q
"""
