package main

func main() {
	//rSum, rSubtr, rMult := maths(1, 4)
	//println(rSum, rSubtr, rMult)

	_, rSubtr, _ := maths(1, 4)
	println(rSubtr)
}

func maths(num1, num2 int) (sum, subtr, mult int) {
	sum = num1 + num2
	subtr = num1 - num2
	mult = num1 * num2

	return
}
