# -*- coding: utf-8 -*-

import os
import sys
import time
import config as u
import tool
import runner as r
from const import *
from PyQt5.QtWidgets import QApplication, QMainWindow


def init_svn(cf):
    temp = str(cf.get_val(ServerPath))

    v = format("%sGameServer\\" % temp)
    resp = tool.get_svn_resp(v)
    cf.append(GsSvn, resp)

    v = format("%sgamegate\\" % temp)
    resp = tool.get_svn_resp(v)
    cf.append(GateSvn, resp)

    v = format("%swsgate\\" % temp)
    resp = tool.get_svn_resp(v)
    cf.append(WsGateSvn, resp)


def init_dir(cf):
    temp_path = os.getcwd()
    idx = temp_path.rfind('\\')
    sql_path = format("%s\\dbserver\\sql\\" % temp_path[0:idx])
    cf.append(SqlPath, sql_path)

    services_path = format("%s\\" % temp_path[0:idx])
    cf.append(ServerPath, services_path)

    monitor_path = format("%s\\py_script\\monitor\\" % temp_path[0:idx])
    cf.append(MonitorPath, monitor_path)


def init_dbname(cf, region, group):
    temp = str(region * 1000 + group)
    guild_db = format("guild_%s" % temp)
    log_db = format("dblog_%s" % temp)
    newgamedb_db = format("newgamedb_%s" % temp)
    cf.append(GuildDb, guild_db)
    cf.append(DbLogDb, log_db)
    cf.append(NewGameDbDb, newgamedb_db)


def submit():
    cf = tool.Conf()
    region = int(ui.region_2.text())
    group = int(ui.group.text())
    latest = int(ui.latest.text())
    latest_area = str(ui.latestArea.text())
    s_host = str(ui.serverHost.text())
    i_host = str(ui.intraneHost.text())
    game_start = str(ui.gameStartTime.text())

    cf.append(GameRegion, region)
    cf.append(GameGroup, group)
    cf.append(Latest, latest)
    cf.append(LatestArea, latest_area)
    cf.append(ServerHost, s_host)
    cf.append(IntranetHost, i_host)
    cf.append(GameStartTime, game_start)

    # db data
    d_host = ui.dataHost.text()
    d_port = ui.dataPort.text()
    d_user = ui.dataUser.text()
    d_pwd = ui.dataPasswd.text()

    cf.append(DataHost, d_host)
    cf.append(DataPort, d_port)
    cf.append(DataUser, d_user)
    cf.append(DataPasswd, d_pwd)

    # db log
    log_host = ui.logHost.text()
    log_port = ui.logPort.text()
    log_user = ui.logUser.text()
    log_pwd = ui.logPasswd.text()

    cf.append(LogHost, log_host)
    cf.append(LogPort, log_port)
    cf.append(LogUser, log_user)
    cf.append(LogPasswd, log_pwd)

    gs_num = int(ui.gsNum.text())
    gate_num = int(ui.gateNum.text())
    wsgate_num = int(ui.wsGateNum.text())

    print('--' * 5, gs_num, gate_num, wsgate_num, '--' * 5)

    cf.append(GsNum, gs_num)
    cf.append(GateNum, gate_num)
    cf.append(WsGateNum, wsgate_num)

    cf.append(ServerID, region * 1000 + group)

    init_dir(cf)
    init_dbname(cf, region, group)
    init_svn(cf)
    runner = r.Runner(cf)
    runner.run()


def init():
    ss = time.localtime()
    ct = '%.4d-%.2d-%.2d' % (ss[0], ss[1], ss[2])
    ui.gameStartTime.setText(ct)
    w_ip = tool.get_ip_www()
    i_ip = tool.get_ip()
    if w_ip != '':
        ui.serverHost.setText(w_ip)
    if i_ip != '':
        ui.intraneHost.setText(i_ip)


if __name__ == "__main__":
    print('hello world')
    app = QApplication(sys.argv)
    tw = QMainWindow()
    ui = u.Ui_region()
    ui.setupUi(tw)
    init()
    ui.submit.clicked.connect(submit)
    tw.setWindowTitle("懒人工具")
    # tw.setWindowIcon(QIcon('tool.png'))
    tw.show()
    sys.exit(app.exec_())
