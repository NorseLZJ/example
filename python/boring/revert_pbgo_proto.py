import os

# 处理字段选项
def dispose_field_type(ft:str):
    ft = ft.replace("*","")
    if "[]" not in ft and ("int" in ft or "string" in ft):
        return ft
    if "[]" not in ft and "float" in ft:
        return "float" 
    if "[]" in ft:
        return f"repeated {ft.replace("[]","")}"
    
    return ft

# 将Go代码转换为Proto文件
def convert_go_to_proto(go_code,packnme):
    proto_code = f"""syntax="proto3";\npackage {packname};\n"""
    state=False
    for line in go_code.split("\n"):
        if "type " in line and "struct" in line:
            message_name = line.split(" ")[1]
            proto_code += f"message {message_name} {{\n"
            state=True
        elif "}" in line and state:
            proto_code += "}\n\n"
            state=False
        elif "protobuf" in line and "import" not in line and "github.com" not in line:
            count=0
            ss = line.split(" ")
            # print(ss)
            fieldName=''
            fieldType=''
            info=''

            for v in ss:
                if v == "":
                    continue
                if count ==0 :
                    fieldName=v
                if count == 1:
                    fieldType = v
                if count ==2:
                    info = v
                count+=1
            
            if info == "":
                print(fieldName,fieldType)
                exit(0)
                        

            fieldType=dispose_field_type(fieldType)
            number=info.split(",")[1]
            proto_code+= f"  {fieldType} {fieldName} = {number};\n"
            
    return proto_code


def check_directory(directory):
    out=[]
    for root, dirs, files in os.walk(directory):
        if 'msgtype' in root:
            for file in files:
                if 'pb.go' in file:
                    # print(os.path.join(root, file))
                    out.append(os.path.join(root, file))
    return out 


if __name__ == "__main__":
    files = check_directory('E:\\codes')
    for file in files:
        result = ''
        ss = file.split("\\")
        packname=ss[len(ss)-2]
        with open(file,"r",encoding="utf-8") as f:
            result  = convert_go_to_proto(f.read(),packname)

        file2 = file.replace("pb.go",'proto')
        print(file,' -> ',file2)
        with open(file2,"w",encoding="utf-8") as f:
            f.write(result)
        ss = file2.split("\\")
        file3=f"proto\\{ss[len(ss)-1]}"
        with open(file3,"w",encoding="utf-8") as f:
            f.write(result)
        
