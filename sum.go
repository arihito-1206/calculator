package calculator

var logMessage = "[LOG]"

// Version of calculator
var Version = "1.0"

func internalSum(num int) int {
	return num - 1
}

// Sum two integer numbers
func Sum(num1, num2 int) int {
	return num1 + num2
}
