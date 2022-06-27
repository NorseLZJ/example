from comm import *
import numpy as np
from redis import *
from datetime import datetime as dt

r = Redis(host="192.168.66.47", port=6381, decode_responses=True)


def check_avg(_symbol: str) -> bool:
    df = get_daily_data(_symbol)
    if df is None:
        return False
    cp_df = collect_data_by_df(df)
    cp_df.dropna(inplace=True, axis=0)
    cp_df.reset_index(inplace=True)

    (_, _, _, _, close, _, _, _, _, _, ma20, ma60, _) = get_params(
        cp_df, len(cp_df) - 1
    )

    # TODO 当前股价在60,20天均线以上 但是超过60均线不足5%
    if ma60 != 0 and ma20 != 0:
        if (close > ma60) and (close > ma20):
            ret = (close - ma60) / ma60 * 100
            if ret < 5.0:
                return True
    return False


class Stock(object):
    def __init__(self, _symbol: str, _name: str, _price):
        self.symbol = _symbol
        self.name = _name
        self.price = _price

    def info(self):
        return format("\t%s %s %f\n" % (self.symbol, self.name, self.price))


if __name__ == "__main__":
    key = format("xstp_%s" % (str(dt.today()).split(" ")[0]))
    data = r.get(key)
    v = None
    if data is None:
        v = ak.stock_rank_xstp_ths(symbol="20日均线")
        r.set(key, v.to_json())
    else:
        v = pd.read_json(data)
    simple = v["股票简称"]
    last_prices = v["最新价"]
    txt = ""
    industry_list = []
    symbol_list = []
    code_list = []
    price_list = []
    for idx in range(len(simple)):
        last_price = last_prices[idx]
        symbol = str(simple[idx])
        if symbol.find("ST") == -1 and symbol.find("退市") == -1 and last_price >= 5:
            # print(simple[idx])
            code = str(v["股票代码"][idx])
            prefix = code[0:3]
            if prefix in ("000", "300", "600") and check_avg(code) is True:
                industry_list.append(get_industry(code))
                symbol_list.append(symbol)
                code_list.append(code)
                price_list.append(last_price)

    df = pd.DataFrame(
        data={
            "industry": industry_list,
            "symbol": symbol_list,
            "code": code_list,
            "price": price_list,
        }
    )
    df.to_excel("today.xlsx")
