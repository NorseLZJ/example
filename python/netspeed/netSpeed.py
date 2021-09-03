"""
netSpeed.py
"""
import time
import psutil
import sys
import threading
from PyQt5.QtWidgets import QApplication, QWidget, QLabel
import speed as s


def get_net_speed():
    sent_0 = psutil.net_io_counters().bytes_sent
    recv_0 = psutil.net_io_counters().bytes_recv
    time.sleep(1)
    sent_1 = psutil.net_io_counters().bytes_sent
    recv_1 = psutil.net_io_counters().bytes_recv
    send = (sent_1 - sent_0) / 1024
    recv = (recv_1 - recv_0) / 1024
    now = time.strftime("[%Y-%m-%d %H:%M:%S]", time.localtime())

    up = format("↑ %.2f KB/s" % send)
    dw = format("↓ %.2f KB/s" % recv)
    return up, dw, now


def auto_set_up_and_dw(ex):
    while True:
        up, dw, now = get_net_speed()
        pass
        ex.set_up_and_dw(up, dw)


if __name__ == '__main__':
    app = QApplication(sys.argv)
    w = QWidget()
    w.resize(100, 100)
    w.move(300, 300)
    w.setWindowTitle("netSpeed")
    ex = s.Ui_Form(w)
    w.show()

    t1 = threading.Thread(target=auto_set_up_and_dw, args=(ex,))
    t1.start()

    app.exec_()
    t1.join()
    sys.exit()
