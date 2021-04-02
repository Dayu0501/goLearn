package ch5

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

type Phone interface {
	call()
}

type nokia struct {
	name string
}

/* nokia 实现Phone接口 */
func (p nokia) call() {
	fmt.Println("我是 Nokia，是一台电话" + p.name)
}

func TestNokia(t *testing.T) {
	nokia1 := nokia{name: "N72"}
	nokia1.call()
}

/* 多态：一个接口的多种形态。
   例如：这就一个接口（老师）下，在不同对象（人）上的不同表现。这就是多态。
*/

//商品接口 go中interface的概念类似的是java中的interface的概念或者是c++中抽象基类的概念
type Good interface {
	settleAccount() int
	orderInfo() string
}

type PhoneNew struct {
	name     string
	quantity int
	price    int
}

type FreeGift struct {
	name     string
	quantity int
	price    int
}

/* PhoneNew实现Good接口 */
func (phone PhoneNew) settleAccount() int {
	return phone.quantity * phone.price
}

func (phone PhoneNew) orderInfo() string {
	return "您要购买" + strconv.Itoa(phone.quantity) + "个" +
		phone.name + "计：" + strconv.Itoa(phone.settleAccount()) + "元"
}

func (gift FreeGift) settleAccount() int {
	return 0
}

func (gift FreeGift) orderInfo() string {
	return "您要购买" + strconv.Itoa(gift.quantity) + "个" +
		gift.name + "计：" + strconv.Itoa(gift.settleAccount()) + "元"
}

func TestInterface3(t *testing.T) {
	iPhone := PhoneNew{
		name:     "iPhone",
		quantity: 1,
		price:    8000,
	}

	earphones := FreeGift{
		name:     "耳机",
		quantity: 1,
		price:    200,
	}

	goods := []Good{iPhone, earphones}
	allPrice := calculateAllPrice(goods)
	fmt.Printf("该订单总共需要支付 %d 元", allPrice)
}

func calculateAllPrice(goods []Good) int {
	var allPrice int
	for _, good := range goods {
		fmt.Println(good.orderInfo())
		allPrice += good.settleAccount()
	}
	return allPrice
}

func TestEmptyInterface(t *testing.T) {
	// 声明一个空接口实例
	var i interface{}

	// 存 int 没有问题
	i = 1
	fmt.Println(i)

	// 存字符串也没有问题
	i = "hello"
	fmt.Println(i)

	// 存布尔值也没有问题
	i = false
	fmt.Println(i)
}

/* 该函数可以接受任意类型的值，接收一个任意类型的值 示例 */
func myfunc(iface interface{}) {
	fmt.Println(iface)
}

/* 接收任意个任意类型的值 示例 */
func myfunc1(ifaces ...interface{}) {
	for _, iface := range ifaces {
		fmt.Println(iface)
	}
}

/* 定义一个可以接收任意类型的 array、slice、map、strcut，例如这边定义一个切片 */
func TestInterface9(t *testing.T) {
	any := make([]interface{}, 5)
	any[0] = 11
	any[1] = "hello world"
	any[2] = []int{11, 22, 33, 44}
	for _, value := range any {
		fmt.Println(value)
	}
}

/* 当空接口承载数组和切片后，该对象无法再进行切片，下面代码会报错 */
func TestEmptyInterface8(t *testing.T) {
	//sli := []int{2, 3, 5, 7, 11, 13}
	//
	//var i interface{}
	//i = sli
	//
	//g := i[1:3]
	//fmt.Println(g)
}

/* 当你使用空接口来接收任意类型的参数时，它的静态类型是 interface{}，但动态类型（是 int，string 还是其他类型）我们并不知道，因此需要使用类型断言。 */
func emptyInterface7(i interface{}) {
	/* 下面代码是类型断言 */
	switch i.(type) {
	case int:
		fmt.Println("参数的类型是 int")
	case string:
		fmt.Println("参数的类型是 string")
	}
}

func TestEmptyInterface7(t *testing.T) {
	a := 10
	b := "hello"

	emptyInterface7(a)
	emptyInterface7(b)

	var c interface{} = 100
	emptyInterface7(c)
}

type iPhone struct {
	name string
}

func (phone iPhone) call() {
	fmt.Println("Hello, iPhone.")
}

func (phone iPhone) send_wechat() {
	fmt.Println("Hello, Wechat.")
}

func TestInterface11(t *testing.T) {
	var phone Phone
	phone = iPhone{name: "apple"}

	phone.call()

	/* 下面函数的调用时错误的，由于静态类型的限制，虽然iPhone实现了send_wechat方法，但是由于是Phone类型的变量，所以不能调用send_wechat方法 */
	/* phone对象显示声明为 Phone 接口类型，因此 phone调用的方法会受到此接口的限制。 */
	//phone.send_wechat()
}

/* Go 语言中的函数调用都是值传递的，变量会在方法调用前进行类型转换。 */
/* 下面代码会报错 */
/*func TestFunc(t *testing.T) {
	a := 10

	switch a.(type) {
	case int:
		fmt.Println("参数的类型是 int")
	case string:
		fmt.Println("参数的类型是 string")
	}
}*/

