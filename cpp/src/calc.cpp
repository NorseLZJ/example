#include "calc.h"

Calc::Calc(int _baseMoney = 10000, int _maxYear = 10, float _baseRate = 0.1,
           float _addRate = 0.1)
{
  baseMoney = _baseMoney;
  maxYear = _maxYear;
  baseRate = _baseRate;
  addRate = _addRate;
}
void Calc::run() { std::cout << "runing..." << std::endl; }
Calc::~Calc() { std::cout << "drop ..." << std::endl; }
void Calc::print()
{
  std::cout << "baseMoney:" << baseMoney << "maxYear:" << maxYear
            << "baseRate:" << baseRate << "addRate:" << addRate << std::endl;
}
