:: 生成服务脚本

C: && cd C:\srvany && instsrv.exe FSV C:\srvany\srvany.exe
REG ADD HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\FSV\Parameters
REG ADD HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\FSV\Parameters /v AppDirectory /t REG_SZ /d I:\Cache\go\bin\
REG ADD HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\FSV\Parameters /v Application /t REG_SZ /d I:\Cache\go\bin\fsv.exe

pause
