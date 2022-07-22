from redis import *
import akshare as ak
from datetime import datetime, date
import datetime as dt
from comm import *
import os

r = Redis(host="127.0.0.1", port=6379, decode_responses=True)
buff_file = "stock_data/buff.csv"


def cache_data(symbol: str, name: str):
    out_file = get_stock_data_file(symbol)
    if invide_stock_code(symbol) is False:
        return
    if name.find("ST") != -1 or name.find("退") != -1:
        if os.path.exists(out_file):
            os.remove(out_file)
        return

    td = dt.date.today()
    end_date = str(td).replace("-", "")
    now_stamp = datetime.timestamp(datetime.now())
    dt_object = datetime.fromtimestamp(int(now_stamp) - 356 * 2 * 86400)
    start_date = (str(dt_object).split(" ")[0]).replace(" ", "")
    name = get_name_akshare(symbol)
    if name == "":
        return None

    # 不存在就获取一个
    if not os.path.exists(out_file):
        try:
            df = ak.stock_zh_a_daily(name, start_date=start_date, end_date=end_date, adjust="qfq")
            df.to_csv(out_file, index=False)
        except Exception as e:
            print("get stock daily data[%s] err:%s" % (symbol, e))
        return
    else:
        df1 = pd.read_csv(out_file)
        last_date = df1.iloc[-1]["date"]
        now = datetime.now()
        month, day = now.month, now.day
        strmonth, strday = "", ""
        if month < 10:
            strmonth = format("0%d" % month)
        else:
            strmonth = format("%d" % month)
        if day < 10:
            strday = format("0%d" % day)
        else:
            strday = format("%d" % day)

        now_date = format("%d-%s-%s" % (now.year, strmonth, strday))
        if now_date != last_date and now.hour >= 15:
            try:
                df = ak.stock_zh_a_daily(name, start_date=last_date, end_date=now_date, adjust="qfq")
                df.to_csv(buff_file)
                dfnew = pd.concat([df1, pd.read_csv(buff_file)], ignore_index=True)
                dfnew.drop_duplicates(inplace=True)
                dfnew.to_csv(out_file, index=False)
            except Exception as e:
                print("get stock daily data[%s] err:%s" % (symbol, e))


if __name__ == "__main__":
    df = ak.stock_lrb_em()
    df.apply(lambda x: cache_data(x["股票代码"], x["股票简称"]), axis=1)
    if os.path.exists(buff_file):
        os.remove(buff_file)
