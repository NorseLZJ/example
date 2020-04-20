#ifndef __CALC_H_
#define __CALC_H_

#include <iostream>

class Calc
{
private:
  int baseMoney, maxYear;
  float baseRate, addRate;

public:
  Calc(int _baseMoney, int _maxYear, float _baseRate, float _addRate);
  void run();
  virtual ~Calc();
  void print();
};

#endif
