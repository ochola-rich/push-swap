// Package stack
package stack

func Pa(a *[]int, b *[]int) string {
	if len(*b) == 0 {
		return ""
	}
	val := (*b)[0]
	*b = (*b)[1:]
	*a = append([]int{val}, (*a)...)

	return "pa"
}

func Pb(a *[]int, b *[]int) string {
	if len(*a) == 0 {
		return ""
	}
	val := (*a)[0]
	*a = (*a)[1:]
	*b = append([]int{val}, (*b)...)

	return "pb"
}

func R(s []int, label string) string {
	if len(s) < 2 {
		return ""
	}
	first := s[0]
	copy(s, s[1:])
	s[len(s)-1] = first

	return label
}

func Rr(s []int, label string) string {
	if len(s) < 2 {
		return ""
	}
	last := s[len(s)-1]
	copy(s[1:], s)
	s[0] = last

	return label
}

func RotateSilent(s []int) {
	if len(s) < 2 {
		return
	}
	f := s[0]
	copy(s, s[1:])
	s[len(s)-1] = f
}

func RevRotateSilent(s []int) {
	if len(s) < 2 {
		return
	}
	l := s[len(s)-1]
	copy(s[1:], s)
	s[0] = l
}
