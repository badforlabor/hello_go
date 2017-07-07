package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"encoding/json"

	"github.com/garyburd/redigo/redis"
)

// 暂时就封装三种数据类型：bytes, string, int，提供了Get/Set方法
// 如果需要其他类型的，先用json转成string格式，这样，这样的目的是，在数据库中观测。
// redis默认端口是6379，需要更改，自己手动改一下。

var (
	Pool *redis.Pool
)

func init() {
	// 用默认端口了
	redisHost := ":6379"
	Pool = newPool(redisHost)
	close()
}

func newPool(server string) *redis.Pool {

	return &redis.Pool{

		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// 监听关闭消息
func close() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}

func redisGetBytes(key string) ([]byte, error) {

	conn := Pool.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error get key %s: %v", key, err)
	}
	return data, err
}
func redisGetString(key string) (string, error) {

	conn := Pool.Get()
	defer conn.Close()

	data, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error get key %s: %v", key, err)
	}
	return data, err
}
func redisGetInt(key string) (int, error) {

	conn := Pool.Get()
	defer conn.Close()

	data, err := redis.Int(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error get key %s: %v", key, err)
	}
	return data, err
}

func redisSetBytes(key string, value []byte) {

	conn := Pool.Get()
	defer conn.Close()

	conn.Do("SET", key, value)
}
func redisSetString(key string, value string) {

	conn := Pool.Get()
	defer conn.Close()

	conn.Do("SET", key, value)
}
func redisSetInt(key string, value int) {

	conn := Pool.Get()
	defer conn.Close()

	conn.Do("SET", key, value)
}
func redisSetJSON(key string, v interface{}) {
	jsonbytes, err := json.Marshal(v)

	if err != nil {
		fmt.Println("json format error.")
	} else {
		redisSetBytes(key, jsonbytes)
	}
}
func redisGetJSON(key string, v interface{}) error {
	jsonbytes, err := redisGetBytes(key)
	if err != nil {
		fmt.Println("get json data error.")
	} else {
		//fmt.Println(string(jsonbytes))
		err = json.Unmarshal(jsonbytes, v)
	}
	return err
}
func redisAutoID() int {

	conn := Pool.Get()
	defer conn.Close()

	n, _ := redis.Int(conn.Do("INCR", "AutoID"))
	return n
}

type jsonStructTest struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
}

func main() {
	redisSetString("test", "test-value")
	test, err := redisGetString("test")
	fmt.Println(test, err)

	redisSetInt("test-int", 123)
	i, err := redisGetInt("test-int")
	fmt.Println(i, err)

	oribytes := []byte("bytes-data")
	redisSetBytes("test-bytes", oribytes)
	bytes, err := redisGetBytes("test-bytes")
	fmt.Println(bytes, oribytes, err)

	orijson := jsonStructTest{AccessToken: "temptoken"}
	redisSetJSON("test-json", &orijson)

	var getjson jsonStructTest
	err = redisGetJSON("test-json", &getjson)
	fmt.Println(getjson)

	fmt.Println("auto-id", redisAutoID())
}
