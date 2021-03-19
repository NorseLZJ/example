#include "people.h"
#include "calc.h"
#include "test.h"
#include "mysql_connection_pool.h"
#include "std.h"

void test(int argc, char **argv)
{
  //testPeople();
  //testCalc();
  //testMySQL();
  testRedis(argc, argv);
}

char *toChar(std::string str)
{
  char *cstr = new char[str.length() + 1];
  std::strcpy(cstr, str.c_str());
  return cstr;
}

void testPeople()
{
  // people
  People p("lzj", 11);
  p.print();
  p.setname("lzj");
  p.setage(13);
  p.print();
}

void testCalc()
{
  // calc
  int baseMoney = 10000, baseYear = 10;
  float baseRate = 0.1;
  std::vector<float> rate;
  for (int val = 0; val < 10; ++val)
  {
    rate.push_back(baseRate);
    baseRate += 0.1;
  }
  Calc c(baseMoney, baseYear, rate);
  c.run();
  c.print();
}

void testMySQL()
{
  // mysql
  std::vector<std::string> sqlS;
  sqlS.push_back(" \
  CREATE TABLE `ban` ( \
  `userId` bigint NOT NULL COMMENT 'id', \
  `openId` varchar(60) NOT NULL DEFAULT '' COMMENT 'openId', \
  PRIMARY KEY(`userId`))ENGINE = InnoDB DEFAULT CHARSET = utf8;");
  sqlS.push_back("INSERT into ban value (1,'lzj1');");
  sqlS.push_back("INSERT into ban value (2,'lzj2');");
  sqlS.push_back("INSERT into ban value (3,'lzj3');");
  sqlS.push_back("INSERT into ban value (4,'lzj4');");
  sqlS.push_back("INSERT into ban value (5,'lzj5');");
  sqlS.push_back("SELECT * from ban;");

  MysqlConnectionPool *pool = new MysqlConnectionPool();
  int ret = 0;
  // cpp_test 手动创建
  ret = pool->initMysqlConnPool("127.0.0.1", 3307, "root", "123456", "cpp_test");
  assert(ret == 0);
  ret = pool->openConnPool(10);
  assert(ret == 0);
  MYSQL_RES *res_ptr;
  int i, j;
  MYSQL_ROW sqlrow;

  for (std::vector<std::string>::iterator it = sqlS.begin(); it != sqlS.end(); ++it)
  {
    char *cstr = toChar(*it);
    mysqlConnection *mysqlConn = pool->fetchConnection();
    assert(mysqlConn != NULL);
    assert(cstr != NULL);
    pool->executeSql(mysqlConn, cstr);
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
    delete[] cstr;
  }
  delete pool;
}

void testRedis(int argc, char **argv)
{
  redisContext *c;
  redisReply *reply;
  char host[] = "172.17.0.3";
  int port = 6379;
  c = redisConnect(host, port);
  if (c == NULL)
  {
    printf("Connection is NULL\n");
    exit(1);
  }
  if (c->err)
  {
    printf("Connection is err : %s\n", c->errstr);
    printf("Connection is err : %d\n", c->err);
    exit(1);
  }

  /* PING server */
  reply = (redisReply *)redisCommand(c, "PING");
  printf("PING: %s\n", reply->str);
  freeReplyObject(reply);

  /* Set a key */
  reply = (redisReply *)redisCommand(c, "SET %s %s", "foo", "hello world");
  printf("SET: %s\n", reply->str);
  freeReplyObject(reply);

  /* Set a key using binary safe API */
  reply = (redisReply *)redisCommand(c, "SET %b %b", "bar", (size_t)3, "hello", (size_t)5);
  printf("SET (binary API): %s\n", reply->str);
  freeReplyObject(reply);

  /* Try a GET and two INCR */
  reply = (redisReply *)redisCommand(c, "GET foo");
  printf("GET foo: %s\n", reply->str);
  freeReplyObject(reply);

  reply = (redisReply *)redisCommand(c, "INCR counter");
  printf("INCR counter: %lld\n", reply->integer);
  freeReplyObject(reply);
  /* again ... */
  reply = (redisReply *)redisCommand(c, "INCR counter");
  printf("INCR counter: %lld\n", reply->integer);
  freeReplyObject(reply);

  /* Create a list of numbers, from 0 to 9 */
  reply = (redisReply *)redisCommand(c, "DEL mylist");
  freeReplyObject(reply);
  /* Disconnects and frees the context */
  redisFree(c);
}