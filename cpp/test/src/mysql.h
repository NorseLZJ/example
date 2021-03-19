#ifndef __MYSQL_H_
#define __MYSQL_H_

#include <iostream>
//#include <vector>
//#include <map>
#include <mysql/mysql.h>

class M_MYSQL
{
private:
  //int initVal, year;
  //std::vector<float> rate;
  //std::map<int, std::vector<int>> ret;

public:
  M_MYSQL();
  //M_MYSQL(int _initVal, int _year, std::vector<float> &_rate);
  //void run();
  virtual ~M_MYSQL();
  //void print();
};

#endif
