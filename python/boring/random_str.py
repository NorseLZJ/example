"""
static , const
"""
import numpy as np

x = [
    "a",
    "b",
    "c",
    "d",
    "e",
    "f",
    "g",
    "h",
    "i",
    "j",
    "k",
    "l",
    "m",
    "n",
    "o",
    "p",
    "q",
    "r",
    "s",
    "t",
    "u",
    "v",
    "w",
    "x",
    "y",
    "z",
    "A",
    "B",
    "C",
    "D",
    "E",
    "F",
    "G",
    "H",
    "I",
    "J",
    "K",
    "L",
    "M",
    "N",
    "O",
    "P",
    "Q",
    "R",
    "S",
    "T",
    "U",
    "V",
    "W",
    "X",
    "Y",
    "Z",
    "0",
    "1",
    "2",
    "3",
    "4",
    "5",
    "6",
    "7",
    "8",
    "9",
]


def rand_str(l=62) -> str:
    v = np.random.randint(1, len(x), len(x))
    ret = (
        x[v[0]]
        + x[v[1]]
        + x[v[2]]
        + x[v[3]]
        + x[v[4]]
        + x[v[5]]
        + x[v[6]]
        + x[v[7]]
        + x[v[8]]
        + x[v[9]]
        + x[v[10]]
        + x[v[11]]
        + x[v[12]]
        + x[v[13]]
        + x[v[14]]
        + x[v[15]]
        + x[v[16]]
        + x[v[17]]
        + x[v[18]]
        + x[v[19]]
        + x[v[20]]
        + x[v[21]]
        + x[v[22]]
        + x[v[23]]
        + x[v[24]]
        + x[v[25]]
        + x[v[26]]
        + x[v[27]]
        + x[v[28]]
        + x[v[29]]
        + x[v[30]]
        + x[v[31]]
        + x[v[32]]
        + x[v[33]]
        + x[v[34]]
        + x[v[35]]
        + x[v[36]]
        + x[v[37]]
        + x[v[38]]
        + x[v[39]]
        + x[v[40]]
        + x[v[41]]
        + x[v[42]]
        + x[v[43]]
        + x[v[44]]
        + x[v[45]]
        + x[v[46]]
        + x[v[47]]
        + x[v[48]]
        + x[v[49]]
        + x[v[50]]
        + x[v[51]]
        + x[v[52]]
        + x[v[53]]
        + x[v[54]]
        + x[v[55]]
        + x[v[56]]
        + x[v[57]]
        + x[v[58]]
        + x[v[59]]
        + x[v[60]]
        + x[v[61]]
    )
    if l >= 62:
        return ret
    return ret[:l]


def rand_str2(l=62) -> str:
    v = np.random.randint(1, len(x), len(x))
    ret = ""
    for i in v:
        ret += x[v[i]]
    if l >= 62:
        return ret
    return ret[:l]
