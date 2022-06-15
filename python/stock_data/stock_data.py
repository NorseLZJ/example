"""
get date from network
"""
import akshare as ak
import talib
import numpy
from select_move import *

''' stock state'''
state_purchase_ing = 1
state_sellout_ing = 2

ONE_DAY_FOUR_LINE = 4

## 使用talib计算MACD的参数
SHORT_WIN = 12  # 短期EMA平滑天数
LONG_WIN = 26  # 长期EMA平滑天数
MACD_WIN = 20  # DEA线平滑天数


class SubPrice(object):
    def __init__(self, price1, date1, price2, date2):
        self.price1 = price1
        self.date1 = date1
        self.price2 = price2
        self.date2 = date2

    def bill_info(self, is_add, val):
        if is_add == 0:
            return format("%s %s [in:%f out:%f] +%d\n" % (self.date1, self.date2, self.price1, self.price2, val))
        else:
            return format("%s %s [in:%f out:%f] %d\n" % (self.date1, self.date2, self.price1, self.price2, val))


class SixtyMinter(object):
    def __init__(self, date, state, dif, price, next_date):
        self.date = date
        self.state = state
        self.dif = dif
        self.price = price
        self.next_date = next_date

    def info_six_minter(self):
        if self.state == state_purchase_ing:
            return format("%s  买入持有   %f\n" % (self.date, self.dif)), self.state, self.price, self.date, self.next_date
        else:
            return format("%s  卖出空仓   %f\n" % (self.date, self.dif)), self.state, self.price, self.date, self.next_date
        return '-' * 10


class StockData(object):
    def __init__(self, code: str):
        self.code = code
        self.sixtyList = []
        self.content = ''
        self.bill_cent = ''
        self.opt_list = []
        # 日线数据
        self.minute60()

    def minute60(self):
        df = ak.stock_zh_a_minute(get_name_akshare(self.code), period='60')
        c_date = df['day']
        c_close = df['close']
        c_close_d = []
        for i in c_close:
            c_close_d.append(float(i))
        x = numpy.array(c_close_d)
        # print(x)
        vv = talib.MACD(x, fastperiod=SHORT_WIN, slowperiod=LONG_WIN, signalperiod=MACD_WIN)
        v1 = numpy.around(vv[0], 3)
        dif = []
        dates2 = []
        closes = []
        for idx in range(len(v1)):
            if numpy.isnan(v1[idx]):
                continue
            dif.append(v1[idx])
            dates2.append(c_date[idx])
            closes.append(c_close_d[idx])
        self.calc(dif, dates2, closes)

    def calc(self, dif, dates, prices_close):
        for idx in range(len(dif)):
            if idx <= 0 or idx > (len(dif) - 2):
                continue  # start or end

            next_date = dates[idx + 1]
            cur_date = dates[idx]
            prev_dif = dif[idx - 1]
            cur_dif = dif[idx]
            next_dif = dif[idx + 1]
            next_close = prices_close[idx + 1]

            # TODO 条件，判定买卖关键位置
            if cur_dif > prev_dif:
                '''当前大于前一个，谷底，操作后一个到后一个+1'''
                self.sixtyList.append(
                    SixtyMinter(cur_date, state_purchase_ing, cur_dif, next_close, next_date))
            elif cur_dif < prev_dif:
                '''当前小于前一个,谷顶'''
                self.sixtyList.append(
                    SixtyMinter(cur_date, state_sellout_ing, cur_dif, next_close, next_date))

        self.info_list()
        self.info_list_rever()

    def info_list_rever(self):
        txt = ''
        for v in reversed(self.sixtyList):
            content, _, _, _, _ = v.info_six_minter()
            txt += content

        self.content = txt

    def info_list(self):
        prev_state = 0
        price1 = 0
        for v in self.sixtyList:
            _, state, price, c_date, _ = v.info_six_minter()

            if prev_state == 0:
                prev_state = state
                continue

            # 处理持仓过程
            if prev_state == state_sellout_ing and state == state_purchase_ing:  # 前一个状态是卖出，现在需要买入
                price1 = price
                date1 = c_date
            elif prev_state == state_purchase_ing and state == state_sellout_ing:
                if price1 != 0 and date1 != '':
                    self.opt_list.append(SubPrice(price1, date1, price, c_date))
            prev_state = state

        # 买卖操作收益
        bill_txt = ''
        all_money = 0
        for v in self.opt_list:
            val = v.price2 - v.price1
            count = int(10000 / v.price1)
            count = int(count / 100) * 100
            if val > 0:  # 赚钱
                add_money = (val * count) - 20  # 20手续费等
                bill_txt += v.bill_info(0, add_money)
                all_money += add_money
            else:  # 亏钱
                sub_money = (val * count) - 20  # 20手续费等
                bill_txt += v.bill_info(-1, sub_money)
                all_money += sub_money
        self.bill_cent = bill_txt
        self.bill_cent += format("总收益:%d" % all_money)

    def get_content(self):
        return self.content, self.bill_cent
