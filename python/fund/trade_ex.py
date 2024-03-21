from __future__ import absolute_import, division, print_function, unicode_literals

import datetime  # For datetime objects

import akshare as ak
import backtrader as bt
import numpy as np
import pandas as pd

import day_and_week as daw
import tools as tool

# TODO 指定标的
# symbol = "600029"
symbol = "300390"
stock_qfq_df = None


class TestStrategy(bt.Strategy):

    def log(self, txt, dt=None):
        """Logging function for this strategy"""
        dt = dt or self.datas[0].datetime.date(0)
        print("%s, %s" % (dt.isoformat(), txt))

    def __init__(self):
        # Keep a reference to the "close" line in the data[0] dataseries
        self.order = None
        self.data_close = self.datas[0].close
        self.data_open = self.datas[0].open

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
        prev_dif, cur_dif, next_dif = np.NAN, np.NAN, np.NAN
        prev_dea, cur_dea, next_dea = np.NAN, np.NAN, np.NAN
        if len(self.data_open) >= 1:
            next_open = self.data_open[0]

            v = str(self.datas[0].datetime.date(0))
            x = pd.to_datetime(v, errors="coerce")
            cur_index = None
            try:
                cur_index = stock_qfq_df.index.get_loc(x)
            except Exception as e:
                print(e)
            if cur_index is None:
                return
            try:
                cur_dif, cur_dea = stock_qfq_df.iloc[cur_index][["dif", "dea"]]
                prev_dif, prev_dea = stock_qfq_df.iloc[cur_index - 1][["dif", "dea"]]
                next_dif, next_dea = stock_qfq_df.iloc[cur_index + 1][["dif", "dea"]]
            except Exception as e:
                print(e)
                return

        # self.log(f"RANGE , close:{today_close}")
        is_ok = daw.is_in_time_list(self.time_list, str(dt))
        if self._size == 0:
            if is_ok and next_open is not np.NAN:
                if tool.bo_gu(prev_dif, cur_dif, next_dif) or tool.jin_cha(prev_dif, prev_dea, cur_dif, cur_dea):
                    cash = np.round(self.broker.getcash())
                    size = int((cash / (next_open * 100))) * 100
                    self.order = self.buy(size=size)  # 执行买入
                    self._size = size
                    self.log(f"买入 , cash:{cash} ,size:{size} ,clsoe:{next_open}")
        else:
            if not is_ok and next_open is not np.NAN:
                # 这里不用波峰判定效果更好一些
                if tool.bo_fen(prev_dif, cur_dif, next_dif):
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
    stock_qfq_df = tool.modify(stock_qfq_df)

    # Create a Data Feed
    # TODO 回测时间修改
    data = bt.feeds.PandasData(
        dataname=stock_qfq_df,
        fromdate=datetime.datetime(2008, 1, 1),
        todate=datetime.datetime(2024, 3, 21),
    )
    cerebro.adddata(data)
    cerebro.broker.setcash(10000.0)

    print("Starting Portfolio Value: %.2f" % cerebro.broker.getvalue())
    cerebro.run()
    print("Final Portfolio Value: %.2f" % cerebro.broker.getvalue())
