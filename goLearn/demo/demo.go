package demo

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"log"
	"time"
)

var redisdb *redis.Client

func init() {
	//连接服务器
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.70.210:6379", // use default Addr
		Password: "123",               // no password set
		DB:       0,                // use default DB
	})

	//心跳
	pong, err := redisdb.Ping().Result()
	log.Println(pong, err)
}

func SendPkgToRedis(AlgTaskInfoPkgList []AlgTaskInfoReqPkg) error {
	strSendJson, _ := json.Marshal(AlgTaskInfoPkgList)
	 return redisdb.RPush("alg_request_queue", strSendJson).Err()
}

func GetPkgFormRedis() ([]string, error) {
	result, err := redisdb.BLPop(1*time.Second, "alg_response_queue").Result()
	return result, err
}

//func ExampleClient_List() {
//	log.Println("ExampleClient_List +++")
//	defer log.Println("ExampleClient_List ---")
//
//	//添加
//	log.Println(redisdb.RPush("alg_response_queue", "message1").Err())
//	log.Println(redisdb.RPush("alg_response_queue", "message2").Err())
//
//	//设置
//	log.Println(redisdb.LSet("list_test", 2, "message set").Err())
//
//	//remove
//	ret, err := redisdb.LRem("list_test", 3, "message1").Result()
//	log.Println(ret, err)
//
//	rLen, err := redisdb.LLen("list_test").Result()
//	log.Println(rLen, err)
//
//	//遍历
//	lists, err := redisdb.LRange("list_test", 0, rLen-1).Result()
//	log.Println("LRange", lists, err)
//
//	//pop没有时阻塞
//	result, err := redisdb.BLPop(1*time.Second, "list_test").Result()
//	log.Println("result:", result, err, len(result))
//}