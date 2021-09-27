# -*- coding: utf-8 -*-

import requests
import re
import socket
import subprocess


class Conf(object):
    def __init__(self):
        self.configMap = {}

    def append(self, key, val):
        self.configMap[key] = val

    def get_val(self, key):
        return self.configMap[key]


def get_ip_www():
    """
    get ip www
    :return:
    """
    ip_line = re.compile(r'你的外网IP地址是.*')
    ip = re.compile(r'\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3}')

    headers = {
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36"
    }
    resp = requests.get('https://tool.lu/ip/', headers=headers)
    html = str(resp.text.encode('utf-8'))
    ss = re.findall(ip_line, html)
    if len(ss) != 1:
        return ""
    ret = str(ss[0])
    ret = re.findall(ip, ret)
    if len(ret) != 1:
        return ""
    return str(ret[0])


def get_ip():
    """
    get ip like 192.168.1.xxx
    :return:
    """
    ret = socket.gethostname()
    ret = socket.gethostbyname(ret)
    return ret


def get_svn_resp(d):
    s1 = str(d).split(':')[0]
    cmd = format("%s: && cd %s && svn info --xml" % (s1, d))
    info, err = windows_cmd(cmd)
    url_re = re.compile(r'<url>.*</url>')
    ss = re.findall(url_re, str(info))
    if len(ss) != 1:
        return ""
    temp = (ss[0])
    # del : <url> </url>
    temp = temp[5:len(temp) - 6]
    return temp


def windows_cmd(cmd):
    backup = subprocess.Popen(cmd, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
    stdout, stderr = backup.communicate()
    return stdout, stderr


def decode_sql(sql):
    prefix = '/*'
    suffix = '*/'
    proc_head = 'DELIMITER ;;'
    proc_tail = 'DELIMITER ;'
    str_s = sql.split('\n')
    if len(str_s) <= 0:
        return None

    ret = []
    is_comment = False
    is_proc = False

    buff = ''
    proc_buff = ''
    for v in str_s:
        if prefix == v:  # 段注释开头
            is_comment = True
        if suffix == v:  # 段注释结尾
            is_comment = False

        if proc_head == v:  # 存储过程开始
            is_proc = True
            proc_buff += v + '\n'
            continue
        if proc_tail == v and is_proc is True:  # 存储过程结束
            proc_buff += v + '\n'
            ret.append(proc_buff)
            is_proc = False
            proc_buff = ''
            continue
        if is_proc:
            proc_buff += v + '\n'
            continue

        if (prefix in v and suffix in v) or ('--' in v):  # 单行注释
            continue

        # 下面处理别的内容
        if is_comment is False and v != suffix:
            if ';' in v and buff == '':  # 单行语句
                ret.append(v)
                continue
            if v == '':
                continue

            buff += v + '\n'
            if ';' in v and buff != '' and 'CREATE TABLE' in buff:
                buff = buff[:len(buff) - 1]
                ret.append(buff)
                buff = ''
    return ret


def r_file(p, encoding=""):
    if encoding != "":
        with open(p, 'r', encoding=encoding) as f:
            cent = f.read()
            return cent

    with open(p, 'r') as f:
        cent = f.read()
        return cent


def w_file(p, cent, encoding=""):
    if encoding != "":
        with open(p, 'w', encoding=encoding) as f:
            f.write(cent)
            return

    with open(p, 'w') as f:
        f.write(cent)


def path_by_idx(s1, s2, idx):
    if idx == 0:
        return format("%s\\%s" % (s1, s2))
    else:
        return format("%s\\%s%d" % (s1, s2, idx))
