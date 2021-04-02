package arrayT_test

import (
	"fmt"
	"testing"
	"time"
)

type arr3 [3]int

func TestName(t *testing.T) {
	var haha arr3 = arr3{1, 2, 3}
	t.Log(haha)

	arr := [4]int{2: 3}
	t.Log(arr)
}

func TestArray(t *testing.T) {
	a := [10]int{1, 3, 4, 5, 6, 7}

	b := [5]string{"laoWang", "xiaoChang", "dingDing", "laoSan"}
	t.Log(b)

	c := [...]string{"wen", "shuang", "nan"}
	t.Log(c)

	/* 切片用在数组的上的时候，数组和该切片的变动，都会相互影响 */
	c[1] = "shuang ge"

	d := []int{1, 2, 3, 4}
	t.Log(d)

	e := c[1:2]
	t.Log("-------------")
	t.Log(e)
	t.Log("++++++++++++++")
	e[0] = "xiao shuang"

	t.Log(e)
	t.Log(c[1])

	for i := 0; i < len(a); i++ {
		t.Log("a is ", a[i])
	}

	for index, e := range a {
		t.Log(index, e)
	}
}

func TestArray1(t *testing.T) {
	a := [...]int{4, 5, 6, 7, 8, 9} //不设定长度的数组
	t.Log(len(a))
}

func TestSlice(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)

	s2 := make([]int, 3, 5) //在go中[]int就是切片的类型，而数组的类型是带有长度的
	t.Log(len(s2), cap(s2))

	s3 := []int{10, 20, 30, 40}
	t.Log(len(s3), cap(s3))

	myarr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("myarr 的长度为：%d，容量为：%d\n", len(myarr), cap(myarr))

	/* 切片的第三个参数，会影响到切片的容量 */
	mysli2 := myarr[1:3:4]
	fmt.Printf("mysli2 的长度为：%d，容量为：%d\n", len(mysli2), cap(mysli2))
	fmt.Println(mysli2)

	// 声明字符串切片
	var strList []string
	t.Log(strList)

	// 声明整型切片
	var numList []int
	t.Log(numList)

	// 声明一个空切片
	var numListEmpty = []int{}
	t.Log(numListEmpty)

	t.Log(numList == nil)
}

func TestSlice2(t *testing.T) {
	myarr := []int{1}
	t.Log(myarr)

	// 追加一个元素
	myarr = append(myarr, 2)
	t.Log(myarr)

	// 追加多个元素
	myarr = append(myarr, 3, 4)
	t.Log(myarr)

	// 追加一个切片, ... 表示解包，不能省略
	myarr = append(myarr, []int{7, 8}...)
	t.Log(myarr)

	// 在第一个位置插入元素
	myarr = append([]int{2}, myarr...)
	t.Log(myarr)

	// 在中间插入一个切片(两个元素)
	myarr = append(myarr[:5], append([]int{5, 6}, myarr[5:]...)...)
	fmt.Println(myarr)
}

func TestSlice3(t *testing.T) {
	arr1 := [4]int{1, 2, 3, 100}
	t.Log(arr1)

	slice1 := arr1[2:3]
	t.Log(slice1)

	/* 当slice是引用一个之前声明过了的数组，slice的append操作，也会修改之前数组的值 */
	slice1 = append(slice1, 4)
	t.Log(slice1)
	t.Log(arr1)

	slice1 = append(slice1, 5, 6, 7, 8)
	t.Log(slice1)
	t.Log(arr1)
}

func TestSliceTest3(t *testing.T) {
	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myslice := numbers4[4:6:8]
	fmt.Printf("myslice为 %d, 其长度为: %d\n", myslice, len(myslice))

	myslice = myslice[:cap(myslice)]
	fmt.Println(myslice)
	fmt.Printf("myslice的第四个元素为: %d", myslice[3])
}

