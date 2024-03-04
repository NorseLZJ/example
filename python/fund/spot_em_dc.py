import akshare as ak

if __name__ == "__main__":
    df = ak.stock_zh_a_spot_em()
    # print(df.head(5))
    df.sort_values(by="60日涨跌幅", inplace=True)
    df.to_excel("spot_em_dc.xlsx", index=False)
