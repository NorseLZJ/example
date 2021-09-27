# coding=utf-8

import os
import shutil
from pattern import *
from const import *
from tool import *

pr_list_f_head = '''
# -*- coding: utf-8 -*-


%s

%s
'''

fmt_watch_pr_list = ''' 
pr_list = [
%s
]
'''

fmt_kill_pr_list = '''
kill_pr_list = [
%s
]

'''

gateConfigList = []

wsGateConfigList = []

pr_list = []

kill_pr_list = []

# keyfile,certfile 可能需要改变，根据域名变化
wsgate_template = '''
logcolor=false
loglevel=error
logfile=./log/wsgate.log
logfilemaxsize=52428800
wsaddr=http://0.0.0.0:%d/
gamenum=%d
%s
keyfile=./config/huanyumid.com.key
certfile=./config/huanyumid.com.pem
'''

gate_gs_head = ''' 
GsNum=%d
%s
'''


def clear_empty(ss):
    """
    :param ss: 原始字符串
    :return: 去掉空行后字符串
    """
    ss = ss.split('\n')
    temp = ''
    for i in range(0, len(ss)):
        if ss[i] == '':
            continue
        temp += ss[i] + '\n'
    return temp


wg_gg_temp = """
gameaddr%d=127.0.0.1:%d
"""

# 处理wsgate配置的时候使用
# 处理gamegate配置的时候填充
wg_gg_str = """"""


def Gate_to_wsGate_config(idx, port):
    cc = format(wg_gg_temp % (idx, port))
    global wg_gg_str
    wg_gg_str = format("%s%s" % (wg_gg_str, cc))


# 处理gate配置的时候使用
# 处理gameserver配置的时候填充
gate_gs_str = """"""

gs_to_gate_template = """
GS_%d_IP=127.0.0.1 
GS_%d_Port=%d
GS_%d_No=%d
"""


def AddProcess(d, name):
    s1 = d.split(":")[0]
    s2 = d.split(":")[1]

    exe = ''
    if d == name:
        # 第一个程序
        exe = format("%s.exe" % name)
    else:
        # 其它程序
        str_s = d.split('\\')
        exe = str_s[len(str_s) - 1] + ".exe"

    exe = exe.lower()
    global pr_list
    global kill_pr_list

    cmd = format("%s: && cd %s:%s && start \\\"%s\\\" /min %s" % (s1, s1, s2, exe, exe))

    temp = [exe, cmd]
    pr_list.append(temp)

    kill_temp = [exe, exe]
    kill_pr_list.append(kill_temp)


def AddProcessNoCopy(src):
    if not os.path.exists(src):
        return

    s1 = src.split(":")[0]
    s2 = src.split(":")[1]
    str_s = s2.split('\\')

    # gameServer,GameGate,wsgate no have index
    src_exe = str_s[len(str_s) - 1]
    src_exe = src_exe[0:len(src_exe) - 1]  # 去掉尾部的 数字
    src_exe = format("%s\\%s.exe" % (src, src_exe))

    print(src_exe)
    exe = str_s[len(str_s) - 1] + ".exe"
    name = src + "\\" + exe

    if not os.path.exists(name):
        if os.path.exists(src_exe):
            shutil.copy(src_exe, name)
        else:
            return

    exe = exe.lower()
    s2 = s2.replace('\\', '/')
    cmd = format("%s: && cd %s:%s && start \\\"%s\\\" /min %s" % (s1, s1, s2, exe, exe))
    # cmd = format("%s: && cd %s:%s && start %s" % (s1, s1, s2, exe))
    temp = [exe, cmd]

    global pr_list
    pr_list.append(temp)

    global kill_pr_list
    kill_temp = [exe, exe]
    kill_pr_list.append(kill_temp)


def process_list(start_file):
    temp_watch = ""
    for k in range(0, len(pr_list)):
        vv = pr_list[k]
        temp_watch += ' ' * 4 + '["' + vv[0] + '","' + vv[1] + '"],'
        if k < len(pr_list) - 1:
            temp_watch += '\n'

    watch_str = format(fmt_watch_pr_list % temp_watch)

    # need kill proc
    temp_kill = ""
    for k in range(0, len(kill_pr_list)):
        vv = kill_pr_list[k]
        temp_kill += ' ' * 4 + '["' + vv[0] + '","' + vv[1] + '"],'
        if k < len(kill_pr_list) - 1:
            temp_kill += '\n'

    kill_str = format(fmt_kill_pr_list % temp_kill)

    f = open(start_file, 'w+')
    code = format(pr_list_f_head % (watch_str, kill_str))
    f.write(code)
    f.close()
    print('-' * 50)


def other_serve(cur_path=""):
    if cur_path == "":
        cur_path = os.getcwd()

    ss = format("%s\\%s" % (cur_path, 'NewMServer'))
    AddProcessNoCopy(ss)

    # ss = format("%s\\%s" % (cur_path, 'CenterServer'))
    # AddProcessNoCopy(ss)

    ss = format("%s\\%s" % (cur_path, 'loginserver'))
    AddProcessNoCopy(ss)

    ss = format("%s\\%s" % (cur_path, 'dbserver'))
    AddProcessNoCopy(ss)

    # ss = format("%s\\%s" % (cur_path, 'WebServer'))
    # AddProcessNoCopy(ss)

    # ss = format("%s\\%s" % (cur_path, 'WorldServer'))
    # AddProcessNoCopy(ss)


