package main

import "testing"

func TestSum(t *testing.T) {
	tables := []struct{
		x int
		y int
		n int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
	}

	for _, table := range tables {
		total := Sum(table.x, table.y)
		t.Logf("Sum(%d, %d): %d", table.x, table.y, total)
		if total != table.n {
			t.Errorf("Sum (%d, %d) está incorreta: obtido %d, esperado %d", table.x, table.y, total, table.n)
		}
	}
	
}
