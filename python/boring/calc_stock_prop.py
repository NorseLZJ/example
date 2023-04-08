import akshare as ak
import numpy as np


def check_position(code):
    # 获取股票的日线数据
    try:
        df = ak.stock_zh_a_daily(symbol=ak.stock_a_code_to_symbol(code), adjust="qfq")
        # 计算60日和20日均线
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
        df.drop(inplace=True, columns=["volume", "outstanding_share", "turnover", "ema12", "ema26"])
        column = ["ma60", "ma20", "ma10", "ma5", "dif", "dea", "macd"]
        df[column] = np.round(df[column].astype(float), 3)

        print(df.tail(10))
        df.to_csv("temp.csv")
        # 获取最新一天的数据
        latest = df.iloc[-1]
        # 获取当前价格，60日均线和20日均线的值
        price = latest["close"]
        ma60 = latest["ma60"]
        ma20 = latest["ma20"]
        if ma60 < ma20:  # 如果60日均线在下方
            if price < ma20 and price > ma60:  # 如果当前价格在两条均线之间
                return True  # 返回True
            else:  # 否则
                return False  # 返回False
        else:  # 如果两条均线重合
            return False  # 返回False
    except Exception as e:
        print(f"出现错误 code:{code} err:{e}")
        return False


if __name__ == "__main__":
    codes = ["601766"]
    for code in codes:
        result = check_position(code)
        print(f"股票{code}当前价格是否在60日和20日均线的中间位置：{result}")