def Gs_to_gate_config(idx, port):
    cc = format(gs_to_gate_template % (idx, idx, port, idx, idx))
    global gate_gs_str
    gate_gs_str = format("%s%s" % (gate_gs_str, cc))


class ModifyConf(object):
    def __init__(self, cf):
        self.cf = cf
        self.cur_path = str(self.cf.get_val(ServerPath))
        self.gs_num = int(self.cf.get_val(GsNum))
        self.gate_num = int(self.cf.get_val(GateNum))
        self.wsgate_num = int(self.cf.get_val(WsGateNum))
        self.region = int(self.cf.get_val(GameRegion))
        self.group = int(self.cf.get_val(GameGroup))
        self.latest = int(self.cf.get_val(Latest))
        self.latest_area = str(self.cf.get_val(LatestArea))

        self.new_game_db = str(self.cf.get_val(NewGameDbDb))
        self.dblog_db = str(self.cf.get_val(DbLogDb))
        self.guild_db = str(self.cf.get_val(GuildDb))

    def modify_config(self):
        monitor_path = str(self.cf.get_val(MonitorPath))
        proc_list_path = format("%sproclist.py" % monitor_path)

        self.db_config()
        self.login_server_config()
        self.new_m_server_config()
        self.gs_config()
        self.gate_config()
        self.wsgate_config()

        other_serve(self.cur_path)
        process_list(proc_list_path)

    def gate_config(self):

        gs_str = format(gate_gs_head % (self.gs_num, gate_gs_str))
        port = 6700
        print('GateNum:', self.gate_num)
        for idx in range(0, self.gate_num):
            temp = path_by_idx(self.cur_path, 'gamegate', idx)

            ini = format("%s\\gamegate.ini" % temp)
            channel = format("%s\\gg_channel.xml" % temp)
            print(port, idx, ini, '#' * 13, channel)

            cent = r_file(ini, encoding='utf-8')
            cent = replace_gate_ini(cent, self.region, self.group, idx, port, self.latest, self.latest_area, gs_str)
            w_file(ini, cent, encoding='utf-8')

            cent = r_file(channel)
            cent = replace_gate_channel(cent, self.region, self.group, idx)
            w_file(channel, cent)

            AddProcessNoCopy(temp)
            Gate_to_wsGate_config(idx, port)
            port += 1

    def new_m_server_config(self):
        ini = format("%s\\NewMserver\\newmserver.ini" % self.cur_path)
        channel = format("%s\\NewMserver\\ms_channel.xml" % self.cur_path)

        cent = r_file(ini)
        cent = replace_newmserver_ini(cent, self.region, self.group, self.dblog_db, self.guild_db)
        w_file(ini, cent)

        cent = r_file(channel)
        cent = replace_newmserver_channel(cent, self.region, self.group)
        w_file(channel, cent)

    def login_server_config(self):
        channel = format("%s\\loginserver\\ls_channel.xml" % self.cur_path)

        cent = r_file(channel)
        cent = replace_login_server_channel(cent, self.region, self.group)
        w_file(channel, cent)

    def db_config(self):
        ini = format("%s\\dbserver\\dbserver.ini" % self.cur_path)
        channel = format("%s\\dbserver\\db_channel.xml" % self.cur_path)

        cent = r_file(ini)
        cent = replace_db_server_ini(cent, self.new_game_db)
        w_file(ini, cent)

        cent = r_file(channel)
        cent = replace_db_server_channel(cent, self.region, self.group)
        w_file(channel, cent)

    def gs_config(self):
        port = 6600
        for idx in range(0, self.gs_num):
            temp = path_by_idx(self.cur_path, "GameServer", idx)

            ini = format("%s\\gameserver.ini" % temp)
            channel = format("%s\\gs_channel.xml" % temp)
            print(port, idx, ini, '#' * 13, channel)

            cent = r_file(ini)
            cent = replace_gs_ini(cent, self.region, self.group, idx, port, self.gs_num)
            w_file(ini, cent)

            cent = r_file(channel)
            cent = replace_gs_channel(cent, self.region, self.group, idx)
            w_file(channel, cent)

            Gs_to_gate_config(idx, port)
            AddProcessNoCopy(temp)
            port += 1

    def wsgate_config(self):
        port = 8001
        for idx in range(0, self.wsgate_num):
            temp = path_by_idx(self.cur_path, 'wsgate', idx)
            cfg = format("%s\\config\\gate.cfg" % temp)
            print(port, cfg)

            c_str = format(wsgate_template % (port, self.gate_num, wg_gg_str))
            c_str = clear_empty(c_str)

            with open(cfg, 'w') as f:
                f.write(c_str)

            AddProcessNoCopy(temp)
            port += 1
