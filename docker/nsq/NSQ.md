#### **Docker nsqd**

* 把 nsq 整了一下，用了下 tmux 可惜的是不能一键启动，因为我们的nsq.sh执行的时候容器还没起来，所以tmux server 是没起来的,所以脚本执行失败。
* 好处是，如果你的项目比较多，可以多开几个这个容器，指定不同的端口。

* [nsq container](https://hub.docker.com/repository/docker/zijianliunorse/s_nsqd)
* [Dockerfile nsq container](https://github.com/NorseLZJ/example/tree/master/docker/nsq)

#### **使用**
* [nsq download](https://nsq.io/deployment/installing.html)
* 先进下载页面，下载 linux版本 解压，把解压的目录里边bin目录的东西拷贝到 nsq 目录，有nsq.sh 的目录里边
* 执行 docker build -t s_nsqd . 
* 就完成了