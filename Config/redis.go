package Config

import (
	"time"
	"github.com/gomodule/redigo/redis"
)

var cnt = 0

func ConnRedis() redis.Conn {
	const p = "redis:6379"
	c, err := redis.Dial("tcp", p)
	if err != nil {
		time.Sleep(time.Second)
		cnt++
		if cnt > 30 {
			panic(err)
		}
		ConnRedis()
	}
	return c
}