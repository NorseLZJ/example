

linux LINUX:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-w -s '
	ls -lhtr

windows WINDOWS:
	go build -ldflags='-w -s '
	ls -lhtr
	cp ginCmd.exe options.exe
