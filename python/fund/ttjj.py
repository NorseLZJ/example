import requests
from bs4 import BeautifulSoup

url = "https://fund.eastmoney.com/003384.html?spm=search"
headers = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36"
}

response = requests.get(url, headers=headers)
soup = BeautifulSoup(response.text, "html.parser")

# 获取基金名称
fund_name = soup.select_one(".fundDetail-tit").text.strip()
print(soup.select_one(".fundDetail-tit").text)

print("基金名称:", fund_name)

# 获取基金净值
net_value = soup.select_one(".dataItem01 .dataNums").text.strip()
print("基金净值:", net_value)

# 获取基金涨跌幅
change_percent = soup.select_one(".dataItem02 .dataNums").text.strip()
print("基金涨跌幅:", change_percent)

# 获取基金规模
fund_size = soup.select_one(".dataItem03 .dataNums").text.strip()
print("基金规模:", fund_size)

# 获取基金经理
# fund_manager = soup.select_one(".managerteam").text.strip()
# print("基金经理:", fund_manager)
