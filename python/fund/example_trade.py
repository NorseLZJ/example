from datetime import datetime

import backtrader as bt
import matplotlib.pyplot as plt
import akshare as ak
import pandas as pd
import numpy as np

plt.rcParams["font.sans-serif"] = ["SimHei"]
plt.rcParams["axes.unicode_minus"] = False

stock_qfq_df = ak.stock_zh_a_hist(symbol="002424", adjust="qfq").iloc[:, :6]
stock_qfq_df.columns = [
    "date",
    "open",
    "close",
    "high",
    "low",
    "volume",
]
stock_qfq_df.index = pd.to_datetime(stock_qfq_df["date"])


class MyStrategy(bt.Strategy):
    """
    主策略程序
    """

    def __init__(self):
        """
        初始化函数
        """
        self.data_close = self.datas[0].close
        self.order = None
        self.buy_price = None
        self.buy_comm = None

        self.ma20 = bt.indicators.SimpleMovingAverage(self.datas[0], period=20)

    def next(self):
        """
        执行逻辑
        """
        if self.order:
            return

        ma20_today = self.ma20[0]
        ma20_yesterday = self.ma20[-1]

        if not self.position:  # 没有持仓
            if (
                ma20_yesterday != np.NAN
                and ma20_today > ma20_yesterday
                and self.data_close[0] > self.ma20[0]
            ):
                self.order = self.buy(size=100)  # 执行买入
            elif self.data_close[0] > self.ma20[0]:
                self.order = self.buy(size=100)  # 执行买入
        else:
            if self.data_close[0] < self.ma20[0]:
                self.order = self.sell(size=100)  # 执行卖出


cerebro = bt.Cerebro()  # 初始化回测系统
start_date = datetime(2023, 1, 1)  # 回测开始时间
end_date = datetime(2024, 3, 5)  # 回测结束时间
data = bt.feeds.PandasData(
    dataname=stock_qfq_df, fromdate=start_date, todate=end_date
)  # 加载数据
cerebro.adddata(data)  # 将数据传入回测系统
cerebro.addstrategy(MyStrategy)  # 将交易策略加载到回测系统中
start_cash = 100000
cerebro.broker.setcash(start_cash)  # 设置初始资本为 100000
cerebro.broker.setcommission(commission=0.002)  # 设置交易手续费为 0.2%
cerebro.run()  # 运行回测系统

port_value = cerebro.broker.getvalue()  # 获取回测结束后的总资金
pnl = port_value - start_cash  # 盈亏统计

print(
    f"初始资金: {start_cash}\n回测期间：{start_date.strftime('%Y%m%d')}:{end_date.strftime('%Y%m%d')}"
)
print(f"总资金: {round(port_value, 2)}")
print(f"净收益: {round(pnl, 2)}")

cerebro.plot(style="candlestick")  # 画图
