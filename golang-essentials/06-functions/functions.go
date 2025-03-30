package main

import "fmt"

func main() {
	fmt.Println("main function")
	print()

	result := sum2(1, 4)
	println(result)
}

func print() {
	fmt.Println("print function")
	test("test function", 1)
	test("test function", 2)
}

func test(content string, number int) {
	fmt.Println(content, number)
	sum(2, 2, "result: ")
}

func sum(num1, num2 int, msg string) {
	result := num1 + num2
	println(msg, result)
}

func sum2(num1, num2 int) int {
	result := num1 + num2

	return result
}
