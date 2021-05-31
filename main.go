package main

import "fmt"

func Cat() string {
	return "miao~~"
}

var (
	password = "123456"
)

func main() {

	fmt.Printf("mysql password is %s", password)
	saying := Cat()
	fmt.Println(saying)
}
