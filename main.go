package main

import "fmt"

func Cat() string {
	return "miao~~"
}

func main() {
	saying := Cat()
	fmt.Println(saying)
}
