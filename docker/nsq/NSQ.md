#### **Docker nsqd**

* 把 nsq 整了一下，用了下 tmux 可惜的是不能一键启动，因为我们的nsq.sh执行的时候容器还没起来，所以tmux server 是没起来的,所以脚本执行失败。
* 好处是，如果你的项目比较多，可以多开几个这个容器，指定不同的端口。

* [nsq container](https://hub.docker.com/repository/docker/zijianliunorse/s_nsqd)