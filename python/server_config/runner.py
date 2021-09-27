# -*- coding: utf-8 -*-

from db_conn import *
from modify_config import *
from tool import *


class Runner(object):
    def __init__(self, cf):
        self.cf = cf

        self.data_conn = DataConn(cf)
        self.log_conn = LogConn(cf)

    def run(self):
        self.insert_server_info()
        self.create_db_and_table()

        self.svn_resp()
        mc = ModifyConf(self.cf)
        mc.modify_config()

        self.data_conn.close()
        self.log_conn.close()

    def svn_resp(self):
        num = int(self.cf.get_val(GsNum))
        if num > 1:
            resp = str(self.cf.get_val(GsSvn))
            self.svn_co_cmd(resp, 'GameServer', num)

        num = int(self.cf.get_val(GateNum))
        if num > 1:
            resp = str(self.cf.get_val(GateSvn))
            self.svn_co_cmd(resp, 'gamegate', num)

        num = int(self.cf.get_val(WsGateNum))
        if num > 1:
            resp = str(self.cf.get_val(WsGateSvn))
            self.svn_co_cmd(resp, 'wsgate', num)

    def svn_co_cmd(self, resp, co_name, max_num):
        path = str(self.cf.get_val(ServerPath))
        s1 = path.split(':')[0]
        for i in range(1, max_num):
            name = format("%s%d" % (co_name, i))
            cmd = format(svn_checkout_cmd % (s1, path, resp, name))
            os.system(cmd)

    def read_exec_sql(self, conn_type, path, db):
        with open(path, 'r', encoding='utf8') as f:
            txt = f.read()
            sql = format(sql_use_create % (db, txt))
            if conn_type == ConnTypeData:
                self.data_conn.exec_sql_slice(decode_sql(sql))
            if conn_type == ConnTypeLog:
                self.log_conn.exec_sql_slice(decode_sql(sql))

    def create_db_and_table(self):
        cf = self.cf
        temp = self.cf.get_val(SqlPath)
        db = cf.get_val(GuildDb)

        if True is self.data_conn.db_not_exist_create(db):
            cc = format("%sguild.sql" % temp)
            self.read_exec_sql(ConnTypeData, cc, db)

        db = cf.get_val(NewGameDbDb)
        if True is self.data_conn.db_not_exist_create(db):
            cc = format("%snewgamedb.sql" % temp)
            self.read_exec_sql(ConnTypeData, cc, db)

        db = cf.get_val(DbLogDb)
        if True is self.log_conn.db_not_exist_create(db):
            cc = format("%sdblog.sql" % temp)
            self.read_exec_sql(ConnTypeLog, cc, db)

    def create_server_info(self):
        ret = self.data_conn.db_exist(centerDbName)
        if ret is True:
            return

        # 创建数据库 和 表结构
        temp = self.cf.get_val(SqlPath)
        sql_path = format("%scenterdb.sql" % temp)
        with open(sql_path, 'r', encoding='utf8') as f:
            txt = f.read()
            sql = format(sql_create_center_db % (centerDbName, centerDbName, txt))
            self.data_conn.exec_sql_slice(decode_sql(sql))

    def insert_server_info(self):
        cf = self.cf
        self.create_server_info()
        sid = int(cf.get_val(ServerID))
        host = cf.get_val(DataHost)
        port = str(cf.get_val(DataPort))
        dbname = cf.get_val(GuildDb)
        user = cf.get_val(DataUser)
        pwd = str(cf.get_val(DataPasswd))
        s_host = cf.get_val(ServerHost)
        i_host = cf.get_val(IntranetHost)
        sql = format(sql_serverInfo % (centerDbName, sid, sid, host, port, dbname, user, pwd, s_host, i_host))
        self.data_conn.insert_sql(sql)
