package ch11

import (
	"errors"
	"fmt"
	"image"
	"os"
	"testing"
	"time"
)

func TestString(t *testing.T) {
	//str := `hello \r \n world`

	s := "pprof.cn博客"
	//println(len(s))

	for index, r := range s {
		println(index)
		fmt.Printf("%v(%c) ", r, r)
	}

	//println(str)
}

type student struct {
	name string
	age  int
}

func TestStruct(t *testing.T) {
	m := make(map[string]*student)
	stus := []student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}

type person struct {
	name string
	city string
	age  int8
}

/* go中自己实现构造函数 */
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

//Dream Person做梦的方法
func (p person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

func TestConstruct(t *testing.T) {
	p9 := newPerson("pprof.cn", "测试", 90)
	fmt.Printf("%#v\n", p9)

	p9.Dream()
}

func TestCopy(t *testing.T) {
	array := []int{10, 20}
	slice := make([]int, 6)
	n := copy(slice, array)
	fmt.Println(n, slice)
}

func TestCondition(t *testing.T) {
	if a := 1; a > 3 {
		println(a)
	} else if a == 1 {
		println(a)
	}
}

// 返回2个函数类型的返回值 f1,f2不是一个闭包
func test01(base int) (func(int) int, func(int) int) {
	// 定义2个函数，并返回
	// 相加
	add := func(i int) int {
		base += i
		return base
	}
	// 相减
	sub := func(i int) int {
		base -= i
		return base
	}
	// 返回
	return add, sub
}

func TestForTest01(t *testing.T) {
	f1, f2 := test01(10)
	// base一直是没有消
	fmt.Println(f1(1), f2(2))
	// 此时base是9
	fmt.Println(f1(3), f2(4))
}

func add(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

func TestForAdd(t *testing.T) {
	tmp1 := add(10)
	fmt.Println(tmp1(1), tmp1(2))
	// 此时tmp1和tmp2不是一个实体了
	tmp2 := add(100)
	fmt.Println(tmp2(1), tmp2(2))
}

type Test struct {
	name string
}

//对比代码1，defer使用闭包的方式 由于闭包用到的变量 i 在执行的时候已经变成4,所以输出全都是4
func TestTest(t *testing.T) {
	var whatever [5]struct{}

	for i := range whatever {
		defer func() { fmt.Println(i) }()
	}
}

//对比代码2 defer语句中的变量，在defer声明时就决定了。
func TestTest1(t *testing.T) {
	var whatever [5]struct{}

	for i := range whatever {
		defer fmt.Println(i)
	}

}

func (t *Test) Close() {
	fmt.Println(t.name, "closed")
}

func Close(t Test) {
	t.Close()
}

func Test1(t *testing.T) {
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		defer t.Close()
	}
}

func Test2(t *testing.T) {
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		defer Close(t)
	}
}

func Test3(t *testing.T) {
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		t2 := t
		defer t2.Close()
	}
}

func foo(a, b int) (i int, err error) {
	//这种方式 就是非defer closure的方式 如果 defer 后面跟的不是一个 closure 最后执行的时候我们得到的并不是最新的值。
	defer fmt.Printf("first defer err %v\n", err)
	//延迟调用参数在注册时求值或复制，可用指针或闭包 "延迟" 读取。
	defer func(err error) { fmt.Printf("second defer err %v\n", err) }(err)
	//这种方式 就是defer closure的方式
	defer func() { fmt.Printf("third defer err %v\n", err) }()
	if b == 0 {
		err = errors.New("divided by zero!")
		return
	}

	i = a / b
	return
}

func TestForFoo(t *testing.T) {
	_, _ = foo(2, 0)
}

func foo1(a, b int) (i int, err error) {
	defer func() {
		fmt.Printf("first defer err %v\n", err)
	}()

	defer func(err error) { fmt.Printf("second defer err %v\n", err) }(err)
	defer func() { fmt.Printf("third defer err %v\n", err) }()
	if b == 0 {
		err = errors.New("divided by zero!")
		return
	}

	i = a / b
	return
}

func TestForFoo1(t *testing.T) {
	_, _ = foo1(2, 0)
}

type User struct {
	id   int
	name string
}

func (self *User) Test() {
	fmt.Printf("%p, %v\n", self, self)
}

func TestForUseTest(t *testing.T) {
	u := User{1, "Tom"}
	u.Test()

	mValue := u.Test
	mValue() // 隐式传递 receiver

	mExpression := (*User).Test
	mExpression(&u) // 显式传递 receiver
}

type PathError struct {
	path       string
	op         string
	createTime string
	message    string
}

func (p *PathError) Error() string {
	return fmt.Sprintf("path=%s \nop=%s \ncreateTime=%s \nmessage=%s", p.path,
		p.op, p.createTime, p.message)
}

func Open(filename string) error {

	file, err := os.Open(filename)
	if err != nil {
		return &PathError{
			path:       filename,
			op:         "read",
			message:    err.Error(),
			createTime: fmt.Sprintf("%v", time.Now()),
		}
	}

	defer file.Close()
	return nil
}

func TestForError(t *testing.T) {
	err := Open("/Users/5lmh/Desktop/go/src/test.txt")
	switch v := err.(type) {
	case *PathError:
		fmt.Println("get path error,", v)

		println(v)
	default:

	}
}

func TestTimer(t *testing.T) {
	timer5 := time.NewTimer(3 * time.Second)
	timer5.Reset(1 * time.Second)
	fmt.Println(time.Now())
	fmt.Println(<-timer5.C)
}

var icons map[string]image.Image

func loadIcons() {
	icons = map[string]image.Image{
		"left":  loadIcon("left.png"),
		"up":    loadIcon("up.png"),
		"right": loadIcon("right.png"),
		"down":  loadIcon("down.png"),
	}
}

func TestIconFor_test(t *testing.T) {
	if icons == nil {
		loadIcons()
	}
}
