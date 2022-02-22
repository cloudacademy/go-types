package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	fmt.Println()

	strNum1 := "1100"

	testInt1, err := strconv.Atoi(strNum1)
	fmt.Println(testInt1, err, reflect.TypeOf(testInt1))

	testInt2, err := strconv.ParseInt(strNum1, 0, 16)
	fmt.Println(testInt2, err, reflect.TypeOf(testInt2))

	testInt3, err := strconv.ParseInt(strNum1, 2, 8)
	fmt.Println(testInt3, err, reflect.TypeOf(testInt3))

	i := 10
	s1 := strconv.FormatInt(int64(i), 10)
	fmt.Println(s1, reflect.TypeOf(s1))

	s2 := strconv.Itoa(i)
	fmt.Println(s1, reflect.TypeOf(s2))
}
