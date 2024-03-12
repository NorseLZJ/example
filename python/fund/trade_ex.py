from __future__ import absolute_import, division, print_function, unicode_literals

import datetime  # For datetime objects

import akshare as ak
import pandas as pd
import numpy as np

import backtrader as bt

import day_and_week as daw

symbol = "600029"


class TestStrategy(bt.Strategy):

    def log(self, txt, dt=None):
        """Logging function for this strategy"""
        dt = dt or self.datas[0].datetime.date(0)
        print("%s, %s" % (dt.isoformat(), txt))

    def __init__(self):
        # Keep a reference to the "close" line in the data[0] dataseries
        self.dataclose = self.datas[0].close
        self.dataopen = self.datas[0].open

        """ 自定义参数 """
        self._size = 0
        self.time_list = daw.trade_time_list(
            daw.stock_data(period="weekly", symbol=symbol)
        )

    def next(self):

        dt = self.datas[0].datetime.date(0)
        dt = f"{dt.isoformat()}"

        # today_close = self.dataclose[0]
        next_open = np.NAN
        if len(self.dataopen) >= 1:
            next_open = self.dataopen[0]

        # self.log(f"RANGE , close:{today_close}")
        isok = daw.is_in_time_list(self.time_list, str(dt))
        if self._size == 0:
            if isok and next_open is not np.NAN:
                cash = np.round(self.broker.getcash())
                size = int((cash / (next_open * 100))) * 100
                self.order = self.buy(size=size)  # 执行买入
                self._size = size
                self.log(f"买入 , cash:{cash} ,size:{size} ,clsoe:{next_open}")
        else:
            if not isok and next_open is not np.NAN:
                self.order = self.sell(size=self._size)  # 执行卖出
                self.log(f"卖出 , close:{next_open}")
                self._size = 0


if __name__ == "__main__":
    cerebro = bt.Cerebro()
    cerebro.addstrategy(TestStrategy)

    stock_qfq_df = ak.stock_zh_a_hist(symbol=symbol, adjust="qfq").iloc[:, :6]
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
