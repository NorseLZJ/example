# -*- coding: utf-8 -*-

# Form implementation generated from reading ui file 'speed.ui'
#
# Created by: PyQt5 UI code generator 5.9.2
#
# WARNING! All changes made in this file will be lost!

from PyQt5 import QtCore, QtGui, QtWidgets


class Ui_Form(object):
    def __init__(self, Form):
        self.setupUi(Form)

    def setupUi(self, Form):
        Form.setObjectName("Form")
        Form.setEnabled(False)
        Form.resize(112, 65)
        self.widget = QtWidgets.QWidget(Form)
        self.widget.setGeometry(QtCore.QRect(0, 0, 92, 50))
        self.widget.setObjectName("widget")
        self.box = QtWidgets.QVBoxLayout(self.widget)
        self.box.setContentsMargins(0, 0, 0, 0)
        self.box.setObjectName("box")
        self.up_speed = QtWidgets.QLabel(self.widget)
        self.up_speed.setStyleSheet("rgb:(85, 170, 0);\n"
                                    "font: 25 12pt \"Microsoft YaHei\";")
        self.up_speed.setObjectName("up_speed")
        self.box.addWidget(self.up_speed)
        self.dw_speed = QtWidgets.QLabel(self.widget)
        self.dw_speed.setStyleSheet("rgb:(85, 170, 0);\n"
                                    "font: 25 12pt \"Microsoft YaHei\";")
        self.dw_speed.setObjectName("dw_speed")
        self.box.addWidget(self.dw_speed)

        self.retranslateUi(Form)
        QtCore.QMetaObject.connectSlotsByName(Form)

    def retranslateUi(self, Form):
        _translate = QtCore.QCoreApplication.translate
        Form.setWindowTitle(_translate("Form", "Form"))
        self.up_speed.setText(_translate("Form", "↑ 0 KB/s      "))
        self.dw_speed.setText(_translate("Form", "↓ 0 KB/s      "))

    def set_up_and_dw(self, up, dw):
        self.up_speed.setText(str(up))
        self.dw_speed.setText(str(dw))
