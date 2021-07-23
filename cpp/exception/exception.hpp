#ifndef __EXCEPTION_H_
#define __EXCEPTION_H_

#include <exception>
#include <string>
#include <memory>
#include <stdlib.h>

const char *logfmt = "file:%s:%d err:%s";

class Exception1 : public std::exception
{

public:
	Exception1(const char *file, const char *err, unsigned int line)
	{
		this->file = file;
		this->err = err;
		this->line = line;
		ret = (char *)malloc(8192);
		std::exception();
	};
	~Exception1()
	{
		printf("destory\n");
		free(ret);
		ret = NULL;
		ret = NULL;
	};

public:
	virtual const char *what() const noexcept override
	{
		sprintf(ret, logfmt, file, line, err);
		//printf("err:%s\n", ret);
		return ret;
	}

private:
	const char *file;
	const char *err;
	unsigned int line;
	char *ret;
};

// some word if you need
class Exception2 : public Exception1
{
public:
	Exception2(const char *file, const char *err, unsigned int line) : Exception1(file, err, line){};
	~Exception2(){};
};

#endif