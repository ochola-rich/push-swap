package main

import (
	"os"
	"slices"
	"strconv"
	"strings"

	"push-swap/internal/algo"
	"push-swap/internal/stack"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	// Parsing Input
	var stackA []int
	input := os.Args[1:]

	if slices.IsSorted(input) {
		return
	}

	if len(input) == 1 {
		input = strings.Fields(input[0])
	}
	for _, arg := range input {
		n, _ := strconv.Atoi(arg)
		stackA = append(stackA, n)
	}
	stackB := []int{}

	// Phase 1: Push A to B until 3 left
	if len(stackA) > 3 {
		stack.Pb(&stackA, &stackB)
	}
	if len(stackA) > 3 {
		stack.Pb(&stackA, &stackB)
	}
	for len(stackA) > 3 {
		plan := algo.GetBestPlan(stackA, stackB)
		algo.ExecuteMove(plan, &stackA, &stackB)
	}

	// Phase 2: Sort Three
	algo.SortThree(stackA)

	// Phase 3: Push B back to A
	for len(stackB) > 0 {
		targetIdxA := algo.FindTargetIdxA(stackB[0], stackA)
		dist, isTop := GetDist(targetIdxA, len(stackA))
		for range dist {
			if isTop {
				stack.Ra(stackA, "ra")
			} else {
				stack.Rra(stackA, "rra")
			}
		}
		stack.Pa(&stackA, &stackB)
	}

	// Phase 4: Align A
	algo.FinalizeA(stackA)
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func GetDist(idx, length int) (int, bool) {
	if idx <= length/2 {
		return idx, true
	}
	return length - idx, false
}
