package math

var factorialArray = [...]int{1, 1, 2, 6, 24, 120, 720, 5040}

func factorial(n int) int {
	if n < len(factorialArray) {
		return factorialArray[n]
	}
	return n * factorial(n-1)
}
