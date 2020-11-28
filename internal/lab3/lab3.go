package lab3

import (
	"fmt"
	"strings"
)

func form(num int) []int {
	var polynom []int

	for num > 0 {
		polynom = append(polynom, num%2)
		num = num >> 1
	}

	return polynom
}

type BinaryPolynom struct {
	numbers []int
}

func NewBinaryPolynom(num int) *BinaryPolynom {
	return &BinaryPolynom{numbers: form(num)}
}

func (bp *BinaryPolynom) String() string {
	var str []string

	for i, monom := range bp.numbers {
		if monom == 0 {
			continue
		}

		var s string

		switch i {
		case 0:
			s = "1"
		case 1:
			s = "x"
		default:
			s = fmt.Sprintf("x^%d", i)
		}

		str = append(str, s)
	}

	return strings.Join(str, "+")
}


//func (p1 *BinaryPolynom) Multiply(p2 *BinaryPolynom) error{
//
//}

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
