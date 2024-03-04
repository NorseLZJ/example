import akshare as ak
import pandas as pd
import numpy as np


def stock_data(period="", symbol="", start_date="") -> pd.DataFrame:
    df = ak.stock_zh_a_hist(period=period, symbol=symbol, start_date=start_date)
    columns = {
        "日期": "date",
        "开盘": "open",
        "收盘": "close",
        "最高": "high",
        "最低": "low",
    }
    df.rename(inplace=True, columns=columns)

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

    column = [
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
    df.drop(columns=["ema12", "ema26"], inplace=True)
    df[column] = np.round(df[column].astype(float), 3)
    return df


if __name__ == "__main__":
    # start_date = "20200101"
    start_date = ""
    symbol_list = ["300390", "600029", "688388"]
    # df_day = stock_data(period="daily", symbol=symbol, start_date=start_date)
    for symbol in symbol_list:
        df_week = stock_data(period="weekly", symbol=symbol, start_date=start_date)
        df_week["flag"] = (df_week["ma20"] > df_week["ma20"].shift()) & (
            df_week["close"] > df_week["ma20"]
        )
        df_week["flag"] = df_week["flag"].astype(int)

        # df_day.to_csv("daily.csv")
        # df_week.to_csv(f"week.csv")
        df_week.to_excel(f"{symbol}_week.xlsx")
