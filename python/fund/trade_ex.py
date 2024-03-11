from __future__ import absolute_import, division, print_function, unicode_literals

import datetime  # For datetime objects
import akshare as ak
import pandas as pd

import backtrader as bt


class TestStrategy(bt.Strategy):

    def log(self, txt, dt=None):
        """Logging function for this strategy"""
        dt = dt or self.datas[0].datetime.date(0)
        print("%s, %s" % (dt.isoformat(), txt))

    def __init__(self):
        # Keep a reference to the "close" line in the data[0] dataseries
        self.dataclose = self.datas[0].close

        """ 自定义参数 """
        self._size = 0

    def next(self):

        dt = self.datas[0].datetime.date(0)
        dt = f"{dt.isoformat()}"

        today_close = self.dataclose[0]

        if dt == "2022-11-25":
            cash = self.broker.getcash()
            size = int((cash / (today_close * 100))) * 100
            self.order = self.buy(size=size)  # 执行买入
            self._size = size
            self.log(
                f"购入 , cash:{self.broker.getcash()} ,size:{size} ,clsoe:{today_close}"
            )
        if dt == "2023-03-17":
            self.order = self.sell(size=self._size)  # 执行卖出
            self.log(f"卖出 , close:{today_close}")


if __name__ == "__main__":
    cerebro = bt.Cerebro()
    cerebro.addstrategy(TestStrategy)

    stock_qfq_df = ak.stock_zh_a_hist(symbol="600029", adjust="qfq").iloc[:, :6]
    stock_qfq_df.columns = [
        "date",
        "open",
        "close",
        "high",
        "low",
        "volume",
    ]
    stock_qfq_df.index = pd.to_datetime(stock_qfq_df["date"])

    # Create a Data Feed
    data = bt.feeds.PandasData(
        dataname=stock_qfq_df,
        fromdate=datetime.datetime(2022, 1, 1),
        todate=datetime.datetime(2024, 3, 8),
    )
    cerebro.adddata(data)
    cerebro.broker.setcash(10000.0)
    print("Starting Portfolio Value: %.2f" % cerebro.broker.getvalue())
    cerebro.run()
    print("Final Portfolio Value: %.2f" % cerebro.broker.getvalue())
