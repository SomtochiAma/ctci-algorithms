package numbers

// import "fmt"

func IsPrime(n int) bool{
	for x := 2; x*x <= n; x++ {
		if n % x == 0 {
			return false
		}
	} 

	return true
}