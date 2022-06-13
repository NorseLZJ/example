"""
5,10,20 天均线选股

5上穿，10，20，均线
"""
import time

import akshare as ak
import talib
import numpy
from datetime import datetime, date
import datetime as dt
from comm import *


def get_day_data(code: str) -> bool:
    """
    计划用5，10，20天线做买入依据
    :return:
    """
    td = dt.date.today()
    end_date = str(td).replace('-', '')
    timestamp = datetime.timestamp(datetime.now())
    dt_object = datetime.fromtimestamp(int(timestamp) - 100 * 86400)  # 拿40天前到现在
    start_date = (str(dt_object).split(' ')[0]).replace(' ', '')

    name = get_name(code)
    if name == '':
        return False

    df = ak.stock_zh_a_daily(name, start_date=start_date, end_date=end_date)
    c_date = df['date']
    c_close = df['close']
    prices = []
    for idx in range(len(c_date)):
        # print(c_date[idx], ' ', c_close[idx])
        prices.append(float(c_close[idx]))

    x = numpy.array(prices)
    all_nan = True
    for v in x:
        if not numpy.isnan(v):
            all_nan = False
            break
    if all_nan is True:
        return False

    prices = numpy.around(x, 2)
    avg5 = talib.SMA(x, timeperiod=5)
    avg10 = talib.SMA(x, timeperiod=10)
    avg20 = talib.SMA(x, timeperiod=20)
    avg60 = talib.SMA(x, timeperiod=60)

    avg5 = numpy.around(avg5, 2)
    avg10 = numpy.around(avg10, 2)
    avg20 = numpy.around(avg20, 2)
    avg60 = numpy.around(avg60, 2)
    idx = len(avg5) - 1
    a5 = avg5[idx] if not numpy.isnan(avg5[idx]) else 0
    a10 = avg10[idx] if not numpy.isnan(avg10[idx]) else 0
    a20 = avg20[idx] if not numpy.isnan(avg20[idx]) else 0
    a60 = avg60[idx] if not numpy.isnan(avg60[idx]) else 0
    cur_price = prices[len(prices) - 1]
    if a60 != 0 and a20 != 0 and a10 != 0 and a5 != 0:
        if (cur_price > a60) and (cur_price > a20):
            return True
    return False


if __name__ == "__main__":
    # print(get_day_data('300595'))
    xx = ak.stock_info_a_code_name()
    codes = xx['code']
    with open('see.txt', 'w') as f:
        for i in codes:
            if get_day_data(str(i)) is True:
                f.write(format("%s\n" % str(i)))
