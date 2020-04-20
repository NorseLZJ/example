#include <iostream>
#include "src/people.h" 

int main() {
  People p("lzj",11);
  p.print();
  p.setname("lzj");
  p.setage(13);
  p.print();

  return 0;
}
