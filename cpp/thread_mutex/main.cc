#include <iostream>
#include <memory>
#include <thread>
#include <vector>
#include <string>
#include <mutex>
#include <stdio.h>

using namespace std;

void out_str(string name);
bool get_str_by_mutex(string &ret);

vector<string> all_str;
std::mutex mtx;

int main()
{
    for (int i = 0; i < 100; i++)
    {
        all_str.push_back(to_string(i));
    }

    std::thread first(out_str, "first");
    std::thread second(out_str, "second");
    std::thread three(out_str, "three");
    std::thread four(out_str, "four");
    std::thread five(out_str, "five");

    first.join();
    second.join();
    three.join();
    four.join();
    five.join();

    return 0;
}

void out_str(string name)
{
    string ret;
    while (get_str_by_mutex(ret))
    {
        printf("tn:%s -> out:%s\n", name.c_str(), ret.c_str());
        //cout << "tn:" << name << "-> out:" << ret << endl;
    }
}

bool get_str_by_mutex(string &ret)
{
    mtx.lock();
    if (all_str.size() == 0)
    {
        mtx.unlock();
        return false;
    }

    //ret = all_str[all_str.size() - 1];
    //all_str.pop_back();
    ret = all_str[0];
    all_str.erase(all_str.begin());

    mtx.unlock();

    return true;
}
