#include <iostream>
#include <memory>
#include "exception.hpp"

int main()
{
    try
    {
        throw Exception2(__FILE__, "---> throw err ... <---", __LINE__);
    }
    //catch (const Exception1 &e)
    //{
    //    std::cerr << e.what() << '\n';
    //}
    catch (const Exception2 &e)
    {
        std::cerr << e.what() << '\n';
    }

    return 0;
}