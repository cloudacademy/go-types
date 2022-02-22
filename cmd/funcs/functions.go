package main

import (
	"fmt"
	"math/big"
	"strings"
)

type operator func(x int) int
type operator2 func(x interface{}) interface{}

type result struct {
	int
	bool
}

func mapper(op operator, a []int) []int {
	res := make([]int, len(a))
	for i, x := range a {
		res[i] = op(x)
	}
	return res
}

func mapper2(op2 operator2, a []interface{}) []interface{} {
	res := make([]interface{}, len(a))
	for i, x := range a {
		res[i] = op2(x)
	}
	return res
}

func main() {
	pow2 := func(num int) int {
		return num * num
	}

	a := []int{1, 2, 3, 4}
	b := mapper(pow2, a)
	fmt.Println(b)

	c := mapper(func(x int) int { return 10 * x }, b)
	fmt.Println(c)

	fmt.Println()

	upper := func(data interface{}) interface{} {
		return strings.ToUpper(data.(string))
	}

	d := []interface{}{"cloudacademy", "was", "here"}
	e := mapper2(upper, d)
	fmt.Println(e)

	length := func(data interface{}) interface{} {
		return len(data.(string))
	}

	f := mapper2(length, d)
	fmt.Println(f)

	isprime := func(data interface{}) interface{} {
		number := int64(data.(int))
		if big.NewInt(number).ProbablyPrime(0) {
			return result{data.(int), true}
		} else {
			return result{data.(int), false}
		}
	}

	g := mapper2(isprime, []interface{}{1, 2, 3, 4, 5, 6, 7, 8})
	//g := mapper2(isprime, []interface{}{"test", "testing"}) //runtime panic
	fmt.Println(g)
}
