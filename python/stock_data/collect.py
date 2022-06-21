import matplotlib.pyplot as plt
import numpy as np
import scipy.stats as stats
import scipy.optimize as opt
from pandas import *
import talib
import pandas as pd
import numpy

from redis import *

r = Redis(host='192.168.66.47', port=6381, decode_responses=True)
short_win = 12  # 短期EMA平滑天数
long_win = 26  # 长期EMA平滑天数
macd_win = 20  # DEA线平滑天数
pd.set_option('display.max_columns', None)
pd.set_option('display.width', 500)

if __name__ == "__main__":
    keys = r.keys('*')
    data = r.get(keys[0])
    df = pd.read_json(data)
    (dif, dea, macd) = talib.MACD(df['close'], fastperiod=short_win, slowperiod=long_win, signalperiod=macd_win)
    ma5 = numpy.around(talib.SMA(df['close'], timeperiod=5), 2)
    ma10 = numpy.around(talib.SMA(df['close'], timeperiod=10), 2)
    ma20 = numpy.around(talib.SMA(df['close'], timeperiod=20), 2)
    ma60 = numpy.around(talib.SMA(df['close'], timeperiod=60), 2)
    dif = numpy.around(dif, 2)
    dea = numpy.around(dea, 2)
    macd = numpy.around(macd, 2)

    df.insert(loc=5, column='dif', value=dif)
    df.insert(loc=5, column='dea', value=dea)
    df.insert(loc=5, column='macd', value=macd)
    df.insert(loc=5, column='ma5', value=ma5)
    df.insert(loc=5, column='ma10', value=ma10)
    df.insert(loc=5, column='ma20', value=ma20)
    df.insert(loc=5, column='ma60', value=ma60)
    df = df.dropna()
    print(df.tail(10))

    '''
    dat = df.set_index('date')[['close', 'open', 'high', 'low']]
    dat.plot(title="Close Price of SINOPEC (600028) during Jan, 2015")
    plt.show()
    '''
