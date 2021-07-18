package stl

func Lcm(a, b int) int {
	return a / Gcd(a, b) * b
}
