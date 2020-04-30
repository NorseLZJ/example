CFLAGS = -g -Wall -lpthread -std=c++11 
CC = g++
MYSQLFLAGS = -L/usr/lib64/mysql -lmysqlclient 
#REDISFLAGS =  -L/usr/local/lib -lhiredis
REDISFLAGS = libhiredis.a

obj = main.cpp \
	src/people.cpp \
	src/calc.cpp \
	src/mysql_connection_pool.cpp \
	src/test.cpp \

run main:
	$(CC) $(CFLAGS) $(MYSQLFLAGS) $(obj) $(REDISFLAGS) -o main 

clear _clean:
	rm main

