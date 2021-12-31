"""
static.py
一些固定的字符串拼接
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
