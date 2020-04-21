#ifndef __M_MYSQL_H_
#define __M_MYSQL_H_

#include <mysql/mysql.h>
#include <semaphore.h>
#include <pthread.h>
#include <deque>
#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>
#include <time.h>
#include <cstring>
#include <errno.h>
#include <assert.h>

#define INIT_ERROR 1
#define INIT_SUCCESS 0

#define MAX_SQL_LEN 4096
#define MAX_SETTING_STRING_LEN 256
#define MAX_POOL_SIZE 256
#define MCP_LOG(fmt, ...)                \
  {                                      \
    fprintf(stderr, fmt, ##__VA_ARGS__); \
  }

typedef struct
{
  char host[MAX_SETTING_STRING_LEN];
  int port;
  char database[MAX_SETTING_STRING_LEN];
  char user[MAX_SETTING_STRING_LEN];
  char password[MAX_SETTING_STRING_LEN];
  char charset[MAX_SETTING_STRING_LEN];
  unsigned int timeout;
} setting;

typedef struct
{
  setting *connSetting;
  MYSQL *sock;
  MYSQL mysql;
  int res;
} mysqlConn;

class M_Mysql
{
public:
  int initMysqlConnPool(const char *host, int port, const char *user, const char *password, const char *database);
  void recycleConnection(mysqlConn *conn);
  int executeSql(mysqlConn *conn, const char *sql);
  mysqlConn *fetchConnection();
  int openConnPool(int coreConnNum);
  void setCharsetOption(setting *connSetting, const char *charset);

public:
  M_Mysql()
  {
  }
  ~M_Mysql()
  {
    closeConnPool();
  }

private:
  int lockPool();
  void closeConnPool();

private:
  int connNum;
  sem_t sem;
  std::deque<mysqlConn *> connPool;
  setting *connSetting;
  pthread_mutex_t mutex;
};

#endif