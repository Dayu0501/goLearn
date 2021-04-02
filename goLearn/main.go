package main

import (
	"encoding/json"
	"fmt"
	"log"
	"myProject/dao"
	"myProject/demo"
	"sync"
	"time"
)

func requestProcess(wg *sync.WaitGroup) {
	data, _ := dao.HttpGet()

	println(len(data))

	account := demo.ReviewData{}

	err := json.Unmarshal([]byte(data), &account)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", account)

	fmt.Println(account.Status)
	fmt.Println(account.Msg)

	//json 处理

	wg.Done()
}

func getRedisDataProcess(wg *sync.WaitGroup) {
	for {
		time.Sleep(1000)
		if v, ok := demo.GetPkgFormRedis(); ok == nil {
			println(len(v))
			//to do
		}

	}
	//wg.Done()
}

func main() {
	obj := dao.Person1{}
	dao.DbSelect()

	fmt.Print("%#v\n", obj)

	var wg sync.WaitGroup
	count := 0
	wg.Add(2)
	go requestProcess(&wg)
	go getRedisDataProcess(&wg)

	wg.Wait()
	fmt.Println("count 的值为：", count)
}
