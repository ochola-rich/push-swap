package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"push-swap/internal/stack"
)

func main() {
	inst := []string{"pb", "pb", "rb", "pb", "sa", "pa", "pa", "pa"}

	if len(os.Args) < 2 {
		return
	}

	// Parsing Input
	var stackA []int
	input := os.Args[1:]
	if len(input) == 1 {
		input = strings.Fields(input[0])
	}
	for _, arg := range input {
		n, _ := strconv.Atoi(arg)
		stackA = append(stackA, n)
	}
	stackB := []int{}
	for _, ins := range inst {
		switch ins {
		case "pb":
			stack.Pb(&stackA, &stackB)
		case "ra":
			stack.Ra(stackA, "ra")
		case "sa":
			stackA[0], stackA[1] = stackA[1], stackA[0]
		case "pa":
			stack.Pa(&stackA, &stackB)
		case "rb":
			stack.Ra(stackB, "rb")

		}
	}
	fmt.Println(stackA, stackB)
}
