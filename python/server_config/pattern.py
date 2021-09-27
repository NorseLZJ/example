# -*- coding: utf-8 -*-

import re

'''
使用到的正则
'''

# gameServer ini
p_gs_area = re.compile(r"Area=\d*")
p_gs_group = re.compile(r"Group=\d*")
p_gs_server_port = re.compile(r"ServerPort=\d*")
p_gs_server_count = re.compile(r"ServerCount=\d*")
p_gs_server_idx = re.compile(r"ServerIndex=\d*")

# gameServer channel
p_gs_ch_server_id = re.compile(r"<server id=(.+?)>")
p_gs_ch_item1 = re.compile(r"<item id=(.+?)-(.+?)-")
p_gs_ch_item2 = re.compile(r"<item id=(.+?)-(.+?)-8-0")

# gameGate ini
p_gg_server_port = re.compile(r"ListenPort=\d*")
p_gg_gate_pipe_id = re.compile(r"GatePipeID=(\d)-(\d)-(\d)-(\d)")
p_gg_gate_latest = re.compile(r"Latest=\d*")
p_gg_gate_latest_area = re.compile(r"LatestArea=.*")
p_gg_gate_gs_str = re.compile(r"\[GameServer](.*?)\[Debug]", re.S)

# gameGate channel
p_gg_ch_server_id = re.compile(r"<server id=(.+?)>")
p_gg_ch_item = re.compile(r"<item id=(.+?)-(.+?)-")

# db server
p_db_server_channel = re.compile(r"<server id=(.+?)>")

p_db_server_dbname1 = re.compile(r"Mysql_1_DBName=.*")
p_db_server_dbname2 = re.compile(r"Mysql_2_DBName=.*")

# login server
p_login_server_channel = re.compile(r"<server id=(.+?)>")

# newMServer
p_newMServer_channel = re.compile(r"<server id=(.+?)>")

# newMServer ini
p_newMServer_ini_area = re.compile(r"Area=\d*")
p_newMServer_ini_group = re.compile(r"Group=\d*")
p_newMServer_ini_sql_dblog = re.compile(r"MysqlDBLog=.*")
p_newMServer_ini_sql_logic = re.compile(r"MysqlLogic=.*")


def replace_newmserver_channel(cent, region, group):
    cent = re.sub(p_newMServer_channel, "<server id=\"%d-%d-5-0\"/>" % (region, group), cent)
    return cent


def replace_newmserver_ini(cent, region, group, db_dblog, db_guild):
    cent = re.sub(p_newMServer_ini_area, "Area=%d" % region, cent)
    cent = re.sub(p_newMServer_ini_group, "Group=%d" % group, cent)
    cent = re.sub(p_newMServer_ini_sql_dblog, "MysqlDBLog=%s" % db_dblog, cent)
    cent = re.sub(p_newMServer_ini_sql_logic, "MysqlLogic=%s" % db_guild, cent)
    return cent


def replace_login_server_channel(cent, region, group):
    cent = re.sub(p_login_server_channel, "<server id=\"%d-%d-2-0\"/>" % (region, group), cent)
    return cent


def replace_db_server_channel(cent, region, group):
    cent = re.sub(p_db_server_channel, "<server id=\"%d-%d-3-0\"/>" % (region, group), cent)
    return cent


def replace_db_server_ini(cent, dbname):
    cent = re.sub(p_db_server_dbname1, "Mysql_1_DBName=%s" % dbname, cent)
    cent = re.sub(p_db_server_dbname2, "Mysql_2_DBName=%s" % dbname, cent)
    return cent


def replace_gate_ini(cent, region, group, idx, port, latest, latest_area, gs_str):
    cur = int(4)  # 这个值，好像无关紧要，就拿配置里边的一个先写一下
    # pipe = format("GatePipeID=%d-%d-%d-%d" % (region, group, cur, idx))
    cent = re.sub(p_gg_server_port, "ListenPort=%d" % port, cent)
    cent = re.sub(p_gg_gate_pipe_id, "GatePipeID=%d-%d-%d-%d" % (region, group, cur, idx), cent)
    cent = re.sub(p_gg_gate_latest, "Latest=%d" % latest, cent)
    cent = re.sub(p_gg_gate_latest_area, "LatestArea=%s" % latest_area, cent)
    cent = re.sub(p_gg_gate_gs_str, "[GameServer]%s\n[Debug]\n" % gs_str, cent)
    return cent


def replace_gate_channel(cent, region, group, idx):
    cent = re.sub(p_gs_ch_server_id, "<server id=\"%d-%d-4-%d\"/>" % (region, group, idx), cent)
    cent = re.sub(p_gg_ch_item, "<item id=\"%d-%d-" % (region, group), cent)
    return cent


def replace_gs_ini(cent, region, group, idx, port, num):
    cent = re.sub(p_gs_area, "Area=%d" % region, cent)
    cent = re.sub(p_gs_group, "Group=%d" % group, cent)
    cent = re.sub(p_gs_server_idx, "ServerIndex=%d" % idx, cent)
    cent = re.sub(p_gs_server_port, "ServerPort=%d" % port, cent)
    cent = re.sub(p_gs_server_count, "ServerCount=%d" % num, cent)
    return cent


def replace_gs_channel(cent, region, group, idx):
    cent = re.sub(p_gs_ch_server_id, "<server id=\"%d-%d-6-%d\"/>" % (region, group, idx), cent)
    cent = re.sub(p_gs_ch_item1, "<item id=\"%d-%d-" % (region, group), cent)
    cent = re.sub(p_gs_ch_item2, "<item id=\"1-1-8-0", cent)
    return cent
