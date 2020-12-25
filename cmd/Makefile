export dir=$(dir $(realpath $(firstword $(MAKEFILE_LIST))))

RELEASE=-ldflags='-w -s ' -x
CC=go install 

fsv _FSV:
	$(CC) $(RELEASE) $(dir)fsv

fjson _FJSON:
	$(CC) $(RELEASE) $(dir)jjson 

rpt _RPT:
	$(CC) $(RELEASE) $(dir)reptile

mip _MIP:
	$(CC) $(RELEASE) $(dir)mip

pp _PP:
	$(CC) $(RELEASE) $(dir)pping

interip _INTERIP:
	$(CC) $(RELEASE) $(dir)internalip

vsfmt _VSFMT:
	$(CC) $(RELEASE) $(dir)vsfmt

goget _GOGET:
	$(CC) $(RELEASE) $(dir)getcode
