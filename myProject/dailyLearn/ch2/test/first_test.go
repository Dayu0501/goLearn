package try_test

import (
	"fmt"
	"testing"
)

func TestFirst(t *testing.T) {
	t.Log("My first try !")
	
}

func TestBitOp(t *testing.T) {
	x := 11
	y := (1 << 0) | (1 << 3) //保证 z 中的第 0 位和第 3 位为 0
	z := x &^ y

	fmt.Printf("x = %b\n", x)
	fmt.Println("\t\t&^")
	fmt.Printf("y = %b\n", y)
	fmt.Println("————————")
	fmt.Printf("z = %04b\n", z)
}
