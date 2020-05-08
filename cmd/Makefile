export dir=$(dir $(realpath $(firstword $(MAKEFILE_LIST))))

fsv _fsv:
	cd $(dir)fileserver
	go build -ldflags='-w -s' && go install
gcode _gcode:
	cd $(dir)getcode
	go build -ldflags='-w -s' && go install
interip _interip:
	cd $(dir)internalIp
	go build -ldflags='-w -s' && go install
mip _mip:
	cd $(dir)mip
	go build -ldflags='-w -s' && go install
jjson _jjson:
	cd $(dir)jjson
	go build -ldflags='-w -s' && go install
pping _pping:
	cd $(dir)pping
	go build -ldflags='-w -s' && go install
