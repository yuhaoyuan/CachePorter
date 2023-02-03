package IntelligentRedisKey

import (
	"context"
	"fmt"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"strconv"
	"testing"
	"time"
)

type TestStruct struct {
	Name   string
	Age    int
	Gender string
}

func GetFromDB(ctx context.Context) (interface{}, error) {
	return &TestStruct{
		Name:   "jesse",
		Age:    18,
		Gender: "man",
	}, nil
}

func newClient() *redis.Client {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	p, _ := strconv.Atoi(s.Port())
	addr := fmt.Sprintf("%v:%v", s.Host(), p)
	options := &redis.Options{
		Network: "tcp",
		Addr:    addr,
	}
	client := redis.NewClient(options)
	return client

}
func TestStringKey(t *testing.T) {
	testKey := IntelligentRedisKey{
		Name:       "test_key_%v",
		NameParams: []interface{}{"haoyuan"},

		Expiration: time.Minute,

		GetFromDBFunc: GetFromDB,
		BindData:      TestStruct{},
	}

	redisCli := newClient()

	data, err := testKey.Read(context.Background(), redisCli)
	if err != nil {
		t.Error(err)
	}
	println(data)
}
