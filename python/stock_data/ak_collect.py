import talib
import pandas as pd
import numpy as np
import mplfinance as mpf
import akshare as ak
from comm import *

pd.set_option("display.max_columns", None)
pd.set_option("display.width", 500)


def is_jc(_dif, _dea, _close):  # 金叉
    if _dif >= _dea:
        return _close * 0.95
    return np.nan


def is_sc(_dif, _dea, _close):  # 金叉
    if _dif >= _dea:
        return np.nan
    return _close * 1.15


if __name__ == "__main__":
    df = ak.stock_zh_index_daily_em(symbol="sz002424")
    df = collect_data_by_df(df)
    df.dropna(inplace=True, axis=0)
    # (date, _, _, _, close, dif, dea, macd, ma5, ma10, ma20, ma60) = get_params(df, len(df) - 1)
    # print(df.head(3))
    # print(df.tail(3))
    x = get_params(df, len(df) - 1)
    df = df.iloc[len(df) - 200 : len(df) - 1]
    # print(x)
    # print(df.tail(10))
    df.to_excel("opt.xlsx", index=False, startrow=3, startcol=1)

    my_color = mpf.make_marketcolors(
        up="red", down="green", edge="inherit", volume="inherit"
    )
    my_stype = mpf.make_mpf_style(marketcolors=my_color)

    datetime_series = pd.to_datetime(df["date"])
    datetime_index = pd.DatetimeIndex(datetime_series.values)

    buy = df.apply(lambda x: is_jc(x["dif"], x["dea"], x["close"]), axis=1)
    sellout = df.apply(lambda x: is_sc(x["dif"], x["dea"], x["close"]), axis=1)

    # buy = (
    #    np.where(
    #        (df["dif"] > df["dif"].shift(-1))
    #        & (df["close"].shift(1) < df["open"].shift(1)),
    #        1,
    #        np.nan,
    #    )
    #    * 0.95
    #    * df["low"]
    # )

    add_plot = [
        mpf.make_addplot(df[["ma5", "ma10", "ma20", "ma60"]]),
        # mpf.make_addplot(df['signal_long'], scatter=True, markersize=5, marker="^", color='r'),
        # mpf.make_addplot(df['signal_short'], scatter=True, markersize=5, marker="s", color='g')
        # mpf.make_addplot(buy, scatter=True, markersize=50, marker=r'$\Uparrow$', color='green')
        mpf.make_addplot(
            buy, scatter=True, markersize=50, marker=r"$\Uparrow$", color="red"
        ),
        mpf.make_addplot(
            sellout, scatter=True, markersize=50, marker=r"$\Downarrow$", color="green"
        ),
    ]

    df2 = df.set_index(datetime_index)
    # df['date'].astype(pd.DatetimeIndex)
    # print(df2.info())
    mpf.plot(
        df2,
        type="candle",
        ylabel="price",
        style=my_stype,
        addplot=add_plot,
        volume=True,
        ylabel_lower="vol",
    )
    # mpf.plot(df2, type='line', ylabel='price', style=my_stype)