func TestSliceGrowing(t *testing.T) {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSlice5(t *testing.T) {
	// 声明字符串切片
	var strList []string
	strList = append(strList, "laowang")

	// 声明整型切片
	//var numList []int
	//下面这段代码是错误的
	//numList[0] = 1

	// 声明一个空切片
	var numListEmpty = []int{}
	numListEmpty = append(numListEmpty, 0)
}

/* go的map是hash表的一种实现*/
func TestMap1(t *testing.T) {
	var a map[string]int //map[string]int 是map的类型
	t.Log(len(a))

	b := map[string]int{"hello": 1, "world": 2}

	t.Log(b)
}

func TestMap2(t *testing.T) {
	a := map[int]int{0: 1}

	if v, ok := a[0]; ok { //map中某个key是否存在需要主动去判断，v代表着key对应的值，ok为true代表存在，为false代表不存在
		t.Log(v)
	} else {
		t.Log("unkown")
	}
}

func TestMap3(t *testing.T) {
	var scores map[string]int

	/* 下面这个是错误的代码，因为map没有初始化 */
	/*scores["laowang"] = 100
	t.Log(scores)*/

	if scores == nil {
		scores = make(map[string]int)
	}
	t.Log(scores)

	scores["changge"] = 200
	scores["laowang"] = 99
	t.Log(scores)

	/* 如果不存在的key会返回其零值 */
	fmt.Println(scores["math"])

	/* 使用 delete 函数，如果 key 不存在，delete 函数会静默处理，不会报错。 */
	delete(scores, "math")
	delete(scores, "laowang")
	t.Log(scores)

	if math, ok := scores["laowang"]; ok {
		t.Log(math)
	} else {
		t.Log("不存在的！")
	}

}

/* map(字典)遍历 */
func TestMap4(t *testing.T) {
	scores := map[string]int{"english": 80, "chinese": 85}

	for subject, score := range scores {
		fmt.Printf("key: %s, value: %d\n", subject, score)
	}

	for subject := range scores {
		fmt.Printf("key: %s\n", subject)
	}

	for _, score := range scores {
		fmt.Printf("value: %d\n", score)
	}
}

func TestPtr(t *testing.T) {
	aint := 1    // 定义普通变量
	ptr := &aint // 定义指针变量
	fmt.Println("普通变量存储的是：", aint)
	fmt.Println("普通变量存储的是：", *ptr)
	fmt.Println("指针变量存储的是：", &aint)
	fmt.Println("指针变量存储的是：", ptr)

	/* 打印指针指向的内存 */
	// 第一种
	fmt.Printf("%p", ptr)

	// 第二种
	fmt.Println(ptr)

	astr := "hello"
	aint = 1
	abool := false
	arune := 'a'
	afloat := 1.2

	fmt.Printf("astr 指针类型是：%T\n", &astr)
	fmt.Printf("aint 指针类型是：%T\n", &aint)
	fmt.Printf("abool 指针类型是：%T\n", &abool)
	fmt.Printf("arune 指针类型是：%T\n", &arune)
	fmt.Printf("afloat 指针类型是：%T\n", &afloat)
}

func TestArrModify(t *testing.T) {
	arr := [10]int{10, 20, 30, 40, 50}
	t.Log(arr)
	modifyArr(arr[:])
	t.Log(arr)
}

func modifyArr(sli []int) bool {
	sli[1] = 10
	return true
}

func TestIfElse(t *testing.T) {
	a := 10

	if a > 5 && a < 20 {
		t.Log("ok")
	} else {
		t.Log("not ok")
	}

	if b := 20; b > a {
		t.Log("b > a")
	} else {
		t.Log("b <= a")
	}
}

func TestSwitch(t *testing.T) {
	a := 80
	switch {
	case a >= 80 && a < 100:
		t.Log("优秀")
	case a >= 60 && a < 80:
		t.Log("一般")
	default:
		t.Log("不知道")
	}

	s := "hello"
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s != "world":
		fmt.Println("world")
		fallthrough
	case s == "haha":
		fmt.Println("haha")
	}
}

func makeTimeOut(ch chan bool, t int)  {
	time.Sleep(time.Second * time.Duration(t))
	ch <- true;
}

func TestSelect(t *testing.T) {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	timeOut := make(chan bool, 1)


	go makeTimeOut(timeOut, 2)
	//c1 <- "laowang"

	//go func() {
	//	c2 <- "hello world"
	//}()

	select {
	case msg := <- c1:
		t.Log(msg)
	case msg := <- c2:
		t.Log(msg)
	case <- timeOut:
		t.Log("time out")

	//default:
	//	t.Log("no msg")
	}
}

func TestSelect2(t *testing.T) {
	c1 := make(chan int, 2)

	c1 <- 2
	select {
	case c1 <- 4:
		fmt.Println("c1 received: ", <-c1)
		fmt.Println("c1 received: ", <-c1)
	default:
		fmt.Println("channel blocking")
	}
}

func TestMainPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log("recover test test main panic !")
		}

	}()

	testPanic()
}

func testPanic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("testPanic end !")
		}
	}()

	panic("testPanic !")
}