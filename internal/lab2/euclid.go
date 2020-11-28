package lab2

import (
	"errors"
)

func GCD(a int, b int) (int, int, int) {
	x, xx, y, yy := 1, 0, 0, 1

	for b != 0 {
		q := a / b
		a, b = b, a%b

		x, xx = xx, x-xx*q
		y, yy = yy, y-yy*q
	}

	return x, y, a
}

func RecursiveGCD(a int, b int) (int, int, int) {
	if a == 0 {
		return b, 0, 1
	}

	g, y, x := RecursiveGCD(b%a, a)
	return g, x - (b/a)*y, y
}

func RecursiveGCDMod(a, b, m int) (int, int, int) {
	gcd, x, y := RecursiveGCD(a, b)

	x = (x%m + m) % m
	y = (y%m + m) % m
	gcd = (gcd%m + m) % m

	return gcd, x, y
}

// gcd (a, m) = 1
// gcd(a, 256) = 1
// ax + by = gcd(a,b)
// ax + my = 1 // % m
// ax = 1
func ModIn(a int, m int) (int, error) {
	g, x, _ := RecursiveGCD(a, m)
	if g != 1 {
		return 0, errors.New("modular inverse does not exist")
	}
	x = (x%m + m) % m
	return x % m, nil
}

// бинарное возведение в степень
func BinPow(a int, n int, m int) int {
	res := 1

	for n > 0 {
		if n&1 > 0 {
			res *= a
			res %= m
		}

		a *= a
		a %= m
		n >>= 1
	}

	return res
}
