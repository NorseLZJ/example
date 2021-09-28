# -*- coding: utf-8 -*-

import os
import sys
import json
from PyQt5.QtCore import *
from PyQt5.QtWidgets import *
from PyQt5.QtWebEngineWidgets import *

url = 'http://sdk424.1iy.cc/sdk/UserCenter/login?f=eyJwcm9tb3RlX2lkIjoiNDgzIiwiZ2FtZV9pZCI6IjQyNCJ9#/'


def read_conf():
    path = 'client.json'
    if not os.path.exists(path):
        return
    with open(path, 'r') as f:
        data = f.read()
        temp = json.loads(data)
        global url
        url = temp['url']


class MainWindow(QMainWindow):
    def __init__(self):
        super(MainWindow, self).__init__()
        self.setWindowTitle('传奇H5')
        self.setGeometry(5, 30, 1355, 730)
        self.browser = QWebEngineView()
        self.browser.load(QUrl(url))
        self.setCentralWidget(self.browser)


if __name__ == '__main__':
    read_conf()
    app = QApplication(sys.argv)
    win = MainWindow()
    win.show()
    app.exit(app.exec_())
