package main

import "fmt"

func Cat() string {
	return "wang~~"
}

func main() {
	saying := Cat()
	fmt.Println(saying)
}
