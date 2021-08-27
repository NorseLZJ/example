#ifndef __ENV_H__
#define __ENV_H__

#include "signle.h"
#include <string>
#include <iostream>
using namespace std;
#include <stdio.h>

namespace norse
{

	class Env
	{
	public:
		Env();
		~Env();
		bool init();
		void set_info(
			const std::string &progeam,
			const std::string &exe,
			const std::string &cwd)
		{
			m_program = progeam;
			m_exe = exe;
			m_cwd = cwd;
		}
		void out_info()
		{
			printf("progeam:%s exe:%s cwd:%s \n", m_program.c_str(), m_exe.c_str(), m_cwd.c_str());
		}

	private:
		std::string m_program;
		std::string m_exe;
		std::string m_cwd;
	};

	typedef norse::Singleton<Env> EnvMgr;
}

#endif