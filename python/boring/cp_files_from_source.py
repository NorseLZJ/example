import os
import shutil


def copy_directory_contents(source_dir):
    # 获取当前脚本执行的文件夹路径
    current_dir = os.getcwd()

    # 构建目标文件夹路径
    target_dir = os.path.join(current_dir, os.path.basename(source_dir))

    # 创建目标文件夹
    os.makedirs(target_dir, exist_ok=True)

    # 遍历源文件夹中的所有文件和文件夹
    for item in os.listdir(source_dir):
        item_path = os.path.join(source_dir, item)
        target_path = os.path.join(target_dir, item)

        # 如果是文件夹，则递归调用复制函数
        if os.path.isdir(item_path):
            copy_directory_contents(item_path)
        # 如果是文件，则直接复制
        elif os.path.isfile(item_path):
            shutil.copy2(item_path, target_path)

    print("复制完成！")


# 指定源文件夹路径
# source_directory = "指定目录的路径"
source_directory = "C:\\code\\Github\\example\\python\\fund"

# 调用复制函数
copy_directory_contents(source_directory)
