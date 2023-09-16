package noSql

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func InitRedis() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.50.222:6379",
		Password: "",
		DB:       0,
	})

	_, err = rdb.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return err
}

func WriteKeyWithName(ctx context.Context, key, value string) error {
	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		return fmt.Errorf("redis 设置 kv 失败: %w", err)
	}
	val, verr := rdb.Get(ctx, key).Result()
	if verr != nil {
		return fmt.Errorf("读取失败: %w", verr)
	}
	fmt.Println(val)
	return nil
}

func GetValueWithKey(k string, ctx context.Context) {
	val, err := rdb.Get(ctx, k).Result()
	if err != nil {
		fmt.Println("取出数据错误:", err)
		return
	}
	fmt.Println(val)
}
func GetHash(ctx context.Context) {
	r, e := rdb.HGet(ctx, "user", "name1").Result()
	if e != nil {
		fmt.Println("hash error :", e)
		return
	}
	fmt.Println(r)
	rs, se := rdb.HGetAll(ctx, "user").Result()
	if se != nil {
		fmt.Println("getall ", se)
		return
	}
	fmt.Println(rs)
}

func GetZSet(ctx context.Context) {
	zKey := "zKey"
	//lan := []*redis.Z{
	//	&redis.Z{Score: 92, Member: "Golang"},
	//	&redis.Z{Score: 91, Member: "Python"},
	//	&redis.Z{Score: 93, Member: "Java"},
	//	&redis.Z{Score: 91, Member: "C+"},
	//}
	//num, err := rdb.ZAdd(ctx, zKey, lan...).Result()
	result, err := rdb.ZRangeByScore(ctx, zKey, &redis.ZRangeBy{
		Min:    "90",
		Max:    "100",
		Offset: 0,
		Count:  0,
	}).Result()
	if err != nil {
		fmt.Println("z set err:", err)
		return
	}
	fmt.Println(result)

}
