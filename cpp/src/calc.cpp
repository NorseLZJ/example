#include "calc.h"

Calc::Calc(int _initVal, int _year, std::vector<float> &_rate)
{
  initVal = _initVal;
  year = _year;
  rate = _rate;
}
void Calc::run()
{

  int val = 0;
  for (std::vector<float>::iterator it = rate.begin(); it != rate.end(); ++it)
  {
    val = initVal;
    for (int cc = 1; cc <= year; ++cc)
    {
      val = val + (val * (*it));
      ret[cc].push_back(val);
    }
  }
}
void Calc::print()
{
  for (int cc = 1; cc <= year; ++cc)
  {
    std::vector<int> group = ret[cc];
    for (std::vector<int>::iterator it = group.begin(); it != group.end(); ++it)
    {
      std::cout << *it << "\t\t";
    }
    std::cout << std::endl;
  }
}

Calc::~Calc() {}