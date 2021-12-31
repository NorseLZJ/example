"""
new services
结合　instsrv.exe srvany.exe 创建服务的脚本
"""
import os
import sys
import json
from static import *

g_path = os.getcwd()  # 全局路径,从这里开始

# 所有的服务都在这里了　用来生成start,stop bat
services = []
startBat = ''
stopBat = ''


def start_and_dir():
    d = ""
    configFile = sys.argv[1]
    with open(configFile, 'r') as f:
        data = f.read()
        val = json.loads(data)
        d = val['srvnay']
    if d != "":
        global g_path
        g_path = d

    s1 = g_path.split(':')[0] + ":"
    c_dir1 = g_path.split(':')[1]
    c_dir1 = format("%s%s" % (s1, c_dir1))
    return s1, c_dir1


def gen_bat():
    configFile = sys.argv[1]
    outName = str(configFile).split(".")[0] + ".bat"
    with open(outName, 'w') as f:
        cent = format("%s\n\n%s" % (startBat, stopBat))
        f.write(cent)


def start():
    configFile = sys.argv[1]
    s1, d = start_and_dir()
    global startBat, stopBat
    with open(configFile, 'r') as f:
        data = f.read()
        val = json.loads(data)
        val = val['configs']
        for v in val:
            c_name = v['name']
            c_dir = v['dir']
            c_exe = v['exe']

            services.append(c_name)
            skip = v['skip']
            if skip == "1":
                continue

            # 创建服务
            cmd_1 = format(cd % (s1, d, c_name, d))
            print(cmd_1)
            os.system(cmd_1)

            cmd_2 = format(add_param % c_name)
            print(cmd_2)
            os.system(cmd_2)

            cmd_3 = format(add_dir % (c_name, c_dir))
            print(cmd_3)
            os.system(cmd_3)

            cmd_4 = format(add_exe % (c_name, c_exe))
            print(cmd_4)
            os.system(cmd_4)

            startBat = format("%s\nnet start %s" % (startBat, c_name))
            stopBat = format("%s\nsc stop %s" % (stopBat, c_name))


if __name__ == '__main__':
    if len(sys.argv) <= 1:
        print('没有给配置文件')
        exit(1)
    start()
    gen_bat()
