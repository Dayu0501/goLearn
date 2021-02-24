package arrayT_test

import "testing"

func TestArray(t *testing.T)  {
	a := [10]int{1, 3, 4, 5, 6, 7}

	for i := 0; i < len(a); i++ {
		t.Log("a is " , a[i])
	}

	for index, e := range a {
		t.Log(index, e)
	}
}

func TestArray1(t *testing.T) {
	a := [...] int {4, 5, 6, 7, 8, 9}   //不设定长度的数组
	t.Log(len(a))
}

func TestSlice (t *testing.T) {
	var s0 [] int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)

	s2 := make([]int, 3, 5)			 //在go中[]int就是切片的类型，而数组的类型是带有长度的
	t.Log(len(s2), cap(s2))
}

func TestSliceGrowing(t *testing.T) {
	var s[]int
	for i := 0; i < 10; i++ {
		 s = append(s, i)
		 t.Log(len(s), cap(s))
	}
}

func TestMap1(t *testing.T) {
	var a map[string]int			//map[string]int 是map的类型
	t.Log(len(a))

	b := map [string] int {"hello" : 1, "world" : 2}

	t.Log(b)
}

func TestMap2(t *testing.T) {
	a := map [int] int {0:1}

	if v, ok := a[0]; ok { 				//map中某个key是否存在需要主动去判断，v代表着key对应的值，ok为true代表存在，为false代表不存在
		t.Log(v)
	} else {
		t.Log("unkown")
	}
}
