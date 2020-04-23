#include "std.h"
char *toChar(std::string str)
{
    char *cstr = new char[str.length() + 1];
    std::strcpy(cstr, str.c_str());
    return cstr;
}