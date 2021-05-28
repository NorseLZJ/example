#include <iostream>
#include <string>
#include <vector>
#include <sstream>
#include <fstream>
#include <io.h>
#include <stdio.h>
#include <stdlib.h>
#include <windows.h>

// extend lib
#include "json.hpp"

#define placement "Placement\\"

using namespace std;
using json = nlohmann::json;


// global
string cwd;
bool noPlace = false;
int x = 0;
int y = 0;


// func
void dir(string path, vector<string> &ret);

void init_env(string cwd);

string get_cwd();

void rep_str(string path, vector<string> &ret);

string get_file_str(string file);

void write_file_str(string file, string data);

void get_argv(int argc, char **argv);

int get_val_xy(string param);

void get_val_xy(string src, int &x, int &y);

vector<string> string_split(std::string src, char *delim);


int main(int argc, char **argv) {

    get_argv(argc, argv);
    vector<string> ret;
    cwd = get_cwd();
    init_env(cwd);
    dir(cwd, ret);
    if (ret.size() <= 0) {
        printf("Files is NULL");
        exit(1);
    }
    rep_str(cwd, ret);

    std::cout << "Write all complete" << std::endl;
    return 0;
}

void dir(string path, vector<string> &ret) {
    long hFile = 0;
    struct _finddata_t fileInfo;
    string pathName, exdName;

    // \\* 代表要遍历所有的类型,如改成\\*.jpg表示遍历jpg类型文件
    if ((hFile = _findfirst(pathName.assign(path).append("\\*.json").c_str(), &fileInfo)) == -1) {
        return;
    }
    do {
        //cout << fileInfo.name << (fileInfo.attrib & _A_SUBDIR ? "[folder]" : "[file]") << endl;
        ret.push_back(fileInfo.name);
    } while (_findnext(hFile, &fileInfo) == 0);
    _findclose(hFile);
    return;
}

string get_cwd() {
    char buf[1000];
    GetCurrentDirectory(1024, buf);
    return string(buf);
}

void rep_str(string path, vector<string> &ret) {
    for (auto v : ret) {
        string v_src = v;
        v = path + "\\" + v;
        cout << v << endl;

        string str = get_file_str(v);

        string tag = "\"mc\":{";
        string src_str = str; //
        size_t idx = str.find(tag);
        string dest = str.substr(idx);
        dest = dest.substr(tag.size() + 1);
        idx = dest.find("{");
        dest = dest.substr(idx);
        dest = dest.substr(0, dest.size() - 2);

        string rep_str = dest;

        auto j3 = json::parse(dest);
        for (int i = 0; i < j3["frames"].size(); i++) {

            int old_x = j3["frames"][i]["x"];
            int old_y = j3["frames"][i]["y"];

            if (!noPlace) {
                string res_txt = string(cwd) + "\\" + placement + string(j3["frames"][i]["res"]);
                string str_txt = get_file_str(res_txt);
                if (str_txt == "")
                    continue;
                int tx, ty = 0;
                get_val_xy(str_txt, tx, ty);
                j3["frames"][i]["x"] = tx + x;
                j3["frames"][i]["y"] = ty + y;
            } else {
                j3["frames"][i]["x"] = old_x + x;
                j3["frames"][i]["y"] = old_y + y;
            }
        }
        string s_new = j3.dump();
        size_t pos = str.find(rep_str);
        if (pos == string::npos) continue;

        str.replace(pos, rep_str.size(), s_new);

        //string w_path = path + "\\out\\new_" + v_src;
        string w_path = path + "\\out\\" + v_src;
        write_file_str(w_path, str);
    }
}

string get_file_str(string file) {
    string str;
    ifstream fin;
    fin.open(file, ios::in);
    stringstream buf;
    buf << fin.rdbuf();
    str = buf.str();
    str.erase(std::remove(str.begin(), str.end(), '\n'), str.end());
    str.erase(std::remove(str.begin(), str.end(), ' '), str.end());
    fin.close();
    return str;
}

void get_argv(int argc, char **argv) {

    string cur = "";
    string param = "";
    for (int i = 1; i < argc; i++) {
        cur = argv[i];
        i++;

        if (i >= argc) {
            return;
        }

        param = argv[i];
        if (cur == "-noPlace" && param == "true") {
            noPlace = true;
        }
        if (cur == "-x") {
            x = get_val_xy(param);
        }
        if (cur == "-y") {
            y = get_val_xy(param);
        }
    }
}

int get_val_xy(string param) {
    int val = atoi(param.substr(1, param.size()).c_str());
    string flag = param.substr(0, 1);
    if (flag == "+") {
        return val;
    }
    return val -= val * 2;
}

void get_val_xy(string src, int &x, int &y) {
    if (src == "") return;

    vector<string> ret = string_split(src, ",");
    if (ret.size() < 2) return;
    x = atoi(ret[0].c_str());
    y = atoi(ret[1].c_str());
}

vector<string> string_split(std::string src, char *delim) {

    string::size_type pos1, pos2;
    vector<std::string> ret;
    pos2 = 0;
    while (pos2 != string::npos) {
        pos1 = src.find_first_not_of(delim, pos2);
        if (pos1 == string::npos)
            break;
        pos2 = src.find_first_of(delim, pos1 + 1);
        if (pos2 == string::npos) {
            if (pos1 != src.size())
                ret.push_back(src.substr(pos1));
            break;
        }
        ret.push_back(src.substr(pos1, pos2 - pos1));
    }
    return ret;
}

void init_env(string cwd) {
    string folder = cwd + "\\out";
    fstream _file;
    if ((_access(folder.c_str(), 0)) == -1) {
        string cmd = "mkdir " + folder;
        system(cmd.c_str());
    }
}

void write_file_str(string file, string data) {
    ofstream out;
    out.open(file, ios_base::out | ios_base::trunc);
    out << data << endl;
}
