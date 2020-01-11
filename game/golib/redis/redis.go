package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

var redisCli *redis.Pool

func Init(t *RedConfig) {
	if redisCli != nil {
		return
	}
	redisCli = &redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			c, err := redis.Dial(t.Net, t.Addr)
			if err != nil {
				return nil, err
			}
			_, err = c.Do("AUTH", t.Pass)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		DialContext:  nil,
		TestOnBorrow: nil,
		MaxIdle:      t.MaxIdle,
		MaxActive:    t.MaxActive,
		IdleTimeout:  time.Second * 10,
		Wait:         t.Wait,
	}
}
