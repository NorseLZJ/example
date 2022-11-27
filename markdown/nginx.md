# nginx 安装和 upstream 配置

## 安装

- 系统

> centos7.9

> 相关命令

```shell
yum install -y pcre zlib openssl
yum install -y pcre-devel openssl-devel zlib-devel

./configure
make
make install
```

## upstream 配置

```nginx
worker_processes  1;
events {
    worker_connections  1024;
}
http {
    include       mime.types;
    default_type  application/octet-stream;
    sendfile        on;
    keepalive_timeout  65;
    upstream a {
        server 192.168.10.160:8000 weight=1;
        server 192.168.10.160:8001 weight=1;
    }

    upstream b {
        server 192.168.10.160:8000 weight=1;
    }

    server {
        listen       80;
        server_name  localhost;
        location / {
            root   html;
            index  index.html index.htm;
        }
        location /a/ {
            proxy_pass http://a/;
            # 结尾的斜杠 / 代表截断 /a/ 向后传递 ，不带的话 /a 向后传递
            #  /a/api/user
            #  /api/user
        }
        location /b/ {
            proxy_pass http://b/;
        }
        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
            root   html;
        }
    }
}
```

## 调用示例
