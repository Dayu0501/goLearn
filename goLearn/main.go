package main

import (
	"fmt"
	"myProject/dao"
)

func main() {
	obj := dao.Person1{}
	dao.DbSelect()

	fmt.Print("%#v",obj)
}