CFLAGS = -g -Wall -lpthread -L/usr/lib64/mysql -lmysqlclient
CC = g++

obj = main.cpp \
	src/people.cpp \
	src/calc.cpp \
	src/mysql_connection_pool.cpp \

run main:
	$(CC) $(CFLAGS) $(obj) -o main 

clear _clean:
	rm main

