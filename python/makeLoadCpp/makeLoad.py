import sys
from xlrd3 import *
from xml.etree.ElementTree import ElementTree
from xml.etree.ElementTree import Element
from xml.etree.ElementTree import SubElement as SE

import static as s

"""
name 对应的输出
name.h name.cpp 
define name 
"""
name = ''
inputFile = ''

# BossChallenge -> CBossChallenge
cname = ''
# xml文件prefix
dbXml = ''

# key map<key(int),XXX>
MapKeyWord = ''

help_txt = """
Usage 
inputFile(xxx.xlsx) 

OUTPUT Example
-----------------------------------------------------------------------
BossChallenge out to 
BossChallenge.h
    gpBossChallenge
    
BossChallenge.cpp
-----------------------------------------------------------------------

Usage example
makeLoad.exe 挑战首领表_vh5.xlsx
"""

# 全局变量动态修改
# struct 结构
struct_s = ''

# 加载函数　读配置字段部分
loadConf_s = ''

# 受保护变量那一条
protected_s = ''

# 变量类型　变量名
m_val_type = ''
m_val_name = ''

# cpp 文件 头部
cppHeadCont = ''

# 加载配置的函数
cppLoadCont = ''


def read_argv():
    if len(sys.argv) < 2:
        print(help_txt)
        sys.exit(1)

    global inputFile
    inputFile = str(sys.argv[1])


def read_struct():
    t_book = open_workbook(inputFile)
    t_sheet = t_book.sheet_by_index(0)

    ss = str(t_sheet.name).split('|')
    if len(ss) < 2:
        print('配置名字格式错误 （中文|English）')
        exit(1)

    global dbXml, cname, name
    dbXml = str(ss[1])
    name = dbXml
    cname = format("C%sCtrl" % dbXml)

    explain = t_sheet.row_values(0)
    types = t_sheet.row_values(1)
    keys = t_sheet.row_values(2)
    server = t_sheet.row_values(3)

    fields = ""
    for idx in range(0, len(server)):
        if str(server[idx]).find('server') == -1:
            continue

        t_type = types[idx]
        t_name = keys[idx]
        cp_t_name = t_name
        t_name = t_name + ";"
        t_explain = explain[idx]
        t_explain = str(t_explain).replace("\n", "")
        t_explain = str(t_explain).replace("\r\n", "")

        set_map_key_word(t_type, cp_t_name)

        big_type = s.get_type_method(t_type)
        if big_type is None:
            print("UnKnow Type")
            exit(1)
        print(big_type)
        t_type = big_type[0]
        load_type = big_type[1]
        if idx == 0:
            fields = format("\t%-20s %-20s //%s" % (t_type, t_name, t_explain))
        else:
            fields = format("%s\n\t%-20s %-20s //%s" % (fields, t_type, t_name, t_explain))

        write_load_field(load_type, cp_t_name)

    conf_name = format("%sConf" % name)
    w_str = format(s.struct_head % (conf_name, fields, conf_name, name))

    global protected_s, m_val_name, m_val_type
    m_val_type = format("map%s" % name)
    m_val_name = format("m_map%s" % name)
    protected_s = format("%s %s;" % (m_val_type, m_val_name))
    # print(w_str)
    # print(load_fields)
    global struct_s
    struct_s = w_str

    write_cpp_head()
    write_cpp_load()


def write_load_field(load_type: str, cp_t_name: str):
    """
    加载的函数字format append
    :return:
    """
    global loadConf_s
    temp = load_field(load_type, cp_t_name)
    if temp != '' and loadConf_s == '':
        loadConf_s = format("%s" % temp)
    elif temp != '':
        loadConf_s = format("%s\n\t\t%s" % (loadConf_s, temp))


def set_map_key_word(t_type: str, key_word: str):
    if t_type == 'int/key' or t_type == 'int32/key':
        global MapKeyWord
        MapKeyWord = key_word


def write_cpp_head():
    temp = format(s.cppHead % (name, cname, cname))
    global cppHeadCont
    cppHeadCont = temp


def write_cpp_load():
    temp = format(s.loadConfFunc % (
        cname, dbXml, dbXml, dbXml, dbXml, m_val_name,
        name, loadConf_s, m_val_name, MapKeyWord))

    global cppLoadCont
    # print(temp)
    cppLoadCont = temp


def load_field(l_type: str, f_name: str):
    if l_type == '':  # string
        tmp = format(s.LoadStrFmt % (f_name, f_name))
        return tmp
    elif l_type == 'atoi':
        tmp = format(s.LoadIntFmt % (f_name, f_name))
        return tmp
    elif l_type == 'ParseStringToVector':
        tmp = format(s.LoadVecFmt % (f_name, f_name))
        return tmp
    elif l_type == 'ParseStringToVectorVector':
        tmp = format(s.LoadVecVecFmt % (f_name, f_name))
        return tmp
    elif l_type == 'ParseStringToMap':
        tmp = format(s.LoadMapFmt % (f_name, f_name))
        return tmp

    return ''


def class_ret():
    w_str = format(s.class_head % (cname, cname, cname, protected_s))
    return w_str


def write_h():
    temp_file = name + "Ctrl.h"
    with open(temp_file, 'w') as f:
        d_name = define(name)
        d_class = class_ret()
        w_str = format(s.h_file_head % (d_name, d_name, struct_s, d_class, name, cname))
        f.write(w_str)


def write_cpp():
    temp_file = name + "Ctrl.cpp"
    with open(temp_file, 'w') as f:
        w_str = format("%s%s" % (cppHeadCont, cppLoadCont))
        f.write(w_str)


def define(cur_str):
    cur_str = str(cur_str)
    case_s = []
    prev_idx = 0
    for i in range(0, len(cur_str)):
        if cur_str[i] in s.upper_cases:
            case_s.append(cur_str[prev_idx:i])
            prev_idx = i

    case_s.append(cur_str[prev_idx:])
    w_str = ''
    first = True
    for v in case_s:
        if v == '':
            continue
        if first:
            w_str = format("%s" % (v.upper()))
            first = False
            continue
        else:
            w_str = format("%s_%s" % (w_str, v.upper()))

    w_str = format("%s_CTRL_H" % w_str)
    return w_str


if __name__ == "__main__":
    read_argv()
    read_struct()
    write_h()
    write_cpp()
