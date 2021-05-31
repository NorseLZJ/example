"""
main is all program start
"""

# -*- coding: UTF-8 -*-

import urllib3
import os
import re
import time
from bs4 import BeautifulSoup
from selenium import webdriver
from pyquery import PyQuery as pq
from xml.sax import saxutils as su

header = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36",
    "Content-Type": "application/javascript"
}

# urls
stock_list = 'http://www.sse.com.cn/market/stockdata/dividends/dividend/'
info = 'http://www.sse.com.cn/assortment/stock/list/info/company/index.shtml'
fen_hong = 'http://www.sse.com.cn/assortment/stock/list/info/profit/index.shtml'

browser = None


def open_browser():
    option = webdriver.ChromeOptions()
    option.add_argument('headless')
    global browser
    browser = webdriver.Chrome('./chromedriver.exe', options=option)


def open_info(val):
    url = format("%s%s" % (info, val))
    browser.get(url)
    html = browser.page_source
    data = str(pq(html))
    print(data)


def open_fen_hong(val):
    val = val.split('&')[0]
    url = format("%s%s" % (fen_hong, val))
    browser.get(url)
    html = browser.page_source
    data = str(pq(html))
    print(data)


def open_list(url):
    # print(url)
    browser.get(url)
    html = browser.page_source
    data = str(pq(html))
    data = su.unescape(data)
    re_rule = r'a href="/assortment/stock/list/info/company/index.shtml(.*?)</a>'
    datalist = re.findall(re_rule, data, re.S)
    if len(datalist) <= 1:
        exit(1)
    for val in datalist:
        val = val.split('"')[0]
        val = val.replace('amp;', '')
        # open_info(val, browser)
        open_fen_hong(val)


if __name__ == '__main__':
    open_browser()
    open_list(stock_list)
