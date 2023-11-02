import akshare as ak
import pandas as pd

last_df = ak.stock_report_fund_hold_detail("510050", "20230631")
print(last_df.head(100))
last_df.to_excel("aaa.xlsx", index=False)


"""
# 获取上期上证50指数的成分股
last_period_stocks = ak.stock_zh_index_cons(index="000016.SH")

# 获取当前上证50指数的成分股
current_period_stocks = ak.stock_zh_index_cons(index="000016.SH", end_date="")

# 合并数据
merged_data = pd.merge(current_period_stocks, last_period_stocks, on="成分券代码", suffixes=("_当前", "_上期"))

# 计算调出的股票
out_stocks = merged_data[merged_data["是否调出"] == "是"]

# 计算调入的股票
in_stocks = merged_data[merged_data["是否调入"] == "是"]

# 输出结果为表格
print("调出股票:")
print(out_stocks)
print("\n调入股票:")
print(in_stocks)
"""
