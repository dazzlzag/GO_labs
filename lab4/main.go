package main

import "fmt"

func main() {
	defer fmt.Println("завершение работы") // 5
	// 1
	hello()
	// 2
	func() {
		fmt.Println("Hello, Go2!")
	}()
	// 3
	var f3 func()
	f3 = func() {
		fmt.Println("Hello, Go3!")
	}
	hello3(f3)
	// 4
	var f4 = hello4()
	f4()
	// 5
	hello5()
	// 6
	fmt.Print("0 1 ")
	fibo(0, 1, 3)
}

func hello() {
	fmt.Println("Hello, Go1!")
}

func hello3(func_v func()) {
	func_v()
}

func hello4() func() {
	var f func()
	f = func() {
		fmt.Println("Hello, Go4!")
	}
	return f
}
func hello5() {
	fmt.Println("Hello, Go5!")
}
func fibo(f1 int, f2 int, i int) {
	fmt.Print(f1+f2, " ")
	if i < 23 {
		fibo(f2, f1+f2, i+1)
	}

}
