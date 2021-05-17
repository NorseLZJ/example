# coding:utf-8

import requests, json, sys
from urllib import parse

center_url = ''

# url param
headers = {"Content-Type": "application/x-www-form-urlencoded"}

# opt
operator = '/operator'

# config
server_list_src = ""
game_list_src = ""
recharge_info_src = ""


def init_argc():
    file = sys.argv[1]
    print(file)
    with open(file, encoding='utf-8') as f:
        d = json.loads(f.read())
        global center_url, server_list_src, game_list_src, recharge_info_src
        center_url = d['center_url']
        server_list_src = d['server_list_src']
        game_list_src = d['game_list_src']
        recharge_info_src = d['recharge_info_src']


def operator_game_list():
    d = json.loads(json.dumps(game_list_src))
    channel = d['channel']
    data = d['data']
    val = {
        "channel": channel,
        "data": json.dumps(data),
    }
    cur_url = format("%s%s%s" % (center_url, operator, '/gamelist'))
    r = requests.post(cur_url, data=val, headers=headers)
    print(r.text)


def operator_server_list():
    d = json.loads(json.dumps(server_list_src))
    channel = d['channel']
    game_id = d['game_id']
    data = d['data']
    val = {
        "channel": channel,
        "game_id": game_id,
        "data": json.dumps(data),
    }
    print(val)
    cur_url = format("%s%s%s" % (center_url, operator, '/serverlist'))
    r = requests.post(cur_url, data=val, headers=headers)
    print(r.text)


def operator_rechargeinfo():
    d = json.loads(json.dumps(recharge_info_src))
    channel = d['channel']
    game_id = d['game_id']
    package_name = d['package_name']
    data = d['data']
    val = {
        "channel": channel,
        "game_id": game_id,
        "package_name": package_name,
        "data": json.dumps(data),
    }
    cur_url = format("%s%s%s" % (center_url, operator, '/rechargeinfo'))
    r = requests.post(cur_url, data=val, headers=headers)
    print(r.text)


if __name__ == '__main__':
    init_argc()
    operator_server_list()
    operator_game_list()
    operator_rechargeinfo()
