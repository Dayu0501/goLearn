package main

import (
	"fmt"
	"sync"
	"time"
)

type info struct {
	rmu sync.RWMutex
}

/* for range + 闭包问题的解决方法： 给闭包中传入参数*/
/* 输出结果不一定是 i am Ethan world,因为goroutine执行顺序有go runtime调度器决定 */
func test1() {
	str := []string{"I", "am", "Ethan", "world"}
	for _, v := range str {
		go func(v string) {
			fmt.Println(v)
		}(v)
	}
	time.Sleep(3 * time.Second)
}

/* 其实真实原因也是如此，因为for range创建了每个元素的副本，而不是直接返回每个元素的引用，如果使用该值变量的地址作为指向每个元素的指针，
   就会导致错误，在迭代时，返回的变量是一个迭代过程中根据切片依次赋值的新变量，所以值的地址总是相同的，导致结果不如预期。
*/
func test2() {
	slice := []int{0, 1, 2, 3}
	myMap := make(map[int]*int)
	for index, value := range slice {
		myMap[index] = &value
	}

	fmt.Println(myMap)

	fmt.Println("********test2***********")

	prtMap(myMap)
}

func test3() {
	slice := []int{0, 1, 2, 3}
	myMap := make(map[int]*int)

	for index, value := range slice {
		v := value
		myMap[index] = &v
	}

	fmt.Println(myMap)
	fmt.Println("********test3***********")
	prtMap(myMap)
}

func prtMap(myMap map[int]*int) {
	for key, value := range myMap {
		fmt.Printf("map[%v]=%v\n", key, *value)
	}
}

type myFloat float64

type mytest interface {
	abs1() int
	abs2() int
}

func (f myFloat) abs1() int {
	fmt.Println("abs1")
	return 1
}

func (f myFloat) abs2() int {
	fmt.Println("abs2")
	return 2
}


func main() {
	/* 下面的for range + 闭包，导致的是，在go routine运行的时候，v这个值指向str中的哪一个值是不一定的 */

	str := []string{"I", "am", "Ethan", "hello"}
	for _, v := range str {

		/* 下面的func(){}是一个匿名函数 */
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(3 * time.Second)

	test1()

	fmt.Println("***************************")

	test2()

	fmt.Println("***************************")

	test3()

	fmt.Println("****************************")

	var cc mytest
	cc = myFloat(1)

	println(cc.abs1())

}

//func main() {
//	defer func() {
//		if err := recover(); err != nil {
//			log.Printf("recover: %v", err)
//		}
//	}()
//
//	/* 以并发的方式异步调用匿名函数 */
//	go func() {
//		log.Printf("test annoyname func")
//	}()
//
//	panic("panic learn!")
//}
