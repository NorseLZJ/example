export dir=$(dir $(realpath $(firstword $(MAKEFILE_LIST))))

FLAGS=-ldflags='-w -s '
CC=go build 
OUTDIR=$(GOBIN)

fsv _fsv:
	$(CC) $(FLAGS) -o $(OUTDIR)/fsv $(dir)fileserver
	# fsv -s ~/std/

fjson _fjson:
	$(CC) $(FLAGS) -o $(OUTDIR)/fjson $(dir)jjson 
	# fjson -f xxx.json
	# fjson -d xxx(dir)

rpt _rpt:
	$(CC) $(FLAGS) -o $(OUTDIR)/rpt $(dir)reptile
	# rpt -f xxx.json

mip _mip:
	$(CC) $(FLAGS) -o $(OUTDIR)/mip $(dir)mip
	# mip

pp _pp:
	$(CC) $(FLAGS) -o $(OUTDIR)/pp $(dir)pping
	# pping -f xxx.json