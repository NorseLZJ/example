import akshare as ak
import talib
import numpy as np
from datetime import datetime, date
import datetime as dt
import pandas as pd
from redis import *

short_win = 12  # 短期EMA平滑天数
long_win = 26  # 长期EMA平滑天数
macd_win = 20  # DEA线平滑天数
pd.set_option("display.max_columns", None)
pd.set_option("display.width", 500)

r = Redis(host="127.0.0.1", port=6379, decode_responses=True)
industry_key = "industry"


def invide_stock_code(symbol: str):
    if len(str(symbol)) != 6:
        return False
    if str(symbol)[0:2] not in ("00", "60", "30"):
        return False
    return True


def get_industry(_symbol: str, simple_name: str = ""):
    if simple_name.find("退") != -1 or simple_name.find("ST") != -1:
        return np.nan
    val = r.hget(industry_key, _symbol)
    if val is None:
        try:
            vv = ak.stock_individual_info_em(_symbol)
        except Exception as e:
            print("get industry(%s,%s) err:%s" % (_symbol, simple_name, e))
            return np.nan
        item = vv["item"]
        value = vv["value"]
        for i in range(len(item)):
            if item[i] == "行业":
                r.hset(industry_key, _symbol, value[i])
                print("new industry(%s,%s)" % (_symbol, simple_name))
                return value[i]
    else:
        return str(val)


def day_60_plus(x):
    """
    60天线是否上升趋势
    :return:
    """
    max_idx = len(x) - 1
    stop_idx = max_idx - 10
    if stop_idx > 0:
        while max_idx > stop_idx:
            if not np.isnan(x[max_idx]) and not np.isnan(x[max_idx - 1]):
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


def get_daily_data(_symbol: str):
    td = dt.date.today()
    end_date = str(td).replace("-", "")
    timestamp = datetime.timestamp(datetime.now())
    dt_object = datetime.fromtimestamp(int(timestamp) - 356 * 2 * 86400)
    start_date = (str(dt_object).split(" ")[0]).replace(" ", "")
    name = get_name_akshare(_symbol)
    if name == "":
        return None
    try:
        # 拿一年左右前复权的数据
        df = ak.stock_zh_a_daily(name, start_date=start_date, end_date=end_date, adjust="qfq")
        return df
    except Exception as e:
        print("get stock daily data[%s] err:%s" % (_symbol, e))
        return None


def get_minute_data(_symbol: str, period: str):
    if period not in ("1", "5", "15", "30", "60"):
        return None
    name = get_name_akshare(_symbol)
    if name == "":
        return None
    try:
        df = ak.stock_zh_a_minute(name, period=period, adjust="qfq")
        return df
    except Exception as e:
        print("get stock minute(%s) data[%s] err:%s" % (period, _symbol, e))
        return None


def get_params_by_key(_df, key_list, idx):
    if idx > len(_df):
        return None
    ret = []
    for key in key_list:
        ret.append(_df.loc[idx][key])
    return ret


def get_params(_df, idx):
    """
    params
        _df pandas.DataFrame
        idx _df index
    return (date,open,high,low,close,dif,dea,macd,ma5,ma10,ma20,ma60)
    """
    if idx > len(_df):
        return None
    return (
        _df.loc[idx]["date"],
        _df.loc[idx]["open"],
        _df.loc[idx]["high"],
        _df.loc[idx]["low"],
        _df.loc[idx]["close"],
        _df.loc[idx]["dif"],
        _df.loc[idx]["dea"],
        _df.loc[idx]["macd"],
        _df.loc[idx]["ma5"],
        _df.loc[idx]["ma10"],
        _df.loc[idx]["ma20"],
        _df.loc[idx]["ma60"],
    )


def collect_data_by_json(_data):
    _df = pd.read_json(_data)
    return collect_data_by_df(_df)


def collect_data_by_df(_df):
    if _df is None:
        return None
    (dif, dea, macd) = talib.MACD(_df["close"], fastperiod=short_win, slowperiod=long_win, signalperiod=macd_win)
    ma5 = np.around(talib.SMA(_df["close"], timeperiod=5), 2)
    ma10 = np.around(talib.SMA(_df["close"], timeperiod=10), 2)
    ma20 = np.around(talib.SMA(_df["close"], timeperiod=20), 2)
    ma60 = np.around(talib.SMA(_df["close"], timeperiod=60), 2)
    dif = np.around(dif, 2)
    dea = np.around(dea, 2)
    macd = np.around(macd, 2)
    _df["upper"], _df["middle"], _df["lower"] = talib.BBANDS(
        _df.close.values,
        timeperiod=20,
        # number of non-biased standard deviations from the mean
        nbdevup=2,
        nbdevdn=2,
        # Moving average type: simple moving average here
        matype=0,
    )

    _df.insert(loc=5, column="dif", value=dif)
    _df.insert(loc=5, column="dea", value=dea)
    _df.insert(loc=5, column="macd", value=macd)
    _df.insert(loc=5, column="ma5", value=ma5)
    _df.insert(loc=5, column="ma10", value=ma10)
    _df.insert(loc=5, column="ma20", value=ma20)
    _df.insert(loc=5, column="ma60", value=ma60)

    return _df


def clean_data_by_name(df: pd.DataFrame, type: str) -> pd.DataFrame:
    def clean_signal(name: str):
        if name.find("ST") != -1:
            return np.nan

        if name.find("退") != -1:
            return np.nan

        return 1

    df["del"] = df.apply(lambda x: clean_signal(x["股票简称"]), axis=1)

    # del_idx = df[df["del"] == True].index
    # df.drop(del_idx, inplace=True)

    df.dropna(inplace=True, axis=1)
    df.drop(columns=["序号"], inplace=True, axis=1)

    if type == "zcfz":
        pass
    elif type == "lrb":
        di = df[df["净利润同比"] <= 0.0].index
        df.drop(di, inplace=True)
        df["industry"] = df.apply(lambda x: get_industry(x["股票代码"], x["股票简称"]), axis=1)
        df.sort_values(by=["industry", "净利润同比"], inplace=True, ignore_index=True, ascending=False)
    elif type == "xjll":
        pass
    df.reset_index(inplace=True)
    return df
