"""
tools
"""
import numpy as np
import pandas as pd


def bo_fen(prev_dif, cur_dif, next_dif):
    if cur_dif > prev_dif and cur_dif > next_dif:
        return True
    return False


def bo_gu(prev_dif, cur_dif, next_dif):
    if cur_dif < prev_dif and cur_dif < next_dif:
        return True
    return False


def jin_cha(prev_dif, prev_dea, cur_dif, cur_dea):
    if prev_dif < prev_dea and cur_dif > cur_dea:
        return True
    return False


def modify(df: pd.DataFrame) -> pd.DataFrame:
    # df.rename(
    #     inplace=True,
    #     columns={
    #         "日期": "date",
    #         "开盘": "open",
    #         "收盘": "close",
    #         "最高": "high",
    #         "最低": "low",
    #     },
    # )

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
