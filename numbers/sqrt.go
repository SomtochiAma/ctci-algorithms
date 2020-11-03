package numbers

func SquareRoot(n int) int{
	return sqrt(n, 1, n)
}

func sqrt(n, min, max int) int{
	if max < min {
		return -1
	}

	guess := (min + max)/2

	if guess*guess < n {
		return sqrt(n, guess+1, max)
	}

	if guess*guess > n {
		return sqrt(n, min, guess-1)
	}

	return guess
}