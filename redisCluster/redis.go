package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	redisCon := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"redis-7000:7000", "redis-7001:7001", "redis-7002:7002"},
	})
	result, err := redisCon.Ping(ctx).Result()
	if err != nil {
		fmt.Println("ping err:", err)
	}
	fmt.Println(result)

	test, _ := redisCon.Get(ctx, "test6").Result()
	//test := redisCon.Set(ctx, "test6", 6, redis.KeepTTL).String()
	fmt.Println(test)

	//pipeline := redisCon.Pipeline()
	//pipeline.Set(ctx, "test1", 5, redis.KeepTTL)
	//pipeline.Set(ctx, "test2", 5, redis.KeepTTL)
	//cmds, _ := pipeline.Exec(ctx)
	//fmt.Println(cmds)

	//test := redisCon.Eval(ctx, "redis.call('SET', 'test1', 7);return 1", nil).String()
	//fmt.Println(test)
}
