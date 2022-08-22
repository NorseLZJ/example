from configparser import Interpolation
import requests
import json
import pandas as pd


pd.set_option("display.max_columns", None)
pd.set_option("display.width", 500)

url = "https://voice.baidu.com/api/newpneumonia?from=page"
care_city = {"陕西": {"西安": 1}, "甘肃": {"庆阳": 1}, "上海": {"浦东新区": 1, "杨浦区": 1}}

if __name__ == "__main__":

    headers = {
        "user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36"
    }
    resp = requests.get(url, headers=headers)
    v = json.loads(resp.text)
    with open("temp.json", "w", encoding="utf8") as f:
        f.write(resp.text)
    caseList = v["data"]["caseList"]
    city_list = []
    add_bt_list = []
    add_wzz_list = []
    dangers_list = []
    for val in caseList:
        exist = care_city.get(val["area"])
        if exist is None:
            continue

        for city in val["subList"]:
            exist2 = exist.get(city["city"])
            if exist2 is None:
                continue
            city_list.append(city["city"])
            add_bt_list.append(city["confirmedRelative"])
            add_wzz_list.append(city["asymptomaticRelative"])
            dangers_list.append(len(city["dangerousAreas"]["subList"]))
            break
    df = pd.DataFrame(
        data={
            "city": city_list,
            "本土": add_bt_list,
            "无症状": add_wzz_list,
            "风险地区": dangers_list,
        }
    )
    print(df)
