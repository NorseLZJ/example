from jqdatasdk import *
from redis import *

data_len = 4 * 240 * 3  # 3 year

r = Redis(host='192.168.66.47', port=6381, decode_responses=True)


def cache_data(display_name: str, name: str, start_date: str):
    key = format("%s.%s" % (display_name, name))
    if r.get(key) is not None:
        return
    if not (name[0:2] in ('60', '00', '30')):
        return
    df1 = get_bars(name, data_len, unit='60m', fields=['date', 'open', 'close'], include_now=True)
    r.set(format("%s.%s" % (display_name, name)), df1.to_json())
    print(key)


if __name__ == "__main__":
    auth('xxxxx', 'xxxxx')
    x = get_all_securities(types=['stock'], date=None)
    codes = x['display_name']
    print(type(codes))
    for k, v in codes.items():
        cache_data(v, k, '')
