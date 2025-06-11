package main

import "fmt"

func doSomething(depth int) {
	if depth > 0 {
		doSomething(depth - 1)
		fmt.Println("depth", depth)
	}
}

func main() {
	doSomething(0)
}
