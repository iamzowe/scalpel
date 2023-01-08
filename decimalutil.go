package scalpel

import (
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
)

func calc(a string, b string, op int, pcs int) (string, error) {
	decimal.DivisionPrecision = pcs
	ad, err1 := decimal.NewFromString(a)
	bd, err2 := decimal.NewFromString(b)
	if err1 != nil || err2 != nil {
		return "0.00", errors.New(fmt.Sprintf("Calc fail,err1=%v,err2=%v", err1, err2))
	}

	if op == 1 {
		ad = ad.Add(bd)
	} else if op == -1 {
		ad = ad.Sub(bd)
	} else if op == 2 {
		ad = ad.Div(bd)
	} else if op == 3 {
		ad = ad.Mul(bd)
	}

	return ad.String(), nil
}

/**
a+b
*/
func Add(a string, b string) (r string, err error) {
	return calc(a, b, 1, 20)
}

/**
a-b
*/
func Sub(a string, b string) (r string, err error) {
	return calc(a, b, -1, 20)
}

/**
a/b
*/
func Div(a string, b string) (r string, err error) {
	return calc(a, b, 2, 20)
}

/**
a*b
*/
func Mul(a string, b string) (r string, err error) {
	return calc(a, b, 3, 20)
}

/**
s1 == s2
*/
func Equal(s1 string, s2 string) bool {
	a, err1 := decimal.NewFromString(s1)
	b, err2 := decimal.NewFromString(s2)

	if err1 != nil || err2 != nil {
		return false
	}
	return a.Equal(b)
}

/**
s1 < s2
*/
func LessThan(s1 string, s2 string) bool {
	a, err1 := decimal.NewFromString(s1)
	b, err2 := decimal.NewFromString(s2)

	if err1 != nil || err2 != nil {
		return false
	}
	return a.LessThan(b)
}

/**
s1 <= s2
*/
func LessThanOrEqual(s1 string, s2 string) bool {
	a, err1 := decimal.NewFromString(s1)
	b, err2 := decimal.NewFromString(s2)

	if err1 != nil || err2 != nil {
		return false
	}
	return a.LessThanOrEqual(b)
}

/**
s1 > s2
*/
func GreaterThan(s1 string, s2 string) bool {
	a, err1 := decimal.NewFromString(s1)
	b, err2 := decimal.NewFromString(s2)

	if err1 != nil || err2 != nil {
		return false
	}
	return a.GreaterThan(b)
}

/**
s1 >= s2
*/
func GreaterThanOrEqual(s1 string, s2 string) bool {
	a, err1 := decimal.NewFromString(s1)
	b, err2 := decimal.NewFromString(s2)

	if err1 != nil || err2 != nil {
		return false
	}
	return a.GreaterThanOrEqual(b)
}
