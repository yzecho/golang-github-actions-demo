package main

import "fmt"

func Cat() string {
	return "miao~~"
}

func smallCat() {
}

func main() {
	var password = "12343fsFBKkfasdfbw4352f32R543F"
	fmt.Printf("mysql password is %s", password)
	saying := Cat()
	fmt.Println(saying)
}
