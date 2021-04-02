package ch9

import "testing"

func TestAdd(t *testing.T) {
	t.Log(Add(1, 2))
}

type TestTable struct {
	xarg int
	yarg int
}

func TestAdd1(t *testing.T) {
	tables := []TestTable{
		{1, 2},
		{2, 4},
		{4, 8},
		{5, 10},
		{6, 12},
	}

	for _, table := range tables {
		result := Add(table.xarg, table.yarg)
		if result == (table.xarg + table.yarg) {
			t.Log("the result is ok")
		} else {
			t.Fatal("the result is wrong")
		}
	}
}