/* 下面的函数就不会报上面函数的错误, Go 会默默地为我们做一件事，就是把传入函数的参数值（注意：Go 语言中的函数调用都是值传递的）的类型隐式的转换成 interface{} 类型。*/
func printType(i interface{}) {

	switch i.(type) {
	case int:
		fmt.Println("参数的类型是 int")
	case string:
		fmt.Println("参数的类型是 string")
	}
}
func TestFunc1(t *testing.T) {
	a := 10
	printType(a)
}

/* 手动的类型转换，只有静态类型为接口类型的对象才可以进行类型断言*/
func TestFunc2(t *testing.T) {
	a := 10

	/* 手动的类型转换
	var a int = 25
	b := interface{}(a)
	*/

	switch interface{}(a).(type) {
	case int:
		fmt.Println("参数的类型是 int")
	case string:
		fmt.Println("参数的类型是 string")
	}
}

/* 当类型断言完成后，会返回一个静态类型为你断言的类型的对象，也就是说，当我们使用了类型断言，Go 实际上又会默认为我们进行了一次隐式的类型转换。 */
func TestFunc3(t *testing.T) {
	var a interface{} = 25
	fmt.Printf("%T\n", a)

	switch a.(type) {
	case int:
		fmt.Println("int")
		/* 下面这段代码是错误的，因为a已经被隐式类型转换成了int型，不能再一次进行类型断言了 */
		//a.(type)
	case string:
		fmt.Println("string")
	}
}

func TestAssert(t *testing.T) {
	var i interface{} = 10
	t1 := i.(int)
	fmt.Println(t1)

	fmt.Println("=====分隔线=====")

	t2 := i.(string)
	fmt.Println(t2)
}

func TestAssert1(t *testing.T) {
	var i interface{} // nil
	var _ = i.(interface{})
}

func TestAssert2(t *testing.T) {
	var i interface{} = 10
	t1, ok := i.(int)
	fmt.Printf("%d-%t\n", t1, ok)

	fmt.Println("=====分隔线1=====")

	t2, ok := i.(string)
	fmt.Printf("%s-%t\n", t2, ok)

	fmt.Println("=====分隔线2=====")

	var k interface{} // nil
	t3, ok := k.(interface{})
	fmt.Println(t3, "-", ok)

	fmt.Println("=====分隔线3=====")
	k = 10
	t4, ok := k.(interface{})
	fmt.Printf("%d-%t\n", t4, ok)

	t5, ok := k.(int)
	fmt.Printf("%d-%t\n", t5, ok)
}

func TestReflect(t *testing.T) {
	var age interface{} = 18
	//var age int = 18
	fmt.Printf("原始接口变量的类型为 %T，值为 %v \n", age, age)

	tt := reflect.TypeOf(age)
	v := reflect.ValueOf(age)

	fmt.Printf("从接口变量到反射对象：Type对象的类型为 %T \n", tt)
	fmt.Printf("从接口变量到反射对象：Value对象的类型为 %T \n", v)

}

func TestReflect1(t *testing.T) {
	var age interface{} = 25

	fmt.Printf("原始接口变量的类型为 %T，值为 %v \n", age, age)

	tt := reflect.TypeOf(age)
	v := reflect.ValueOf(age)

	// 从接口变量到反射对象
	fmt.Printf("从接口变量到反射对象：Type对象的类型为 %T \n", tt)
	fmt.Printf("从接口变量到反射对象：Value对象的类型为 %T \n", v)

	// 从反射对象到接口变量
	i := v.Interface()
	fmt.Printf("从反射对象到接口变量：新对象的类型为 %T 值为 %v \n", i, i)
}

func TestReflect2(t *testing.T) {
	var name string = "Go编程时光"

	v := reflect.ValueOf(name)
	fmt.Println("可写性为:", v.CanSet())
}

func TestReflect3(t *testing.T) {
	var name string = "Go编程时光"
	v1 := reflect.ValueOf(&name)
	fmt.Println("v1 可写性为:", v1.CanSet())

	v2 := v1.Elem()
	fmt.Println("v2 可写性为:", v2.CanSet())
}

func TestReflect4(t *testing.T) {
	var name string = "Go编程时光"
	fmt.Println("真实世界里 name 的原始值为：", name)

	v1 := reflect.ValueOf(&name)
	v2 := v1.Elem()

	v2.SetString("Python编程时光")
	fmt.Println("通过反射对象进行更新后，真实世界里 name 变为：", name)
}

type Profile struct {
	name   string
	age    int
	gender string
}

func TestReflect5(t *testing.T) {
	m := Profile{}

	tt := reflect.TypeOf(m)
	fmt.Println("Type 的值: ", tt)
	fmt.Printf("Type的类型: %T\n", tt)
	fmt.Println("Kind: ", tt.Kind())
}

