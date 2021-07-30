package main

import (
	"errors"

	"github.com/davyxu/cellnet"
	"github.com/mediocregopher/radix.v2/redis"
)

type RedisCmd struct {
	rdbPool cellnet.RedisPoolOperator
}

var (
	rcmd *RedisCmd
)

func init() {
	rcmd = &RedisCmd{}
}

func SetClient(pool cellnet.RedisPoolOperator) {
	if pool == nil {
		panic("redis pool is nil")
	}
	rcmd.rdbPool = pool
}

func (rc *RedisCmd) client() (*redis.Client, bool) {
	client := rcmd.rdbPool.Operate(func(c interface{}) interface{} {
		if c == nil {
			log.Infof("redis client is nil")
			return nil
		}
		exec, ok := c.(*redis.Client)
		if !ok {
			log.Infof("client is not redis")
			return nil
		}
		return exec
	})

	if client == nil {
		return nil, false
	}
	return client.(*redis.Client), true
}

func (rc *RedisCmd) Set(k, v interface{}) error {
	if k == nil || v == nil {
		return errors.New("param is nil")
	}
	client, ok := rc.client()
	if !ok || client == nil {
		return errors.New("client nil")
	}
	if r := client.Cmd("SET", k, v); r.Err != nil {
		return r.Err
	}
	return nil
}

func (rc *RedisCmd) MSet(args ...interface{}) error {
	if len(args)%2 != 0 { // must be even numbers
		return errors.New("args error")
	}
	client, ok := rc.client()
	if !ok || client == nil {
		return errors.New("client nil")
	}
	if r := client.Cmd("MSET", args...); r.Err != nil {
		return r.Err
	}
	return nil
}
