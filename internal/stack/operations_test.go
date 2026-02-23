package stack

import (
	"reflect"
	"testing"
)

func TestPa(t *testing.T) {
	a := []int{2}
	b := []int{1, 3}
	res := Pa(&a, &b)

	if res != "pa" {
		t.Errorf("expected 'pa', got %q", res)
	}
	if !reflect.DeepEqual(a, []int{1, 2}) {
		t.Errorf("expected stack A [1, 2], got %v", a)
	}
}

func TestPb_EmptyStack(t *testing.T) {
	a := []int{}
	b := []int{1}
	res := Pb(&a, &b)

	if res != "" {
		t.Errorf("expected empty string for empty stack operation, got %q", res)
	}
}

func TestR(t *testing.T) {
	s := []int{1, 2, 3}
	R(s, "ra")
	expected := []int{2, 3, 1}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("expected %v, got %v", expected, s)
	}
}

func TestRr(t *testing.T) {
	s := []int{1, 2, 3}
	Rr(s, "rra")
	expected := []int{3, 1, 2}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("expected %v, got %v", expected, s)
	}
}
