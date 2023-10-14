# 导入akshare模块
import akshare as ak

# 导入numpy模块
import numpy as np
# 定义股票代码和日期范围
stock_code = "sh000001"  # 浦发银行
start_date = "2022-01-01"  # 开始日期
end_date = "2022-12-31"  # 结束日期

# 获取股票的日线行情数据
stock_df = ak.stock_zh_a_daily(symbol=stock_code, start_date=start_date, end_date=end_date)

# 随机生成5个买入点和5个卖出点
buy_points = np.random.choice(stock_df.index, size=5, replace=False)
sell_points = np.random.choice(stock_df.index, size=5, replace=False)


# 假设本金10万
capital = 100000

# 定义手续费参数
commission_rate = 0.0003  # 交易佣金率，最高不超过0.003，最低按照5元收取
commission_min = 5  # 交易佣金最低值
stamp_tax_rate = 0.001  # 印花税率，只在卖出时收取
transfer_fee_rate = 0.00002  # 过户费率，只在沪市收取，最低按照2元收取
transfer_fee_min = 2  # 过户费最低值

trade_returns = []
cum_returns = []
fees = []
capitals = []
for i in range(5):
    buy_price = stock_df.loc[buy_points[i], "close"]
    sell_price = stock_df.loc[sell_points[i], "close"]
    trade_return = (sell_price - buy_price) / buy_price
    trade_returns.append(trade_return)
    cum_return = np.prod(1 + np.array(trade_returns)) - 1
    cum_returns.append(cum_return)
    # 计算买入手续费，包括交易佣金和过户费（如果是沪市）
    buy_commission = max(capital * commission_rate, commission_min)
    if stock_code.startswith("sh"):
        buy_transfer_fee = max(capital * transfer_fee_rate, transfer_fee_min)
    else:
        buy_transfer_fee = 0
    buy_fee = buy_commission + buy_transfer_fee
    # 计算卖出手续费，包括交易佣金、过户费（如果是沪市）和印花税
    sell_commission = max(capital * commission_rate, commission_min)
    if stock_code.startswith("sh"):
        sell_transfer_fee = max(capital * transfer_fee_rate, transfer_fee_min)
    else:
        sell_transfer_fee = 0
    sell_stamp_tax = capital * stamp_tax_rate
    sell_fee = sell_commission + sell_transfer_fee + sell_stamp_tax
    # 计算总手续费和资金变化曲线
    fee = buy_fee + sell_fee
    fees.append(fee)
    capital = capital * (1 + trade_return) - fee
    capitals.append(capital)

max_drawdown = min(cum_returns) - max(cum_returns)
annual_return = (1 + cum_returns[-1]) ** (250 / len(stock_df)) - 1
risk_free_rate = 0.0281
annual_std = np.std(trade_returns) * np.sqrt(250)
sharpe_ratio = (annual_return - risk_free_rate) / annual_std

print(f"{stock_code} 在 {start_date} 到 {end_date} 期间的交易结果如下：")
print(f"买入点：{buy_points}")
print(f"卖出点：{sell_points}")
print(f"每次交易的收益率：{trade_returns}")
print(f"累计收益率：{cum_returns}")
print(f"最大回撤：{max_drawdown}")
print(f"年化收益率：{annual_return}")
print(f"年化无风险利率：{risk_free_rate}")
print(f"年化收益率的标准差：{annual_std}")
print(f"夏普比率：{sharpe_ratio}")

plt.plot(capitals, label="capital")
plt.title(f"{stock_code} capital curve")
plt.xlabel("trade times")
plt.ylabel("capital")
plt.legend()
plt.show()
