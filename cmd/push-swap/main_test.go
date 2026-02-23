package main

import (
	"testing"
)

func TestGetDist_Main(t *testing.T) {
	dist, isTop := GetDist(0, 3)
	if dist != 0 || !isTop {
		t.Errorf("expected (0, true), got (%d, %t)", dist, isTop)
	}
}

func TestMin(t *testing.T) {
	if Min(10, 20) != 10 {
		t.Error("Min(10, 20) failed")
	}
	if Min(5, -1) != -1 {
		t.Error("Min(5, -1) failed")
	}
}
