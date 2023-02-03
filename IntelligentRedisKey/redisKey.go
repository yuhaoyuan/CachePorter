package IntelligentRedisKey

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// struct 到 string 的读取一体
type RedisKey interface {
	Read(ctx context.Context, client *redis.Client) (interface{}, error)
}

type IntelligentRedisKey struct {
	Name string // key的名字, example_key_{uid}

	NameParams []interface{}

	Expiration       time.Duration                            // 过期时间，0=不过期
	ExpireFunc       func(param ...interface{}) time.Duration // 指定特殊的过期时间
	ExpireFuncParams []interface{}

	GetFromDBFunc func(ctx context.Context) (interface{}, error)

	BindData interface{} // 存在这个key中的数据结构
}

func (c *IntelligentRedisKey) GetFromRedis(ctx context.Context, client *redis.Client) (string, error) {
	key := fmt.Sprintf(c.Name, c.NameParams)
	dataString, err := client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return dataString, nil
}

func (c *IntelligentRedisKey) WriteRedis(ctx context.Context, client *redis.Client) error {
	key := fmt.Sprintf(c.Name, c.NameParams)
	value, err := json.Marshal(c.BindData)
	if err != nil {
		return err
	}

	expiration := c.Expiration
	if c.ExpireFunc != nil {
		expiration = c.ExpireFunc(c.ExpireFuncParams)
	}

	err = client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		// log
		return err
	}
	return nil
}

func (c *IntelligentRedisKey) Read(ctx context.Context, client *redis.Client) (interface{}, error) {
	// read from redis
	dataString, err := c.GetFromRedis(ctx, client)
	if err != nil || dataString == "" {
		// read from db
		dataStruct, err := c.GetFromDBFunc(ctx)
		if err != nil {
			return nil, err
		}
		c.BindData = dataStruct
	}
	// set to redis
	err = c.WriteRedis(ctx, client)
	if err != nil {
		// log
	}
	return c.BindData, err
}
