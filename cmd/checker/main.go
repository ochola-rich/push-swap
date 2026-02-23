package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"push-swap/internal/stack"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	// Parsing Input
	var stackA []int
	input := os.Args[1:]

	if len(input) == 1 {
		input = strings.Fields(input[0])
	}
	if hasDuplicates(stackA) { // Ensure no duplicates
		fmt.Println("Error")
		return
	}
	if slices.IsSorted(input) {
		return
	}
	for _, arg := range input {
		n, err := strconv.Atoi(arg)
		// Ensure only numbers
		if err != nil {
			fmt.Println("Error")
			return
		}
		stackA = append(stackA, n)
	}

	stackB := []int{}

	// Scan operations until EOF
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		switch line {
		case "pa":
			stack.Pa(&stackA, &stackB)
		case "pb":
			stack.Pb(&stackA, &stackB)
		case "ra":
			stack.R(stackA, "ra")
		case "rb":
			stack.R(stackB, "rb")
		case "sa":
			stackA[0], stackA[1] = stackA[1], stackA[0]
		case "sb":
			stackB[0], stackB[1] = stackB[1], stackB[0]
		case "rr":
			stack.RotateSilent(stackA)
			stack.RotateSilent(stackB)
		case "rra":
			stack.RevRotateSilent(stackA)
		case "rrb":
			stack.RevRotateSilent(stackB)
		case "rrr":
			stack.RevRotateSilent(stackA)
			stack.RevRotateSilent(stackB)
		}
	}
	// Validate only after EOF(Ctrl+D)
	if slices.IsSorted(stackA) && len(input) == len(stackA) {
		fmt.Println("OK")
		return
	}

	fmt.Println("KO")
}


func hasDuplicates(elements []int) bool {
    encountered := map[int]bool{}
    for _, v := range elements {
        if encountered[v] {
            return true // Duplicate found!
        }
        encountered[v] = true
    }
    return false
}
