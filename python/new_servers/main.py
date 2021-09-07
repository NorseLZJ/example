"""
new services
结合　instsrv.exe srvany.exe 创建服务的脚本
"""
import os
import json

"""
bat 
"""

cd = """
%s && cd %s\srvany && instsrv.exe %s %s\srvany\srvany.exe
"""

add_param = """
REG ADD HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\%s\Parameters
"""

add_dir = """
REG ADD HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\%s\Parameters /v AppDirectory /t REG_SZ /d %s 
"""

add_exe = """
REG ADD HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\%s\Parameters /v Application /t REG_SZ /d %s 
"""

g_path = os.getcwd()  # 全局路径,从这里开始


def start_and_dir():
    d = ""
    with open('config.json', 'r') as f:
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


def start():
    s1, d = start_and_dir()
    with open('config.json', 'r') as f:
        data = f.read()
        val = json.loads(data)
        val = val['configs']
        for v in val:
            c_name = v['name']
            c_dir = v['dir']
            c_exe = v['exe']

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


if __name__ == '__main__':
    start()
