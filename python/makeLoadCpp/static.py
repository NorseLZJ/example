# class head
class_head = """
class %s
{
    DECLARE_SINGLETON(%s);
    
public:
    %s() {}
    bool Init();
    bool LoadConfig();
    
private:

protected:
    %s
};"""

# struct head

struct_head = """
struct %s
{%s
};
typedef std::map<int,%s> map%s;"""

h_file_head = """#ifndef %s 
#define %s 

#include "GlobeObj.h"

%s 

%s 

#define gp%sCtrl (%s::Instance())

#endif """

# type dict
type_dict = {
    'int':
        ['int', 'atoi'],

    'int32':
        ['int', 'atoi'],

    'int/key':
        ['int', 'atoi'],

    'int32/key':
        ['int', 'atoi'],

    'string':
        ['string', ''],

    'string/vec':
        ['vector<int>', 'ParseStringToVector'],

    'string/vecvec':
        ['vector<vector<int>>', 'ParseStringToVectorVector'],

    'string/map':
        ['map<int,int>', 'ParseStringToMap'],
}

# 大小写字符判断
upper_cases = {
    "A": True,
    "B": True,
    "C": True,
    "D": True,
    "E": True,
    "F": True,
    "G": True,
    "H": True,
    "I": True,
    "J": True,
    "K": True,
    "L": True,
    "M": True,
    "N": True,
    "O": True,
    "P": True,
    "Q": True,
    "R": True,
    "S": True,
    "T": True,
    "U": True,
    "V": True,
    "W": True,
    "X": True,
    "Y": True,
    "Z": True,
}

LoadStrFmt = """info.%s = elem->Attribute("%s");"""
LoadIntFmt = """info.%s = atoi(elem->Attribute("%s"));"""
LoadVecFmt = """ParseStringToVector(elem->Attribute("%s"), info.%s, '|');"""
LoadVecVecFmt = """ParseStringToVectorVector(elem->Attribute("%s"), info.%s, "|;");"""
LoadMapFmt = """ParseStringToMap(elem->Attribute("%s"), info.%s, "|;");"""

# cpp
# include 单例　init函数
cppHead = """#include "%sCtrl.h"
#include "tinyxml.h"

IMPLEMENT_SINGLETON(%s);

bool %s::Init()
{
    if(!LoadConfig())
    {
        return false;
    } 
    return true;
}
"""

loadConfFunc = """
bool %s::LoadConfig()
{
    string strXMLFN = "db/%sDB.xml";
    string xmldbname = "%sDB";
    string xmlelem = "%s";
    string xmlelems = "%ss";
    
    char cf[256];
    snprintf(cf, sizeof(cf), "%%s%%s", gpLoadConfig->m_strEnvirDir.c_str(), strXMLFN.c_str());
    TiXmlDocument doc(cf);
    bool retVal = doc.LoadFile();

    if(!retVal)
    {
        CLog::ErrorLog("%%s file open error!", strXMLFN.c_str());
        return false;
    }

    TiXmlNode* ndRoot = doc.FirstChild(xmldbname.c_str());

    if(!ndRoot)
        return false;

    TiXmlNode* ndItems = ndRoot->FirstChild(xmlelems.c_str());

    if(!ndItems)
    {
        return false;
    }

    %s.clear();
    
    TiXmlNode* ndPosTemp = ndItems->FirstChild(xmlelem.c_str());
    for( ; ndPosTemp != NULL; ndPosTemp = ndPosTemp->NextSibling(xmlelem.c_str()))
    {
        TiXmlElement* elem = ndPosTemp->ToElement();

        if(!elem)
        {
            CLog::ErrorLog("%%s file get Element error!", strXMLFN.c_str());
            return false;
        }

        %sConf info;
        %s
        
        %s[info.%s] = info;
    }

    return true;
}
"""
