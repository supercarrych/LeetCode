package LeetCode

func sumOfMultiples(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			sum = sum + i
		}
	}
	return sum
}

func f(n, m int) int {
	return (m + n/m*m) * (n / m) / 2
}
func sumOfMultiples2(n int) int {

	return f(n, 3) + f(n, 5) + f(n, 7) - f(n, 3*5) - f(n, 3*7) - f(n, 5*7) + f(n, 3*5*7)

}
