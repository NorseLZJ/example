pub mod static_variable {
    use std::fs;
    use std::ops::Index;
    use calamine::{Reader, open_workbook, Xlsx};

    fn type_load_method(c_type: &str, field_name: &str) -> (String, String) {
        return match c_type {
            "int" | "int32" | "int/key" | "int32/key" => {
                let load = format!("\t\tinfo.{0} = atoi({1}({2}{0}{2}));", field_name, ELEM, CHAR_QUOTATION_MARKS);
                let new_type = String::from("int");
                (load, new_type)
            }
            "string/vec" => {
                let load = format!("\t\tinfo.{0} = {1}({2}({3}{0}{3}), info.{0}, '|');", field_name, PAR_STR_TO_VEC, ELEM, CHAR_QUOTATION_MARKS);
                let new_type = String::from("vector<int>");
                (load, new_type)
            }
            "string/vecvec" => {
                let load = format!("\t\tinfo.{0} = {1}({2}({3}{0}{3}), info.{0}, '|;');", field_name, PAR_STR_TO_VEC_VEC, ELEM, CHAR_QUOTATION_MARKS);
                let new_type = String::from("vector<vector<int>>");
                (load, new_type)
            }
            "string/map" => {
                let load = format!("\t\tinfo.{0} = {1}({2}({3}{0}{3}), info.{0}, '|;');", field_name, PAR_STR_TO_MAP, ELEM, CHAR_QUOTATION_MARKS);
                let new_type = String::from("map<int,int>");
                (load, new_type)
            }
            "string/map64" => {
                let load = format!("\t\tinfo.{0} = {1}({2}({3}{0}{3}), info.{0}, '|;');", field_name, PAR_STR_TO_MAP, ELEM, CHAR_QUOTATION_MARKS);
                let new_type = String::from("map<int,UINT64>");
                (load, new_type)
            }
            _ => {
                let load = String::from("");
                let new_type = String::from("");
                (load, new_type)
            }
        };
    }

    const BRACES_LEFT: &'static str = "{";
    const BRACES_RIGHT: &'static str = "}";
    const CHAR_QUOTATION_MARKS: &'static char = &'"';
    const ELEM: &'static str = "elem->Attribute";
    const PAR_STR_TO_VEC: &'static str = "ParseStringToVector";
    const PAR_STR_TO_VEC_VEC: &'static str = "ParseStringToVectorVector";
    const PAR_STR_TO_MAP: &'static str = "ParseStringToMap";

    pub struct CppConfig {
        path: String,
        name: String,
        conf_name: String,
        define_name: String,
        struct_str: String,
        load_struct_str: String,
        out_prefix: String,
        protected: String,
        m_val_type: String,
        m_val_name: String,
        class_ctrl_name: String,
        key: String,
    }


    impl CppConfig {
        pub fn new(path: &String) -> Self {
            CppConfig {
                path: String::from(path),
                name: String::from(""),
                conf_name: String::from(""),
                define_name: String::from(""),
                struct_str: String::from(""),
                load_struct_str: String::from(""),
                out_prefix: String::from(""),
                protected: String::from(""),
                m_val_type: String::from(""),
                m_val_name: String::from(""),
                class_ctrl_name: String::from(""),
                key: String::from(""),
            }
        }

        pub fn run(&mut self) {
            let mut workbook: Xlsx<_> = open_workbook(&self.path).expect("文件路径不对吧大哥 ! ");

            let s_str = workbook.sheet_names()[0].to_string();
            let s: Vec<_> = s_str.split("|").collect();
            self.some_setting(s[1]);

            let mut sff = String::from("");
            let mut lfs = String::from("");
            if let Some(Ok(r)) = workbook.worksheet_range_at(0) {
                let (explains, types, fields, register) = (r.index(0), r.index(1), r.index(2), r.index(3));

                for i in 0..r.index(3).len() {
                    if register[i].to_string().contains("server") {
                        let (v_name, v_type) = (fields[i].to_string(), types[i].to_string());
                        let explain = explains[i].to_string().replace("\r\n", "").replace("\n", "");
                        if v_name == "" || v_type == "" {
                            continue;
                        }
                        // add end suffix ';'
                        let c_name = format!("{}; ", v_name);

                        self.try_set_key(&v_type, &v_name);
                        let (l_str, c_type) = type_load_method(&v_type, &v_name);
                        if l_str != "" && c_type != "" {
                            //println!("{} \t\t {}", new_type, load);
                            lfs = format!("{}\n{}", lfs, l_str);
                            sff = format!("{}\n\t{:<w1$} {:<w$} //{:<w$}", sff, c_type, c_name, explain, w1 = 20, w = 30);
                        }
                    }
                }
                self.set_struct(&sff);
                self.set_load_struct(&lfs);
            }
            self.write_h();
            self.write_cpp();
        }

        /// init 根据名字初始化一些公共变量
        fn some_setting(&mut self, name: &str) {
            self.conf_name = format!("{}Conf", name);
            self.out_prefix = format!("{}Ctrl", name);
            self.name = String::from(name);
            self.m_val_type = format!("map{}", name);
            self.m_val_name = format!("m_map{}", name);
            self.protected = format!("{} {};", self.m_val_type, self.m_val_name);
            self.class_ctrl_name = format!("C{}Ctrl", name);

            self.set_define_name(name);
        }

        fn try_set_key(&mut self, c_type: &str, field_name: &str) {
            if c_type == "int/key" || c_type == "int32/key" {
                self.key = String::from(field_name);
            }
        }

        fn set_define_name(&mut self, name: &str) {
            let mut cur: String = String::from("");
            let mut g_idx = 0;
            let mut v;
            for i in name.to_string().bytes() {
                if i >= 97 {
                    v = (i - 32) as char;
                } else {
                    v = i as char;
                }

                if i < 97 && g_idx != 0 {
                    cur.push('_');
                    cur.push(v);
                } else {
                    cur.push(v);
                }
                g_idx += 1;
            }
            cur.push_str("_CTRL_H");
            //println!("{}", cur);
            self.define_name = cur;
        }

        fn set_struct(&mut self, fields: &String) {
            let cstr = format!("struct {} {} {} \n{};", self.conf_name, BRACES_LEFT, fields, BRACES_RIGHT).to_string();
            self.struct_str = cstr;
        }

        fn set_load_struct(&mut self, load_fields: &String) {
            self.load_struct_str = String::from(load_fields);
        }

        fn write_h(&self) {
            let cur = format!("{}.h", self.out_prefix);
            let typedef_str = format!("typedef std::map<int,{0}> map{1};", self.conf_name, self.name);
            let gp_def_str = format!("#define gp{}Ctrl ({}::Instance())", self.name, self.class_ctrl_name);
            let class_str = format!(
                "class {0} \
\n{1} \
\n    DECLARE_SINGLETON({0}); \
\n\n\
public: \
\n    {0}() {1}{2} \
\n    bool Init(); \
\n    bool LoadConfig(); \
\n\nprivate: \
\n\nprotected: \
\n    {3} \
\n{2};", self.class_ctrl_name, BRACES_LEFT, BRACES_RIGHT, self.protected);

            let str = format!("#ifndef {0} \
\n#define {0}\
\n\n\
#include {5}GlobeObj.h{5}
\n
{1}\
\n{2}\
\n\n\
{3}\
\n\n{4}\
\n\n#endif", self.define_name, self.struct_str, typedef_str, class_str, gp_def_str, CHAR_QUOTATION_MARKS);

            fs::write(cur, str).unwrap();
        }

        fn write_cpp(&self) {
            let cur = format!("{}.cpp", self.out_prefix);
            let include = format!("#include {1}{0}.h{1}\
\n#include {1}tinyxml.h{1}\
\n\n\
IMPLEMENT_SINGLETON({2});\
\n\
\nbool {2}::Init()\
\n{3}\
\n    if (!LoadConfig())\
\n    {3}\
\n        return false;\
\n    {4}\
\n    return true;\
\n{4}\
\n\n", self.out_prefix, CHAR_QUOTATION_MARKS, self.class_ctrl_name, BRACES_LEFT, BRACES_RIGHT);

            let load_fn = format!("\
bool {0}::LoadConfig() \
\n{2}\
\n    string strXMLFN = {4}db/{1}DB.xml{4};\
\n    string xmldbname = {4}{1}DB{4};\
\n    string xmlelem = {4}{1}{4};\
\n    string xmlelems = {4}{1}s{4};\
\n
\n    char cf[256]; \
\n    snprintf(cf, sizeof(cf), {4}%s%s{4}, gpLoadConfig->m_strEnvirDir.c_str(), strXMLFN.c_str()); \
\n    TiXmlDocument doc(cf); \
\n    bool retVal = doc.LoadFile(); \
\n \
\n    if(!retVal) \
\n    {2} \
\n        CLog::ErrorLog({4}%s file open error!{4}, strXMLFN.c_str()); \
\n        return false; \
\n    {3} \
\n \
\n    TiXmlNode* ndRoot = doc.FirstChild(xmldbname.c_str()); \
\n    if(!ndRoot) \
\n        return false; \
\n \
\n    TiXmlNode* ndItems = ndRoot->FirstChild(xmlelems.c_str()); \
\n    if(!ndItems) \
\n    {2} \
\n        return false; \
\n    {3} \
\n\n    {5}.clear();  \
\n    TiXmlNode* ndPosTemp = ndItems->FirstChild(xmlelem.c_str()); \
\n    for( ; ndPosTemp != NULL; ndPosTemp = ndPosTemp->NextSibling(xmlelem.c_str())) \
\n    {2} \
\n        TiXmlElement* elem = ndPosTemp->ToElement(); \
\n        if(!elem) \
\n        {2} \
\n            CLog::ErrorLog({4}%s file get Element error!{4}, strXMLFN.c_str()); \
\n            return false; \
\n        {3} \
\n        {8} info;
{6}
\n        {5}[info.{7} = info]
    {3}\n
    return true; \
\n{3}\n
", self.class_ctrl_name, self.name, BRACES_LEFT, BRACES_RIGHT, CHAR_QUOTATION_MARKS, self.m_val_name, self.load_struct_str, self.key, self.conf_name);

            let str = format!("{}\n{}", include, load_fn);
            fs::write(cur, str).unwrap();
        }
    }
}