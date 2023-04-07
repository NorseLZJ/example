import akshare as ak
import time
import pandas as pd
import os


def get_daily(code: str) -> pd.DataFrame:
    now = time.time() - 24 * 60 * 60
    before = now - 5 * 24 * 60 * 60
    end_time = time.strftime("%Y-%m-%d", time.localtime(now))
    start_time = time.strftime("%Y-%m-%d", time.localtime(before))
    try:
        daily_df = ak.stock_zh_a_daily(
            symbol=ak.stock_a_code_to_symbol(code), adjust="qfq", start_date=start_time, end_date=end_time
        )
        return daily_df
    except Exception as e:
        print(f"get daily {code} err:{e}")
        return None


def get_all_stocks():
    """
    获取目前市场上所有的股票，并去除ST的和价格小于等于2块的
    返回值：
    stock_df: DataFrame, 股票代码、名称和行业的DataFrame
    """
    path = "stock.csv"
    try:
        # 获取A股列表
        if os.path.exists(path) is False:
            stock_df = ak.stock_info_a_code_name()
            stock_df.to_csv(path, index=False)

        stock_df = pd.read_csv(path, dtype=str)
        stock_df = stock_df[~stock_df["name"].str.contains("ST")]
        cni_list = stock_df[["code", "name"]].values.tolist()
        stock_list = []
        for code, name in cni_list:
            daily_df = get_daily(code)
            if daily_df is None:
                continue
            price = daily_df.iloc[-1]["close"]
            if price > 2:
                stock_list.append([code, name])
                print(f"good stock {code} {name}")
        stock_df = pd.DataFrame(stock_list, columns=["code", "name"])
        return stock_df
    except Exception as e:
        print(f"出现错误：{e}")
        return None


if __name__ == "__main__":
    # 测试一下函数
    stock_df = get_all_stocks()
    if stock_df is not None:
        print(f"目前市场上所有的股票（去除ST和价格小于等于2块）有{len(stock_df)}只，如下：")
        # print(stock_df)
        stock_df.to_csv("stock.csv", index=False)
