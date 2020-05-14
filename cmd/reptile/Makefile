export dir=$(dir $(realpath $(firstword $(MAKEFILE_LIST))))

FLAGS=-ldflags='-w -s '
CC=go build
BIN=${GOBIN}

run _run:
	#$(CC) $(FLAGS) -o $(BIN)/reptile
	$(CC) $(FLAGS) -o reptile.exe
