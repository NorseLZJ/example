export dir=$(dir $(realpath $(firstword $(MAKEFILE_LIST))))

FLAGS=-ldflags='-w -s '
CC=go build
BIN=${GOBIN}

fsv _fsv:
	$(CC) $(FLAGS) -o $(BIN)/fileserver ${dir}fileserver/main.go
getCode _getCode:
	$(CC) $(FLAGS) -o $(BIN)/getcode ${dir}getcode/main.go
internalIp _internalIp:
	$(CC) $(FLAGS) -o $(BIN)/internalip ${dir}internalip/main.go
mIp _mIp:
	$(CC) $(FLAGS) -o $(BIN)/mip ${dir}mip/main.go
jJson _jJson:
	$(CC) $(FLAGS) -o $(BIN)/jjson ${dir}jjson/main.go
pPing _pPing:
	$(CC) $(FLAGS) -o $(BIN)/pping ${dir}pping/main.go
