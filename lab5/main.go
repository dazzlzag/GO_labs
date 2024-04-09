package main

import "fmt"

// 6
type square int

func main() {
	// 1
	var strAddress *string
	// 2
	var stringParam1 = "test string"
	fmt.Println(stringParam1)
	fmt.Println(&stringParam1)
	// 3
	strAddress = &stringParam1
	*strAddress = "new string"
	fmt.Println(stringParam1)
	// 4
	// дубль 2
	// 5
	var localV int = 10
	change(&localV)
	fmt.Println(localV)
	// 6
	var s square = 25
	fmt.Println(s)
	// 7
	var s2 square = 30
	s2 += 15
	fmt.Println(s2)
	// 8
	var s3 square = 34
	s3 += 10
	fmt.Println(squarePrint(s3))
}

func change(a *int) {
	*a = 555
}

func squarePrint(s square) string {
	return fmt.Sprintf("%d м2", s)
}
