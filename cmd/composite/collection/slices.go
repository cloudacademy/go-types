package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age  int
}

func getElement(s []int) string {
	b := strings.Builder{}
	for _, val := range s {
		b.WriteString(" " + fmt.Sprint(val))
	}

	return b.String()
}

type scientists []*person

func (s scientists) print() {
	b := strings.Builder{}
	for _, val := range s {
		b.WriteString(" " + (*val).name)
	}

	fmt.Printf("scientists -> %s\n", b.String())
}

func main() {
	//slice
	slice1 := make([]int, 3, 5) //length 3, capacity 5
	fmt.Println(slice1)
	fmt.Println(slice1[:5])
	slice1[0] = 100
	slice1[1] = 101
	slice1[2] = 102
	//slice1[3] = 103 //compile error - index out of range
	fmt.Printf("slice1 ->%s\n", getElement(slice1))

	slice2 := make([]bool, 3) //length 3, capacity 3
	fmt.Println(slice2)

	slice3 := scientists{{name: "Isaac Newton", age: 25}, {name: "Leonhard Euler", age: 30}}
	fmt.Println(slice3)
	slice3.print()

	fmt.Printf("type: %T\n", slice1)
	fmt.Printf("type: %T\n", slice2)
	fmt.Printf("type: %T\n", slice3)

	slice1 = append(slice1, 0)
	fmt.Printf("slice1 length:%d, capacity:%d\n", len(slice1), cap(slice1))
	slice1 = append(slice1, 1)
	fmt.Printf("slice1 length:%d, capacity:%d\n", len(slice1), cap(slice1))
	slice1 = append(slice1, 2) //increases capacity by another 5
	fmt.Printf("slice1 length:%d, capacity:%d\n", len(slice1), cap(slice1))
	slice1 = append(slice1, 3)
	fmt.Printf("slice1 length:%d, capacity:%d\n", len(slice1), cap(slice1))

	fmt.Printf("slice1 ->%s\n", getElement(slice1))

	slice4 := slice1[1:3]
	fmt.Println(slice4)
}
