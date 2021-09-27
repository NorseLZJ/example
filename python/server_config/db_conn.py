# -*- coding: utf-8 -*-

import pymysql
from const import *


def connect_db(host, port, user, passwd):
    port = int(port)
    db = pymysql.connect(host=host, port=port, user=user, passwd=passwd)
    db.set_charset('utf8')
    return db


class Conn(object):
    def __init__(self, conn):
        self.conn = conn

    def exec_sql(self, sql):
        ex = self.conn.cursor()
        if ex is None:
            raise Exception("conn is nil")

        try:
            ex.execute(sql)
        except Exception as e:
            # print("sql:%s err:%s" % (sql, e))
            print("err:%s\n" % e)
            # exit(1)

    def exec_sql_slice(self, sql_s):
        ex = self.conn.cursor()
        if ex is None:
            raise Exception("conn is nil")

        for i in sql_s:
            try:
                ex.execute(str(i))
            except Exception as e:
                # print("sql:%s err:%s" % (i, e))
                print("err:%s\n" % e)

    def insert_sql(self, sql):
        ex = self.conn.cursor()
        if ex is None:
            raise Exception("conn is nil")

        try:
            ex.execute(sql)
            # self.conn.commit()
        except Exception as e:
            # print("sql:%s err:%s" % (sql, e))
            print("err:%s\n" % e)
            # exit(1)

    def db_exist(self, db):
        ex = self.conn.cursor()
        if ex is None:
            raise Exception("conn is nil")

        sql = format(sql_db_exist % db)
        ex.execute(sql)

        row = ex.fetchone()
        if row is not None:
            return True

        return False

    def db_not_exist_create(self, db):
        if True is self.db_exist(db):
            return False
        sql = "CREATE DATABASE IF NOT EXISTS %s default character set utf8mb4 collate utf8mb4_unicode_ci;" % db
        self.exec_sql(sql)

        return True

    def close(self):
        self.conn.close()


class DataConn(Conn):
    def __init__(self, cf):
        host = cf.configMap[DataHost]
        port = cf.configMap[DataPort]
        user = cf.configMap[DataUser]
        pwd = cf.configMap[DataPasswd]
        conn = connect_db(host, port, user, pwd)
        if conn is None:
            return
        super(DataConn, self).__init__(conn)


class LogConn(Conn):
    def __init__(self, cf):
        host = cf.configMap[LogHost]
        port = cf.configMap[LogPort]
        user = cf.configMap[LogUser]
        pwd = cf.configMap[LogPasswd]
        conn = connect_db(host, port, user, pwd)
        if conn is None:
            return
        super(LogConn, self).__init__(conn)
