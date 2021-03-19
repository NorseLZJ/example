#include "netcmd.pb.h"
#include <iostream>
#include <string>

using namespace std;

void read_billboard(string msg);
void read_enterGsNotify(string msg);

int main(int argc, char **argv)
{
    netcmd::Billboard vv;
    vv.set_msg("hello world");

    string ret = vv.SerializeAsString();
    //cout << ret << endl;
    read_billboard(ret);

    netcmd::EnterGsNotify bb;
    bb.set_servertime("1.1.1.1");
    bb.set_serverversion(1);
    bb.set_serveropentime("2020-02-12:12:45");
    bb.set_clientfuncswitch(2);

    ret.clear();
    ret = bb.SerializeAsString();
    //cout << ret << endl;

    read_enterGsNotify(ret);
    return 0;
}

void read_billboard(string msg)
{
    netcmd::Billboard vv;
    if (!vv.ParseFromString(msg))
    {
        cerr << "Failed to parse billboard." << endl;
        return;
    }
    cout << vv.msg() << endl;
}

void read_enterGsNotify(string msg)
{
    netcmd::EnterGsNotify vv;
    if (!vv.ParseFromString(msg))
    {
        cerr << "Failed to parse billboard." << endl;
        return;
    }
    cout << vv.servertime() << endl;
    cout << vv.serverversion() << endl;
    cout << vv.serveropentime() << endl;
    cout << vv.clientfuncswitch() << endl;
}