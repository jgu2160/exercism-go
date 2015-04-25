package diffsquares

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}

func SumOfSquares(n int) int {
	a := Narray(n)
	sum := 0
	for _, value := range a {
		sum = sum + value*value

	}
	return sum
}

func SquareOfSums(n int) int {
	a := Narray(n)
	sum := 0
	for _, value := range a {
		sum = sum + value

	}
	return sum * sum
}

func Narray(n int) []int {
	var a []int
	for i := 1; i < n+1; i++ {
		a = append(a, i)
	}
	return a
}
