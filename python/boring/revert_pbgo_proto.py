import os
import re


# 处理字段选项
def dispose_field_type(ft: str):
    ft = ft.replace("*", "")
    if "[]" not in ft and ("int" in ft or "string" in ft):
        return ft
    if "[]" not in ft and "float" in ft:
        return "float"
    if "[]" in ft:
        return f"repeated {ft.replace("[]","")}"
    return ft


# 将Go代码转换为Proto文件
def convert_go_to_proto(go_code, packnme):
    proto_code = f"""syntax="proto3";\npackage {packname};\n"""
    state = False
    for line in go_code.split("\n"):
        line = re.sub(r"\s+", " ", line)
        if "type " in line and "struct" in line:
            message_name = line.split(" ")[1]
            proto_code += f"message {message_name} {{\n"
            state = True
        elif "}" in line and state:
            proto_code += "}\n\n"
            state = False
        elif "protobuf" in line and "import" not in line and "github.com" not in line:
            ss = line.split(" ")
            (field_name, field_type, info) = ss[1], ss[2], ss[3]
            fieldType = dispose_field_type(field_type)
            number = info.split(",")[1]
            proto_code += f"  {fieldType} {field_name} = {number};\n"

    return proto_code


def check_directory(directory):
    out = []
    for root, dirs, files in os.walk(directory):
        if "msgtype" in root:
            for file in files:
                if "pb.go" in file:
                    # print(os.path.join(root, file))
                    out.append(os.path.join(root, file))

    return out


"""
给定目录，递归所有.pb.go 文件，然后转换成proto,会在pb.go 目录生成一份，还会在当前目录proto下生成一份,需要手动创建proto
"""
if __name__ == "__main__":
    files = check_directory("E:\\codes")
    for file in files:
        result = ""
        ss = file.split("\\")
        packname = ss[len(ss) - 2]
        with open(file, "r", encoding="utf-8") as f:
            result = convert_go_to_proto(f.read(), packname)

        file2 = file.replace("pb.go", "proto")
        print(file, " -> ", file2)
        with open(file2, "w", encoding="utf-8") as f:
            f.write(result)
        ss = file2.split("\\")
        file3 = f"proto\\{ss[len(ss)-1]}"
        with open(file3, "w", encoding="utf-8") as f:
            f.write(result)
