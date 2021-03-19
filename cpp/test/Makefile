CFLAGS = -g -Wall -std=c++2a
CC = g++
MYSQLFLAGS = -L/usr/lib64/mysql -lmysqlclient 
#REDISFLAGS =  -L/usr/local/lib -lhiredis
REDISFLAGS = libhiredis.a

OBJ = main.cpp \
	src/people.cpp \
	src/calc.cpp \
	src/mysql_connection_pool.cpp \
	src/test.cpp \

test TEST:
	$(CC) $(CFLAGS) $(MYSQLFLAGS) $(OBJ) $(REDISFLAGS) -o main 

bitree BITREE:
	gcc bitree.c -o main -g -Wall

thread THREAD:
	$(CC) thread.cpp -o main -lpthread $(CFLAGS) 

clear CLEAN:
	rm main 
