package ch6

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInterface(t *testing.T) {
	/* (int)代表的是类型，(25)代表的是值 */
	age := (int)(25)
	//或者使用
	age1 := (interface{})(25)

	fmt.Printf("type: %T, data: %v \n", age, age)
	fmt.Printf("type: %T, data: %v \n", age1, age1)
}

/*func testMethod() (*File, error) {
	var reader io.Reader

	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}

	reader = tty
}*/

type Person1 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Addr string `json:"addr,omitempty"`
}

func TestMethod(t *testing.T) {

	obj := Person1{
		Name: "Jack",
		Age:  22,
		Addr: "China",
	}

	// 三种获取 field
	field, _ := reflect.TypeOf(obj).FieldByName("Age")
	println(field.Name)
	field = reflect.ValueOf(obj).Type().Field(0)         // i 表示第几个字段
	println(field.Name)
	field = reflect.ValueOf(&obj).Elem().Type().Field(1) // i 表示第几个字段
	println(field.Name)

	// 获取 Tag
	tag := field.Tag
	println(tag)

	// 获取键值对
	labelValue := tag.Get("label")
	println(labelValue)

	labelValue,ok := tag.Lookup("label")
	println(ok)
}

type Person struct {
	Name        string `label:"Name is: "`
	Age         int    `label:"Age is: "`
	Gender      string `label:"Gender is: " default:"unknown"`
}

func TestTag(t *testing.T) {
	person := Person{
		Name:        "MING",
		Age:         29,
	}

	_ = Print(person)
}

func Print(obj interface{}) error {
	// 取 Value
	v := reflect.ValueOf(obj)

	// 解析字段
	for i := 0; i < v.NumField(); i++ {

		// 取tag
		field := v.Type().Field(i)
		tag := field.Tag

		// 解析label 和 default
		label := tag.Get("label")
		defaultValue := tag.Get("default")

		value := fmt.Sprintf("%v", v.Field(i))
		if value == "" {
			// 如果没有指定值，则用默认值替代
			value = defaultValue
		}

		fmt.Println(label + value)
	}

	return nil
}