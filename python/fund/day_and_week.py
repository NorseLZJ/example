import akshare as ak
import pandas as pd
import numpy as np


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
    return modify(df)


def etf_data(period="", symbol="", start_date="") -> pd.DataFrame:
    df = ak.fund_etf_hist_em(
        period=period, symbol=symbol, start_date=start_date, adjust="qfq"
    )
    return modify(df)


if __name__ == "__main__":
    start_date = ""
    symbol_list = ["300390", "600029", "688388"]
    etf_list = ["515120", "561160"]

    for symbol in symbol_list:
        df_week = stock_data(period="weekly", symbol=symbol, start_date=start_date)
        df_week["flag"] = (df_week["ma20"] > df_week["ma20"].shift()) & (
            df_week["close"] > df_week["ma20"]
        )
        df_week["flag"] = df_week["flag"].astype(int)
        df_week.to_excel(f"{symbol}_week.xlsx")

    for symbol in etf_list:
        df_week = etf_data(period="weekly", symbol=symbol, start_date=start_date)
        df_week["flag"] = (df_week["ma20"] > df_week["ma20"].shift()) & (
            df_week["close"] > df_week["ma20"]
        )
        df_week["flag"] = df_week["flag"].astype(int)
        df_week.to_excel(f"etf_{symbol}_week.xlsx")
