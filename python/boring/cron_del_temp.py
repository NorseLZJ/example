import os
import shutil
import logging

__dir = format("C:\\Users\\%s\\AppData\\Local\\Temp\\" % (os.getlogin()))

if __name__ == "__main__":
    if os.path.exists("D:\\Log\\") is False:
        os.mkdir("D:\\Log\\")

    logging.basicConfig(filename='D:\\Log\\Del_Temp.log',
                        level=logging.DEBUG,
                        format='%(asctime)s %(message)s',
                        datefmt='%m/%d/%Y %I:%M:%S %p')
    items = os.listdir(__dir)
    for i in items:
        _path = format("%s%s" % (__dir, i))
        if os.path.isdir(_path):
            try:
                shutil.rmtree(_path)
            except Exception as e:
                logging.info(format("dir failed %s, e:%s" % (_path, e)))
        else:
            try:
                os.remove(_path)
            except Exception as e:
                logging.info(format("file failed %s, e:%s" % (_path, e)))
