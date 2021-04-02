package ch5

import (
	"fmt"
	"testing"
)

type profile struct {
	name string
	age int
	gender string
	mother *profile
	father *profile
}

type company struct {
	companyName string
	companyAddr string
}

type staff struct {
	name string
	age int
	gender string
	position string
	company
}

func (person profile) fmtProfile() {
	fmt.Printf("名字：%s\n", person.name)
	fmt.Printf("年龄：%d\n", person.age)
	fmt.Printf("性别：%s\n", person.gender)
}

/* 结构体的方法如果要改变结构体的属性的话，接受者需要为指针类型
   出于性能的问题，当结构体过大的时候，也是需要指针作为接受者
*/
func (person *profile) increaseAge() {
	person.age += 1
}

func TestStructMethod(t *testing.T) {
	myself := profile{name : "LaoWang", age : 18, gender : "male"}
	myself.fmtProfile()

	myself.increaseAge()
	myself.fmtProfile()
}

func TestStructMethod2(t *testing.T) {
	com := company{
		companyName: "diyi",
		companyAddr: "yuzhou",
	}

	staffInfo := staff{
		name:     "小明",
		age:      28,
		gender:   "男",
		position: "云计算开发工程师",
		company:com,
	}

	fmt.Printf("%s 在 %s 工作\n", staffInfo.name, staffInfo.companyName)
	fmt.Printf("%s 在 %s 工作\n", staffInfo.name, staffInfo.company.companyName)
}