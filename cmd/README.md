#### 非常简单的文件服务器 (

- 配置文件含有程序运行的端口和文件所在路径
- 因为有时候我们所处可能连接的是 wifi 所以再开始会获取我们电脑当前的所有内部 ip 地
  址,包含虚拟机,docker,等一些 ip,如果你不能使用 ifconfig 或者 ipconfig 获得自己的 ip
  那么程序中给出的 ip 地址肯定有你需要的 ip 地址

- 使用: go run file_srv.go 然后访问你的 ip 加上端口号 例如 127.0.0.1:8888 ,端口可以
  配置文件中指定, 需要注意的是 windows and linux 配置文件路径分隔符不一样

#### get_coe program

- 简单的批量 go get 代码的小玩意
- 要 get 的代码通过 json 配置，还有一点别的必备参数，然后你就可以啥都不管了
- 原因：因为 有时候初始化配置，会需要 go get 一些代码，一个一个来有点累，并且需要不断的查看是不是完成，所以有了这个小玩意

#### **jjson**

- 格式化目录中的 json 或者单个 json 文件
- jjson usege
- jjson -f xxx.json
- jjson -d xxx(dir)
