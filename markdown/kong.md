# kong install

**kong 和 konga docker 方式安装**

```shell
docker network create kong-net
```

```shell
docker run -d --name kong-database \
                --network=kong-net \
                -p 5432:5432 \
                -e "POSTGRES_USER=kong" \
                -e "POSTGRES_DB=kong" \
                -e "POSTGRES_PASSWORD=kong" \
                postgres:9.6
```

```shell
 docker run --rm \
    --network=kong-net \
    --link kong-database:kong-database \
    -e "KONG_DATABASE=postgres" \
    -e "KONG_PG_HOST=kong-database" \
    -e "KONG_PG_USER=kong" \
    -e "KONG_PG_PASSWORD=kong" \
    -e "KONG_CASSANDRA_CONTACT_POINTS=kong-database" \
    kong kong migrations bootstrap
```

```shell
docker run -d --name kong \
    --network=kong-net \
    --link kong-database:kong-database \
    -e "KONG_DATABASE=postgres" \
    -e "KONG_PG_HOST=kong-database" \
    -e "KONG_PG_PASSWORD=kong" \
    -e "KONG_CASSANDRA_CONTACT_POINTS=kong-database" \
    -e "KONG_PROXY_ACCESS_LOG=/dev/stdout" \
    -e "KONG_ADMIN_ACCESS_LOG=/dev/stdout" \
    -e "KONG_PROXY_ERROR_LOG=/dev/stderr" \
    -e "KONG_ADMIN_ERROR_LOG=/dev/stderr" \
    -e "KONG_ADMIN_LISTEN=0.0.0.0:8001, 0.0.0.0:8444 ssl" \
    -p 8000:8000 \
    -p 8443:8443 \
    -p 8001:8001 \
    -p 8444:8444 \
    kong
```

# konga install

```shell
docker run -d --name konga-database \
                --network=kong-net \
                -p 5433:5432 \
                -e "POSTGRES_USER=konga" \
                -e "POSTGRES_DB=konga" \
                -e "POSTGRES_PASSWORD=konga" \
                postgres:9.6
```

```shell
docker run --rm --network=kong-net \
    pantsel/konga:latest -c prepare -a postgres -u postgres://konga:konga@konga-database:5432/konga

```

```shell
docker run -d -p 1337:1337 \
             --network=kong-net \
             -e "DB_ADAPTER=postgres" \
             -e "DB_URI=postgres://konga:konga@konga-database:5432/konga" \
             -e "NODE_ENV=production" \
             --name konga \
             pantsel/konga
```

[原始文档](https://github.com/NorseLZJ/example/blob/master/markdown/kong.md)
