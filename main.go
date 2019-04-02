package main

import (
	_ "RedisTest/redis"
	"RedisTest/redis"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"RedisTest/model"
	"encoding/json"
)

func main() {
	stringTest()
	HashTest()
	jsonTest()
	listTest()
}

/*
string
批量写入读取

MGET key [key …]
MSET key value [key value …]
 */

func stringTest() {
	fmt.Println("-------------------------------------------------")
	db := db.RedisPool.Get()
	db.Do("select", 1)
	defer db.Close()
	res, err := db.Do("set", "username", "chenliang")
	fmt.Println(res, err)
	res, err = redis.String(db.Do("get", "username"))
	fmt.Println(res, err)

	res, err = db.Do("mset", "username", "chenliang", "age", 12, "address", "北京")
	fmt.Println(res, err)
	mRes, _ := redis.Values(db.Do("mget", "username", "age", "address"))
	var username string
	var age int64
	var address string
	redis.Scan(mRes, &username, &age, &address)
	fmt.Println(username, age, address)
	res, err = db.Do("set", "a1", "value1", "ex", 5) //设置生命周期为5秒
	res, err = redis.String(db.Do("get", "a1"))
	fmt.Println(res, err)

	res, err = db.Do("exists", "a100")
	fmt.Println(res, err)
}

/*
批量写入读取对象(Hashtable)
HMSET key field value [field value …]
HMGET key field [field …]
 */

func HashTest() {
	fmt.Println("-------------------------------------------------")
	db := db.RedisPool.Get()
	defer db.Close()
	res, err := db.Do("hset", "users", "name", "tom")
	fmt.Println("添加", res, err)

	res, err = redis.String(db.Do("hget", "users", "name"))
	fmt.Println("获取", res, err)

	res, err = db.Do("del", "users", "name")
	fmt.Println("删除", res, err)
	res, err = redis.String(db.Do("hget", "users", "name"))
	fmt.Println("获取", res, err)

	db.Do("hmset", "users", "a1", "v1", "a2", "v2")
	ress, errs := redis.Strings(db.Do("hmget", "users", "a1", "a2"))
	fmt.Println(ress, errs)
}

func jsonTest() {
	user := model.User{Id: 0, Name: "tom", Address: "北京", Phone: "110", Sex: 1}
	bs, _ := json.Marshal(user)
	db := db.RedisPool.Get()
	defer db.Close()
	db.Do("set", "json", string(bs))
	res, err := redis.String(db.Do("get", "json"))
	fmt.Println(res, err)
}

func listTest() {

	db := db.RedisPool.Get()
	defer db.Close()
	db.Do("del","list")
	db.Do("lpush", "list", "a1")
	db.Do("lpush", "list", "a2")
	db.Do("lpush", "list", "a3")
	vs, _ := redis.Values(db.Do("lrange", "list", 0, -1))
	for i, v := range vs {
		fmt.Println(i, string(v.([]byte)))
	}

}
