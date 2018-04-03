package datautils

import "testing"

func TestMax(t *testing.T) {
	data := []struct {
		x int64
		y int64
		m int64
	}{
		{1, 1, 1},
		{1, 2, 2},
		{2, 1, 2},
	}

	for _, d := range data {
		max := Max(d.x, d.y)
		if max != d.m {
			t.Errorf("Max(%d, %d) was wrong: %d is not %d", d.x, d.y, max, d.m)
		}
	}
}
