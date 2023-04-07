"""main start ... """

from scapy.all import *

"""
https://123pan.com:8443
"""


def http_filter(packet):
    if packet.haslayer("TCP") and packet["TCP"].dport == 8443:
        if packet.haslayer("Raw"):
            print(packet["TCP"])


"""
获取网卡名字
wmic nic get AdapterType, Name, Installed, MACAddress, PowerManagementSupported, Speed
"""


if __name__ == "__main__":
    dpkt = sniff(iface="Realtek PCIe GbE Family Controller", count=0, prn=http_filter)  # 这里是针对单网卡的机子，多网卡的可以在参数中指定网卡
