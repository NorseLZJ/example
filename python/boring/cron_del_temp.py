import os
import shutil
import logging

__dir = format("C:\\Users\\%s\\AppData\\Local\\Temp\\" % (os.getlogin()))
__local = format("C:\\Users\\%s\\AppData\\Local\\" % (os.getlogin()))


def rm_dir(path):
    try:
        shutil.rmtree(path)
    except Exception as e:
        logging.info(format("dir failed %s, e:%s" % (path, e)))


def rm_file(path):
    try:
        os.remove(path)
    except Exception as e:
        logging.info(format("file failed %s, e:%s" % (path, e)))


def clear_temp():
    """
    清理temp目录
    """
    items = os.listdir(__dir)
    for i in items:
        _path = format("%s%s" % (__dir, i))
        if os.path.isdir(_path):
            rm_dir(_path)
        else:
            rm_file(_path)


def black_items(item):
    """
    文件夹包含以下的都可以删除，无害
    """
    black_list = ["360", "2345","updater"]
    for i in black_list:
        if str(item).find(i) != -1:
            return True
    return False


def clear_local():
    """
    清理程序升级残留目录
    """
    items = os.listdir(__local)
    for i in items:
        if black_items(i) is False:
            continue
        _path = format("%s%s" % (__local, i))
        rm_dir(_path)


if __name__ == "__main__":
    if os.path.exists("D:\\Log\\") is False:
        os.mkdir("D:\\Log\\")

    logging.basicConfig(
        filename="D:\\Log\\Del_Temp.log",
        level=logging.DEBUG,
        format="%(asctime)s %(message)s",
        datefmt="%m/%d/%Y %I:%M:%S %p",
    )
    clear_temp()
    clear_local()
