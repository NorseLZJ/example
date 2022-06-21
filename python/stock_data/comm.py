import akshare as ak
import talib
import numpy
from datetime import datetime, date
import datetime as dt


def day_60_plus(x):
    """
    60天线是否上升趋势
    :return:
    """
    max_idx = len(x) - 1
    stop_idx = max_idx - 10
    if stop_idx > 0:
        while max_idx > stop_idx:
            if not numpy.isnan(x[max_idx]) and not numpy.isnan(x[max_idx - 1]):
                if x[max_idx] >= x[max_idx - 1]:
                    max_idx -= 1
                else:
                    return False
            else:
                return False
    return True


def get_name_akshare(code: str):
    if code[0:2] == "00" or code[0:2] == "30":  # sz
        return format("sz%s" % code)
    return format("sh%s" % code)


def get_daily_date(_symbol: str):
    td = dt.date.today()
    end_date = str(td).replace('-', '')
    timestamp = datetime.timestamp(datetime.now())
    dt_object = datetime.fromtimestamp(int(timestamp) - 356 * 86400)
    start_date = (str(dt_object).split(' ')[0]).replace(' ', '')
    name = get_name_akshare(_symbol)
    if name == '':
        return None
    try:
        # 拿一年左右前复权的数据
        df = ak.stock_zh_a_daily(name, start_date=start_date, end_date=end_date, adjust='qfq')
        return df
    except Exception as e:
        print("get stock[%s] err:%s" % (_symbol, e))
        return None


def get_avg_list(_xx, _date):
    avg5 = talib.SMA(_xx, timeperiod=5)
    avg10 = talib.SMA(_xx, timeperiod=10)
    avg20 = talib.SMA(_xx, timeperiod=20)
    avg60 = talib.SMA(_xx, timeperiod=60)

    avg5 = numpy.around(avg5, 2)
    avg10 = numpy.around(avg10, 2)
    avg20 = numpy.around(avg20, 2)
    avg60 = numpy.around(avg60, 2)

    _avg5 = {}
    _avg10 = {}
    _avg20 = {}
    _avg60 = {}
    for i in range(len(_date)):
        a5 = avg5[i] if not numpy.isnan(avg5[i]) else 0.0
        a10 = avg10[i] if not numpy.isnan(avg10[i]) else 0.0
        a20 = avg20[i] if not numpy.isnan(avg20[i]) else 0.0
        a60 = avg60[i] if not numpy.isnan(avg60[i]) else 0.0
        key = str(_date[i]).split(' ')[0]
        _avg5[key] = a5
        _avg10[key] = a10
        _avg20[key] = a20
        _avg60[key] = a60
    return _avg5, _avg10, _avg20, _avg60
