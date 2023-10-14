import akshare as ak
import os
import pandas as pd

"""
获取ETF基金
"""

path = "ETF.csv"

if __name__ == "__main__":
    if os.path.exists(path) is False:
        etf = ak.fund_etf_category_sina(symbol="ETF基金")
        etf.to_csv("ETF.csv", index=False)

    etf = pd.read_csv(path)
    for idx in range (len(etf)):
        etf.iloc[idx][""]
        pass
    print(etf.head(10))
    print(etf.tail(10))
