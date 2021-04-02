package ch9

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestFmt(t *testing.T) {
	fmt.Print("hello", "world\n")
	fmt.Println("hello", "world", "china")
	fmt.Printf("hello world\n")
}

func TestFmt1(t *testing.T) {
	n := 1024
	fmt.Printf("%d 的 2 进制：%b \n", n, n)
	fmt.Printf("%d 的 8 进制：%o \n", n, n)
	fmt.Printf("%d 的 10 进制：%d \n", n, n)
	fmt.Printf("%d 的 16 进制：%x \n", n, n)
}

type Profile struct {
	name   string
	gender string
	age    int
}

func TestFmt2(t *testing.T) {
	var people = Profile{name: "wangbm", gender: "male", age: 27}
	fmt.Printf("%v \n", people) // output: {wangbm male 27}
	fmt.Printf("%T \n", people) // output: main.Profile

	// 打印结构体名和类型
	fmt.Printf("%#v \n", people) // output: main.Profile{name:"wangbm", gender:"male", age:27}
	fmt.Printf("%+v \n", people) // output: {name:wangbm gender:male age:27}
	fmt.Printf("%% \n")          // output: %
}

/* 打印布尔值 */
func TestFmt3(t *testing.T) {
	fmt.Printf("%t \n", true)  //output: true
	fmt.Printf("%t \n", false) //output: false
}

func TestFmt4(t *testing.T) {
	fmt.Printf("%s \n", []byte("Hello, Golang")) // output: Hello, Golang
	fmt.Printf("%s \n", "Hello, Golang")         // output: Hello, Golang

	fmt.Printf("%q \n", []byte("Hello, Golang")) // output: "Hello, Golang"
	fmt.Printf("%q \n", "Hello, Golang")         // output: "Hello, Golang"
	fmt.Printf("%q \n", `hello \r\n world`)      // output: "hello \\r\\n world"

	fmt.Printf("%x \n", "Hello, Golang") // output: 48656c6c6f2c20476f6c616e67
	fmt.Printf("%X \n", "Hello, Golang") // output: 48656c6c6f2c20476f6c616e67
}

/* 打印指针 */
func TestFmt5(t *testing.T) {
	var people = Profile{name: "wangbm", gender: "male", age: 27}
	fmt.Printf("%p", &people) // output: 0xc0000a6150
}

func TestFmt6(t *testing.T) {
	n := 1024
	fmt.Printf("%d 的 2 进制：%b \n", n, n)
	fmt.Printf("%d 的 8 进制：%o \n", n, n)
	fmt.Printf("%d 的 10 进制：%d \n", n, n)
	fmt.Printf("%d 的 16 进制：%x \n", n, n)

	// 将 10 进制的整型转成 16 进制打印： %x 为小写， %X 为小写
	fmt.Printf("%x \n", 1024)
	fmt.Printf("%X \n", 1024)

	// 根据 Unicode码值打印字符
	fmt.Printf("ASCII 编码为%d 表示的字符是： %c \n", 65, 65) // output: A

	// 根据 Unicode 编码打印字符
	fmt.Printf("%c \n", 0x4E2D) // output: 中
	// 打印 raw 字符时
	fmt.Printf("%q \n", 0x4E2D) // output: '中'

	// 打印 Unicode 编码
	fmt.Printf("%U \n", '中') // output: U+4E2D
}

func TestFunc7(t *testing.T) {
	f := 12.34
	fmt.Printf("%b\n", f)
	fmt.Printf("%e\n", f)
	fmt.Printf("%E\n", f)
	fmt.Printf("%f\n", f)
	fmt.Printf("%g\n", f)
	fmt.Printf("%G\n", f)
}

func TestFmt8(t *testing.T) {
	n := 12.34
	fmt.Printf("%f\n", n)    // 以默认精度打印
	fmt.Printf("%9f\n", n)   // 宽度为9，默认精度
	fmt.Printf("%.2f\n", n)  // 默认宽度，精度2
	fmt.Printf("%9.2f\n", n) //宽度9，精度2
	fmt.Printf("%9.f\n", n)  // 宽度9，精度0
}

func TestFmt9(t *testing.T) {
	cmd := exec.Command("ls", "-l", "/var/log/")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func TestFmt10(t *testing.T) {
	cmd := exec.Command("ls", "-l", "/var/log/")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}

func TestFmt11(t *testing.T) {
	cmd := exec.Command("ls", "-l", "/var/log/*.log")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()

	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func TestFmt12(t *testing.T) {
	c1 := exec.Command("grep", "ERROR", "/var/log/messages")
	c2 := exec.Command("wc", "-l")
	c2.Stdin, _ = c1.StdoutPipe()
	c2.Stdout = os.Stdout
	_ = c2.Start()
	_ = c1.Run()
	_ = c2.Wait()
}

func ChangeYourCmdEnvironment(cmd *exec.Cmd) error {
	env := os.Environ()
	cmdEnv := []string{}

	for _, e := range env {
		cmdEnv = append(cmdEnv, e)
	}
	cmdEnv = append(cmdEnv, "NAME=wangbm")
	cmd.Env = cmdEnv

	return nil
}

func TestFmt13(t *testing.T) {
	cmd1 := exec.Command("bash", "/home/wangbm/demo.sh")
	ChangeYourCmdEnvironment(cmd1) // 添加环境变量到 cmd1 命令: NAME=wangbm
	out1, _ := cmd1.CombinedOutput()
	fmt.Printf("output: %s", out1)

	cmd2 := exec.Command("bash", "/home/wangbm/demo.sh")
	out2, _ := cmd2.CombinedOutput()
	fmt.Printf("output: %s", out2)
}

func TestFmt14(t *testing.T) {
	var name string
	flag.StringVar(&name, "name", "jack", "your name")

	flag.Parse()  // 解析参数
	fmt.Println(name)
}
