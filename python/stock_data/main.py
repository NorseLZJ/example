"""
stock date get from network
"""
import akshare as ak

'''
Document 
https://zhuanlan.zhihu.com/p/393083394
https://www.akshare.xyz/tutorial.html
'''
if __name__ == '__main__':
    '''
    sh 上海600
    sz 深圳00，创业300
    每次取最后[4,8]个就行了，对比下没有的添加，然后触发更新，
    均线，买卖操作等
    '''
    df = ak.stock_zh_a_minute('sz002424', '60')
    size = len(df)
    for idx in range(len(df)):
        size -= 1
        if size <= 8:
            print(df['day'][idx], ' ', df['close'][idx])
