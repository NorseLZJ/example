#include <iostream>
#include "env.h"
using namespace std;

void test()
{
	norse::EnvMgr::GetInstance()->out_info();
}

int main(int argc, char **argv)
{

	norse::EnvMgr::GetInstance()->init();
	norse::EnvMgr::GetInstance()->set_info("1", "2", "3");
	norse::EnvMgr::GetInstance()->out_info();
	test();
	cout << "norse" << endl;
	return 0;
}