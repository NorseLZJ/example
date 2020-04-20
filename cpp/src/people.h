#ifndef __PEOPLE_H_
#define __PEOPLE_H_

#include <iostream>

class People
{
private:
  int age;
  std::string name;

public:
  People(std::string _name, int _age);
  virtual ~People();
  void setname(std::string _name);
  void setage(int _age);
  std::string getname();
  int getage();
  void print() { std::cout << "name:" << name << "age:" << age << std::endl; }
};

#endif
