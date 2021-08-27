#include <string>
#include "env.h"

namespace norse
{

	Env::Env()
	{
	}

	Env::~Env()
	{
	}

	bool Env::init()
	{
		cout << "Env init..." << endl;
		return true;
	}
}