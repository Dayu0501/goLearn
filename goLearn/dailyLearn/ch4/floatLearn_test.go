package ch4

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestFloat32(t *testing.T) {
	// 8388608 10000018
	var float1 float32 = 10000018
	t.Log(float1)
	t.Log(float1 + 1)

	t.Log("--------------------------")
	var float2 float32 = 100000182
	t.Log(float2)
	t.Log(float2 + 1)

	var a float32 = 2567
	t.Log(a)
	var b float32 = 3456
	t.Log(b)
	var c float32 = a + b

	t.Log(a + b == c)
}

func TestByte(t *testing.T) {
	var a byte = 65
	// 8进制写法: var a byte = '\101'     其中 \ 是固定前缀
	// 16进制写法: var a byte = '\x41'    其中 \x 是固定前缀

	var b uint8 = 66
	fmt.Printf("a 的值: %c \nb 的值: %c\n", a, b)

	// 或者使用 string 函数
	// fmt.Println("a 的值: ", string(a)," \nb 的值: ", string(b))
}

func TestRune(t *testing.T) {
	var a byte = 'A' //表示的是ASCII的字符（范围小，不能表示中文字符）
	var b rune = 'B'
	fmt.Printf("a 占用 %d 个字节数\nb 占用 %d 个字节数\n", unsafe.Sizeof(a), unsafe.Sizeof(b))

	var c rune = '中' //表示的是unicode编码的字符，范围到，可以表示中文字符


	/* 因为uint8 和 uint32 ，直观上让人以为这是一个数值，但是实际上，它也可以表示一个字符 */
	var d int32 = 20013
	var e int32 = '中'
	fmt.Printf("%c\n", c)
	fmt.Printf("%c\n", d)
	fmt.Printf("%d\n", d)
	fmt.Printf("%c\n", e)
}

/* 说明了 string 的本质，其实是一个 byte数组 */
func TestString(t *testing.T) {
	var mystr01 string = "hello"
	var mystr02 [5]byte = [5]byte{104, 101, 108, 108, 111}
	fmt.Printf("mystr01: %s\n", mystr01)
	fmt.Printf("mystr02: %s\n", mystr02)

	/*  Go 语言的 string 是用 uft-8 进行编码的，英文字母占用一个字节，而中文字母占用 3个字节，所以 hello,中国 的长度为 5+1+（3＊2)= 12个字节。 */
	var country string = "hello,中国"
	fmt.Println(len(country))

	/* 解释型表示法 */
	var mystr03 string = "heihei \r\n hehe"
	fmt.Println(mystr03)

	/* 原生型表示法 对转移字符相当于"\\r\\n"，就是现实转义字符的本体，所以这类字符串没办法写换行符的字符*/
	var mystr04 string = `heihei \r\n hehe`
	fmt.Println(mystr04)

	var mystr05 string = `\r\n`
	fmt.Print(`\r\n`)
	fmt.Printf("的解释型字符串是： %q", mystr05)

}

func TestArray(t *testing.T) {
	arr01 := [...]int{1, 2, 3}
	arr02 := [...]int{1, 2, 3, 4}
	fmt.Printf("%d 的类型是: %T\n", arr01, arr01)
	fmt.Printf("%d 的类型是: %T\n", arr02, arr02)
}
