#ifndef __CALC_H_
#define __CALC_H_

#include <iostream>
#include <vector>
#include <map>

class Calc
{
private:
  int initVal, year;
  std::vector<float> rate;
  std::map<int, std::vector<int>> ret;

public:
  Calc(int _initVal, int _year, std::vector<float> &_rate);
  void run();
  virtual ~Calc();
  void print();
};

#endif
