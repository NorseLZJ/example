#include <iostream>
#include "src/people.h"
#include "src/calc.h"
#include "src/m_mysql.h"

int main()
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
  M_Mysql sql;
  int ret = sql.initMysqlConnPool("192.168.20.177", 6380, "lzj", "123456", "cpp_test");
  if (ret != INIT_SUCCESS)
  {
    return ret;
  }
  return 0;
}
