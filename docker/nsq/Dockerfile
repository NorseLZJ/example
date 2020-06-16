FROM alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update \
    && apk upgrade \
    && apk add tmux 

ADD ./nsq /root/nsq
ADD ./.tmux.conf /root/.tmux.conf

#CMD ["/root/nsq/nsq.sh"] # 容器没有启动，tmux server 没有启动，脚本报错
CMD ["top"]
EXPOSE 4150 4151 4160 4161 4170 4171
# docker build -t s_nsqd .
