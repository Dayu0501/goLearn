package test

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type Student struct {
	Name  string
	Age   int
	Class string
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Addr string `json:"addr,omitempty"`
}

func TestRefelecTagLearn(t *testing.T) {
	var haiJunHua = Person{"huaHaiJun", 88, "HeNanXinYang"}

	// 三种获取 field
	field, _ := reflect.TypeOf(haiJunHua).FieldByName("Name")
	t.Log(field)

	field, _ = reflect.TypeOf(haiJunHua).FieldByName("Age")
	t.Log(field)

	field, _ = reflect.TypeOf(haiJunHua).FieldByName("Addr")
	t.Log(field)

	t.Log("------------------------------------------------------")

	field = reflect.ValueOf(haiJunHua).Type().Field(0)         // i 表示第几个字段
	t.Log(field)

	field = reflect.ValueOf(&haiJunHua).Elem().Type().Field(1) // i 表示第几个字段
	t.Log(field)

	field = reflect.ValueOf(&haiJunHua).Elem().Type().Field(2) // i 表示第几个字段
	t.Log(field)

	// 获取 Tag
	var tag = field.Tag
	t.Log(tag)

	// 获取键值对
	labelValue := tag.Get("label")
	t.Log(labelValue)

	labelValue, _ = tag.Lookup("label")
	t.Log(labelValue)
}

func TestStructTagLearn(t *testing.T) {
	p1 := Person{
		Name: "Jack",
		Age:  22,
	}

	data1, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}

	// p1 没有 Addr，就不会打印了
	fmt.Printf("%s\n", data1)

	// ================

	p2 := Person{
		Name: "Jack",
		Age:  22,
		Addr: "China",
	}

	data2, err := json.Marshal(p2)
	if err != nil {
		panic(err)
	}

	// p2 则会打印所有
	fmt.Printf("%s\n", data2)
}

func newStudent(age int, name string, class string) *Student {
	return &(Student{name, age, class})
}

func TestStructLearn(t *testing.T) {
	var stu Student
	stu.Age = 1
	stu.Name = "Super Man"
	stu.Class = "class three two"

	t.Log(stu.Name)

	stu2 := newStudent(2, "Super girl", "class three one")

	t.Log(stu2.Name)

}
