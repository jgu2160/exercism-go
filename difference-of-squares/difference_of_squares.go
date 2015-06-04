package diffsquares

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}

func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i < n+1; i++ {
		sum = sum + i*i
	}
	return sum
}

func SquareOfSums(n int) int {
	sum := 0
	for i := 1; i < n+1; i++ {
		sum = sum + i
	}
	return sum * sum
}
