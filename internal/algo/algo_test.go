package algo

import (
	"testing"
)

func TestFindTargetIdxB(t *testing.T) {
	stackB := []int{10, 5, 2}
	valA := 7
	// Should find the index of 5 (the closest smaller value)
	got := FindTargetIdxB(valA, stackB)
	if got != 1 {
		t.Errorf("expected index 1, got %d", got)
	}
}

func TestFindTargetIdxA_Smallest(t *testing.T) {
	stackA := []int{5, 10, 20}
	valB := 25
	// If valB is larger than everything, it should target the smallest value (index 0)
	got := FindTargetIdxA(valB, stackA)
	if got != 0 {
		t.Errorf("expected index 0, got %d", got)
	}
}

func TestGetDist(t *testing.T) {
	tests := []struct {
		idx    int
		length int
		wantD  int
		wantT  bool
	}{
		{1, 5, 1, true},
		{4, 5, 1, false},
		{2, 6, 2, true},
	}

	for _, tt := range tests {
		dist, isTop := GetDist(tt.idx, tt.length)
		if dist != tt.wantD || isTop != tt.wantT {
			t.Errorf("GetDist(%d, %d) = (%d, %t); want (%d, %t)",
				tt.idx, tt.length, dist, isTop, tt.wantD, tt.wantT)
		}
	}
}
