package stl

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
