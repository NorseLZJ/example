#include <iostream>
#include <assert.h>
#include "src/people.h"
#include "src/calc.h"
#include "src/mysql_connection_pool.h"

int main(int argc, char **argv)
{
  // people
  //People p("lzj",11);
  //p.print();
  //p.setname("lzj");
  //p.setage(13);
  //p.print();

  // calc
  //int baseMoney = 10000, baseYear = 10;
  //float baseRate = 0.1;
  //std::vector<float> rate;
  //for (int val = 0; val < 10; ++val)
  //{
  //  rate.push_back(baseRate);
  //  baseRate += 0.1;
  //}
  //Calc c(baseMoney, baseYear, rate);
  //c.run();
  //c.print();

  // mysql
  char str[] = "CREATE TABLE `ban` ( \
  `userId` bigint NOT NULL COMMENT 'id', \
  `openId` varchar(60) NOT NULL DEFAULT '' COMMENT 'openId', \
  PRIMARY KEY (`userId`) \
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8;";
  //char str[] = "CREATE TABLE `ban` ( \
  //`userId` bigint NOT NULL COMMENT 'id', \
  //`openId` varchar(60) NOT NULL DEFAULT '' COMMENT 'openId', \
  //PRIMARY KEY (`userId`) \
  //) ENGINE=InnoDB DEFAULT CHARSET=utf8;";

  MysqlConnectionPool *pool = new MysqlConnectionPool();
  int ret = 0;
  ret = pool->initMysqlConnPool("127.0.0.1", 3307, "root", "123456", "cpp_test");
  assert(ret == 0);
  ret = pool->openConnPool(10);
  assert(ret == 0);
  int num = 1;
  MYSQL_RES *res_ptr;
  int i, j;
  MYSQL_ROW sqlrow;
  while (num > 0)
  {
  mysqlConnection *mysqlConn = pool->fetchConnection();
  assert(mysqlConn != NULL);
  pool->executeSql(mysqlConn, str);
  res_ptr = mysql_store_result(mysqlConn->sock);
  if (res_ptr)
    {
      printf("%lu Rows\n", (unsigned long)mysql_num_rows(res_ptr));
      j = mysql_num_fields(res_ptr);
      while ((sqlrow = mysql_fetch_row(res_ptr)))
      {
        for (i = 0; i < j; i++)
          printf("%s\t", sqlrow[i]);
        printf("\n");
      }
      if (mysql_errno(mysqlConn->sock))
      {
        fprintf(stderr, "Retrive error:%s\n", mysql_error(mysqlConn->sock));
      }
    }
    mysql_free_result(res_ptr);
    pool->recycleConnection(mysqlConn);
    num--;
  }
  delete pool;

  return 0;
}
