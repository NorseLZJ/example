import akshare as ak
import pandas as pd
import numpy as np

from datetime import datetime, timedelta
import datetime as dt2

std_cache = "cache/"


def modify(df: pd.DataFrame) -> pd.DataFrame:
    df.rename(
        inplace=True,
        columns={
            "日期": "date",
            "开盘": "open",
            "收盘": "close",
            "最高": "high",
            "最低": "low",
        },
    )

    df["ma60"] = df["close"].rolling(60).mean()
    df["ma20"] = df["close"].rolling(20).mean()
    df["ma10"] = df["close"].rolling(10).mean()
    df["ma5"] = df["close"].rolling(5).mean()
    # 计算EMA12和EMA26
    df["ema12"] = df["close"].ewm(span=12).mean()
    df["ema26"] = df["close"].ewm(span=26).mean()
    df["dif"] = df["ema12"] - df["ema26"]
    df["dea"] = df["dif"].ewm(span=9).mean()
    # 计算MACD
    df["macd"] = 2 * (df["dif"] - df["dea"])

    df.drop(columns=["ema12", "ema26"], inplace=True)

    round_column = [
        "ma60",
        "ma20",
        "ma10",
        "ma5",
        "dif",
        "dea",
        "macd",
        "close",
        "open",
        "high",
        "low",
    ]

    df[round_column] = np.round(df[round_column].astype(float), 3)
    return df


def stock_data(period="", symbol="", start_date="") -> pd.DataFrame:
    df = ak.stock_zh_a_hist(
        period=period, symbol=symbol, start_date=start_date, adjust="qfq"
    )
    df = modify(df)
    df["flag"] = (df["ma20"] > df["ma20"].shift()) & (df["close"] > df["ma20"])
    df["flag"] = df["flag"].astype(int)
    df.to_excel(f"{std_cache}{symbol}_week.xlsx")
    return df


def etf_data(period="", symbol="", start_date="") -> pd.DataFrame:
    df = ak.fund_etf_hist_em(
        period=period, symbol=symbol, start_date=start_date, adjust="qfq"
    )
    modify(df)
    df["flag"] = (df["ma20"] > df["ma20"].shift()) & (df["close"] > df["ma20"])
    df["flag"] = df["flag"].astype(int)
    df.to_excel(f"{std_cache}etf_{symbol}_week.xlsx")
    return df


def trade_time_list(df: pd.DataFrame) -> list:
    ok_date = df_week.loc[df_week["flag"] == 1, "date"]
    ok_date = ok_date.to_list()
    time_list = []
    for v in ok_date:
        date_object = datetime.strptime(str(v), "%Y-%m-%d")

        t1 = date_object + timedelta(days=3)
        time_start = datetime.combine(
            t1.date(), datetime.strptime("09:30", "%H:%M").time()
        )

        t2 = date_object + timedelta(days=7)
        time_end = datetime.combine(
            t2.date(), datetime.strptime("15:00", "%H:%M").time()
        )
        # print(f"T1 : {time_start} T2: {time_end}")
        time_list.append([time_start, time_end])
    return time_list


def is_in_time_list(time_list: list, date_str: str) -> bool:
    date_object = dt2.datetime.strptime(date_str, "%Y-%m-%d")
    time_object = dt2.datetime(
        date_object.year, date_object.month, date_object.day, 10, 0
    )
    is_within_interval = any(start <= time_object <= end for start, end in time_list)
    return is_within_interval


if __name__ == "__main__":
    start_date = "20220101"
    # symbol_list = ["300390", "600029", "688388"]
    symbol_list = ["600029"]
    etf_list = ["515120", "561160"]

    for symbol in symbol_list:
        df_week = stock_data(period="weekly", symbol=symbol, start_date=start_date)
        time_list = trade_time_list(df_week)
        # given_time = dt2.datetime(2023, 5, 3, 10, 0)
        # is_within_interval = any(start <= given_time <= end for start, end in time_list)
        date_str = "2023-05-03"
        print(is_in_time_list(time_list, date_str))
        date_str = "2023-04-30"
        print(is_in_time_list(time_list, date_str))
        # print(is_within_interval)

    # for symbol in etf_list:
    #     df_week = etf_data(period="weekly", symbol=symbol, start_date=start_date)
