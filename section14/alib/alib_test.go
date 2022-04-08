package alib

import (
	"testing"
)

var Debug bool = false

func TestAverage1(t *testing.T) {
	if Debug {
		t.Skip("スキップさせるよ")
	}
	v := Average([]int{2, 4, 6})
	if v != 4 {
		t.Errorf("%v != %v", v, 4)
	}
}
