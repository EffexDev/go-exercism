package diffsquares

func SquareOfSum(n int) int {
    total := 0
    for i := 1; i <= n; i++ {
        total += i
    }
    return total * total
}

func SumOfSquares(n int) int {
    total := 0
	for i := 1; i <= n; i++ {
        total += i * i
    }
    return total
}

func Difference(n int) int {
    if SquareOfSum(n) >  SumOfSquares(n) {
        return SquareOfSum(n) - SumOfSquares(n)
    } else {
        return SumOfSquares(n) - SquareOfSum(n)
    }
}
