export dir=$(dir $(realpath $(firstword $(MAKEFILE_LIST))))

FLAGS=-ldflags='-w -s '
CC=go build
BIN=${GOBIN}

fsv _fsv:
	$(CC) $(FLAGS) -o $(BIN)/fileserver.exe ${dir}fileserver/main.go


getCode _getCode:
	$(CC) $(FLAGS) -o $(BIN)/getcode.exe ${dir}getcode/main.go


internalIp _internalIp:
	$(CC) $(FLAGS) -o $(BIN)/internalip.exe ${dir}internalip/main.go


mip _mip:
	$(CC) $(FLAGS) -o $(BIN)/mip.exe ${dir}mip/main.go


jJson _jJson:
	$(CC) $(FLAGS) -o $(BIN)/jjson.exe ${dir}jjson/main.go


pPing _pPing:
	$(CC) $(FLAGS) -o $(BIN)/pping.exe ${dir}pping/main.go
