package constant_test

import (
	"testing"
)

type MyInt int	//给int起一个别名

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstantTry1(t *testing.T) {
	a := 1 //0001
	t.Log(Readable, Writable, Executable)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a

	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)	//打印变量的类型
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
}

func TestCompareArray(t *testing.T) {
	//a := [...] int {1, 2, 3, 4}
	//b := [...] int {1, 2, 3, 4}
}