package redis

import (
	"fmt"
	"log"
	"testing"

	"github.com/gomodule/redigo/redis"
)

var rc = &RedConfig{
	Net:         "tcp",
	Addr:        "127.0.0.1:6379",
	Pass:        "123456",
	MaxIdle:     10,
	MaxActive:   10,
	IdleTimeout: 60,
	Wait:        false,
}

func TestInit(t *testing.T) {
	Init(rc)

	const (
		expire = 3600
	)

	conn := redisCli.Get()
	defer conn.Close()

	key := "go"
	_, err := conn.Do("SET", key, "1", "EX", expire)
	if err != nil {
		log.Fatal(err)
	}
	val, err := redis.String(conn.Do("GET", key))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("val:", val)
	conn.Do("DEL", key)
}
