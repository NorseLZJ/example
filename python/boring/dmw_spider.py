"""
大麦网的活动简单获取
"""

import time
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver import ActionChains
import pandas as pd

URL = "https://search.damai.cn/search.html?keyword=&spm=a2oeg.home.searchtxt.dsearchbtn2"

options = webdriver.ChromeOptions()
options.add_argument(
    "user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"
)

driver = webdriver.Chrome(executable_path="chromedriver.exe", options=options)
driver.get(URL)
time.sleep(1)

citys, titles, addresss, dates = [], [], [], []
try:
    es = driver.find_elements(by=By.CLASS_NAME, value="items")
    while True:
        e = es.pop()
        if e is not None:
            title = e.find_element(by=By.CLASS_NAME, value="items__txt")
            ss = str(title.text).split("\n")
            a = ss[0].split(" ")
            citys.append(a[0])
            titles.append(a[1])
            addresss.append(ss[2])
            dates.append(ss[3])
            print(len(ss), ss)
        else:
            break
    if len(citys) != 0:
        df = pd.DataFrame(data={"city": citys, "title": titles, "address": addresss, "dates": dates})
        df.to_excel("dmw.xlsx", index=False)
except Exception as e:
    if len(citys) != 0:
        df = pd.DataFrame(data={"city": citys, "title": titles, "address": addresss, "dates": dates})
        df.to_excel("dmw.xlsx", index=False)
    print("出错了:", e)

driver.quit()
