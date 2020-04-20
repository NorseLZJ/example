#include "people.h"

People::People(std::string _name, int _age) {
  name = _name;
  age = _age;
}
People::~People() {}
void People::setname(std::string _name) { name = _name; }
void People::setage(int _age) { age = _age; }
std::string People::getname() { return name; }
int People::getage() { return age; }
