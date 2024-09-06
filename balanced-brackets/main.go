package main

import (
	"flag"
	"log"
)

var pairs = map[rune]rune{
	'{': '}',
	'[': ']',
	'(': ')',
}

// isBalanced returns whether the given expression
// has balanced brackets.
func isBalanced(expr string) bool {
	stack := make([]rune, 0)

	for _, br := range expr {
		if len(stack) == 0 {
			stack = append(stack, br)
      continue
		}

		last := stack[len(stack)-1]
		if pairs[last] == br {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, br)
		}
	}

	return len(stack) == 0
}

// printResult prints whether the expression is balanced.
func printResult(expr string, balanced bool) {
	if balanced {
		log.Printf("%s is balanced.\n", expr)
		return
	}
	log.Printf("%s is not balanced.\n", expr)
}

func main() {
	expr := flag.String("expr", "", "The expression to validate brackets on.")
	flag.Parse()
	printResult(*expr, isBalanced(*expr))
}
