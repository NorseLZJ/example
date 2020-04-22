CFLAGS = -Wall -lpthread -L/usr/lib64/mysql -lmysqlclient
CC = g++

param = -g -Wall -lpthread
obj = main.cpp \
	src/people.cpp \
	src/calc.cpp \
	src/m_mysql.cpp \

run main:
	$(CC) $(CFLAGS) $(obj) -o main $(param) 

clear _clean:
	rm main

