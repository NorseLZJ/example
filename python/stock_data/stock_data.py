"""
get date from network
"""
import akshare as ak
import pandas as pd
import talib
import numpy
from select_move import *
import comm as cm

''' stock state'''
purchase_ing = 1
sellout_ing = 2

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
        if self.state == purchase_ing:
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
        self.avg60 = {}
        self.avg_calc()

        # 日线数据
        self.minute60()

    def get_avg(self, _date):
        if self.avg60 is None:
            return None
        key = str(_date).split(' ')[0]
        xx = self.avg60.get(key)
        if xx == 0.0:
            return None
        return xx

    def avg_calc(self):
        df = get_daily_date(self.code)
        if df is None:
            return False
        c_date = df['date']
        c_close = df['close']
        prices = []
        for i in range(len(c_date)):
            prices.append(float(c_close[i]))

        xx = numpy.array(prices)
        if numpy.isnan(xx[len(xx) - 1]):  # 最后一个数据是nan，那前边的就不处理了
            return False

        # prices = numpy.around(xx, 2)
        _, _, _, self.avg60 = get_avg_list(xx, c_date)

    def minute60(self):
        df = ak.stock_zh_a_minute(get_name_akshare(self.code), period='60')
        days = df['day']
        closes = df['close']
        f_closes = []
        for i in closes:
            f_closes.append(float(i))

        xx = numpy.array(f_closes)
        vv = talib.MACD(xx, fastperiod=SHORT_WIN, slowperiod=LONG_WIN, signalperiod=MACD_WIN)
        v1 = numpy.around(vv[0], 3)
        i = 0
        while True:
            if numpy.isnan(v1[i]):
                i += 1
            else:
                break

        # vv = pd.DataFrame(v1[i:], days[i:], f_closes[i:])
        # self.calc(vv, vv, vv)
        d_t = type(days[i:])
        self.calc(v1[i:], days[i:], f_closes[i:])

    def calc(self, dif, dates, prices_close):
        for i in range(len(dif)):
            if i <= 0 or i > (len(dif) - 2):
                continue  # start or end

            next_date = dates[i + 1]
            cur_date = dates[i]
            prev_dif = dif[i - 1]
            cur_dif = dif[i]
            next_close = prices_close[i + 1]
            cur_close = float(prices_close[i])
            avg60 = self.get_avg(cur_date)

            if avg60 is None:
                continue
            # TODO 条件，判定买卖关键位置
            if cur_dif > prev_dif and cur_close >= avg60:
                '''当前大于前一个，谷底，操作后一个到后一个+1'''
                self.sixtyList.append(SixtyMinter(cur_date, purchase_ing, cur_dif, next_close, next_date))
            elif cur_dif < prev_dif:
                '''当前小于前一个,谷顶'''
                self.sixtyList.append(SixtyMinter(cur_date, sellout_ing, cur_dif, next_close, next_date))

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
            if prev_state == sellout_ing and state == purchase_ing:  # 前一个状态是卖出，现在需要买入
                price1 = price
                date1 = c_date
            elif prev_state == purchase_ing and state == sellout_ing:
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
                add_money = (val * count) - 10  # 10手续费等
                bill_txt += v.bill_info(0, add_money)
                all_money += add_money
            else:  # 亏钱
                sub_money = (val * count) - 10  # 10手续费等
                bill_txt += v.bill_info(-1, sub_money)
                all_money += sub_money
        self.bill_cent = bill_txt
        self.bill_cent += format("总收益:%d" % all_money)

    def get_content(self):
        return self.content, self.bill_cent
