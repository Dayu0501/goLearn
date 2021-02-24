/*
	应用程序入口
    1, 必须是main包：package main
	2, 必须是main方法：func main()
	3, 文件名不一定是main.go

	Go中main函数不支持函数的返回值;
    main可以通过os.Exit来返回状态;

	测试程序
	1, 编写文件以_test结尾：xxx_test.go
	2, 测试方法必须以Test开头：func Testxxx(t *testing.T) {...}
	3, 大写字母开头的方法，代表着包外可以访问；
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args) //输出命令行参数
	if len(os.Args) > 1 {
		fmt.Println("hello world", os.Args[1])
	}




	os.Exit(-1)
}