func TestReflect6(t *testing.T) {
	m := Profile{}

	tt := reflect.TypeOf(&m)

	fmt.Println("&m Type: ", tt)
	fmt.Println("&m Kind: ", tt.Kind())

	fmt.Println("m Type: ", tt.Elem())
	fmt.Println("m Kind: ", tt.Elem().Kind())
}

func TestReflect7(t *testing.T) {
	m := Profile{}

	v := reflect.ValueOf(&m)

	fmt.Println("&m Type: ", v.Type())
	fmt.Println("&m Kind: ", v.Kind())

	fmt.Println("m Type: ", v.Elem().Type())
	fmt.Println("m Kind: ", v.Elem().Kind())
}

/* 反射对象到普通类型的转换 */
func TestReflect8(t *testing.T) {
	var age int = 25

	v1 := reflect.ValueOf(age)
	fmt.Printf("转换前， type: %T, value: %v \n", v1, v1)
	v2 := v1.Int()
	fmt.Printf("转换后， type: %T, value: %v \n", v2, v2)
}

func TestReflect9(t *testing.T) {
	var age int = 25

	v1 := reflect.ValueOf(&age)
	fmt.Printf("转换前， type: %T, value: %v \n", v1, v1)
	v2 := v1.Pointer()
	fmt.Printf("转换后， type: %T, value: %v \n", v2, v2)
}

func TestReflect10(t *testing.T) {
	var age int = 25

	v1 := reflect.ValueOf(age)
	fmt.Printf("转换前， type: %T, value: %v \n", v1, v1)
	v2 := v1.Interface()
	fmt.Printf("转换后， type: %T, value: %v \n", v2, v2)
}

func TestReflect100(t *testing.T) {
	var numList []int = []int{1, 2, 3, 4, 5}

	v1 := reflect.ValueOf(numList)
	fmt.Printf("转换前， type: %T, value: %v \n", v1, v1)

	// Slice 函数接收两个参数
	v2 := v1.Slice(1, 3)
	fmt.Printf("转换后， type: %T, value: %v \n", v2, v2)

	v3 := v1.Slice3(1, 2, 3)
	fmt.Printf("转换后， type: %T, value: %v \n", v3, v3)
}

func appendToSlice(arrPtr interface{}) {
	valuePtr := reflect.ValueOf(arrPtr)
	value := valuePtr.Elem()

	value.Set(reflect.Append(value, reflect.ValueOf(3)))

	fmt.Println(value)
	fmt.Println(value.Len())
}

func TestReflect11(t *testing.T) {
	arr := []int{1, 2}

	appendToSlice(&arr)

	fmt.Println(arr)
}

type Person struct {
	name   string
	age    int
	gender string
}

func (p Person) SayBye() {
	fmt.Println("Bye")
}

func (p Person) SayHello() {
	fmt.Println("Hello")
}

func TestReflect13(t *testing.T) {
	p := Person{"写代码的明哥", 27, "male"}

	v := reflect.ValueOf(p)

	fmt.Println("字段数:", v.NumField())
	fmt.Println("第 1 个字段：", v.Field(0))
	fmt.Println("第 2 个字段：", v.Field(1))
	fmt.Println("第 3 个字段：", v.Field(2))

	fmt.Println("==========================")
	// 也可以这样来遍历
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("第 %d 个字段：%v \n", i+1, v.Field(i))
	}
}

func TestReflect14(t *testing.T) {
	p := &Person{"写代码的明哥", 27, "male"}

	tt := reflect.TypeOf(p)

	fmt.Println("方法数（可导出的）:", tt.NumMethod())
	fmt.Println("第 1 个方法：", tt.Method(0).Name)
	fmt.Println("第 2 个方法：", tt.Method(1).Name)

	fmt.Println("==========================")
	// 也可以这样来遍历
	for i := 0; i < tt.NumMethod(); i++ {
		fmt.Printf("第 %d 个方法：%v \n", i+1, tt.Method(i).Name)
	}
}

func TestReflect15(t *testing.T) {
	p := &Person{"写代码的明哥", 27, "male"}

	tt := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	for i := 0; i < v.NumMethod(); i++ {
		fmt.Printf("调用第 %d 个方法：%v ，调用结果：%v\n",
			i+1,
			tt.Method(i).Name,
			v.Elem().Method(i).Call(nil))
	}
}

func TestReflect16(t *testing.T) {
	p := &Person{"写代码的明哥", 27, "male"}

	v := reflect.ValueOf(p)

	println("---------------------")

	v.MethodByName("SayHello").Call(nil)
	v.MethodByName("SayBye").Call(nil)
}

func (p Person)SelfIntroduction(name string, age int)  {
	fmt.Printf("Hello, my name is %s and i'm %d years old.", name, age)
}

func TestReflect17(t *testing.T) {
	p := &Person{}

	//t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)
	name := reflect.ValueOf("wangbm")
	age := reflect.ValueOf(27)
	input := []reflect.Value{name, age}
	v.MethodByName("SelfIntroduction").Call(input)
}
