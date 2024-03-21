import akshare as ak
import pandas as pd

if __name__ == "__main__":
    df = ak.stock_zh_a_hist()
    df.index = pd.to_datetime(df["日期"])
    # df.set_index("日期", drop=True, inplace=True)
    x = pd.to_datetime("1997-04-04", errors="coerce")
    cur_index = df.index.get_loc(x)
    cur_open, cur_close = (df.iloc[cur_index][["开盘", "收盘"]])
    print(cur_open, " - ", cur_close)
