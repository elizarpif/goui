package lab3

import (
	"fmt"
	"strings"
)

// 1001 -> x^3 + 1
// 1. Напишите функцию, представляющую элемент из 𝐺𝐹(256)в полиномиальной форме.
func PolynomForm(baseNum string) string {
	var polynom []string
	var i = len(baseNum)

	for _, num := range baseNum {
		i--

		if num == '0' {
			continue
		}

		// последнее число
		if i == 0 {
			polynom = append(polynom, "1")
		} else if i == 1 {
			polynom = append(polynom, "x")
		} else {
			polynom = append(polynom, fmt.Sprintf("x^%d", i))
		}
	}

	return strings.Join(polynom, "+")
}
