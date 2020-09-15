package main

import (
	"fmt"
	"stacker/stack"
)

func main() {
	var haystack stack.Stack
	haystack.Push("hay")
	haystack.Push(-15)
	haystack.Push([]string{"pin", "clip", "needle"})
	haystack.Push(97.99)

	for {
		item, err := haystack.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}
}