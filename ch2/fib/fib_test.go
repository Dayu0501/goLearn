package fib

/* package名字和该文件所在的文件的名字，没有关系的 */

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T)  {
	var a int = 1
	var b int = 1

	/* 或者下面写法（因为go语言存在着类型推断）
	var (
	a int = 1
	b int = 1
	)

	在或者
	var (
	a = 1
	b = 1
	)

	在或者直接赋值并初始化
	a := 1   //声明并赋值
	b := 1

	在或者
	var a int

	.
	.
	.

	a = 1    //赋值，该操作的前提是，已经使用上面的声明过了
	*/

	t.Log(a)
	for i:=0; i < 5; i++ {
		t.Log(b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	//tmp := a
	//a = b
	//b = tmp

	a, b = b, a //此处的交换代码等价于上面的三杯水的交换方式

}
