// Package algo
package algo

import (
	"fmt"
	"math"

	"push-swap/internal/models"
	"push-swap/internal/stack"
)

func GetBestPlan(stackA, stackB []int) models.MovePlan {
	var bestPlan models.MovePlan
	bestPlan.TotalCost = math.MaxInt
	lenA, lenB := len(stackA), len(stackB)

	for i, valA := range stackA {
		targetIdxB := FindTargetIdxB(valA, stackB)
		distA, isATop := GetDist(i, lenA)
		distB, isBTop := GetDist(targetIdxB, lenB)

		plan := calculateStrategy(distA, distB, isATop, isBTop)
		if plan.TotalCost < bestPlan.TotalCost {
			bestPlan = plan
		}
	}
	return bestPlan
}

func calculateStrategy(distA, distB int, isATop, isBTop bool) models.MovePlan {
	p := models.MovePlan{}
	if isATop && isBTop {
		p.Shared = Min(distA, distB)
		p.SharedType = "rr"
		p.RemainingA, p.RemainingB = distA-p.Shared, distB-p.Shared
		p.MoveA, p.MoveB = "ra", "rb"
	} else if !isATop && !isBTop {
		p.Shared = Min(distA, distB)
		p.SharedType = "rrr"
		p.RemainingA, p.RemainingB = distA-p.Shared, distB-p.Shared
		p.MoveA, p.MoveB = "rra", "rrb"
	} else {
		p.RemainingA, p.RemainingB = distA, distB
		if isATop {
			p.MoveA = "ra"
		} else {
			p.MoveA = "rra"
		}
		if isBTop {
			p.MoveB = "rb"
		} else {
			p.MoveB = "rrb"
		}
	}
	p.TotalCost = p.Shared + p.RemainingA + p.RemainingB
	return p
}

func FindTargetIdxB(valA int, stackB []int) int {
	targetIdx := -1
	closestSmaller := math.MinInt64
	for i, valB := range stackB {
		if valB < valA && valB > closestSmaller {
			closestSmaller = valB
			targetIdx = i
		}
	}
	if targetIdx == -1 {
		maxVal := math.MinInt64
		for i, valB := range stackB {
			if int64(valB) > int64(maxVal) {
				maxVal = valB
				targetIdx = i
			}
		}
	}
	return targetIdx
}

func FindTargetIdxA(valB int, stackA []int) int {
	targetIdx := -1
	closestBigger := math.MaxInt64
	for i, valA := range stackA {
		if valA > valB && valA < closestBigger {
			closestBigger = valA
			targetIdx = i
		}
	}
	if targetIdx == -1 {
		minVal := math.MaxInt64
		for i, valA := range stackA {
			if valA < minVal {
				minVal = valA
				targetIdx = i
			}
		}
	}
	return targetIdx
}

func ExecuteMove(p models.MovePlan, a, b *[]int) {
	for i := 0; i < p.Shared; i++ {
		if p.SharedType == "rr" {
			stack.RotateSilent(*a)
			stack.RotateSilent(*b)
			fmt.Println("rr")
		} else {
			stack.RevRotateSilent(*a)
			stack.RevRotateSilent(*b)
			fmt.Println("rrr")
		}
	}
	for i := 0; i < p.RemainingA; i++ {
		if p.MoveA == "ra" {
			stack.Ra(*a, "ra")
		} else {
			stack.Rra(*a, "rra")
		}
	}
	for i := 0; i < p.RemainingB; i++ {
		if p.MoveB == "rb" {
			stack.Ra(*b, "rb")
		} else {
			stack.Rra(*b, "rrb")
		}
	}
	stack.Pb(a, b)
}

func SortThree(a []int) {
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	if a[0] == max {
		stack.Ra(a, "ra")
	} else if a[1] == max {
		stack.Rra(a, "rra")
	}
	if a[0] > a[1] {
		a[0], a[1] = a[1], a[0]
		fmt.Println("sa")
	}
}

func FinalizeA(a []int) {
	minIdx := 0
	minVal := a[0]
	for i, val := range a {
		if val < minVal {
			minVal = val
			minIdx = i
		}
	}
	dist, isTop := GetDist(minIdx, len(a))
	for range dist {
		if isTop {
			stack.Ra(a, "ra")
		} else {
			stack.Rra(a, "rra")
		}
	}
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
