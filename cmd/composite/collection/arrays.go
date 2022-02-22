package main

import (
	"fmt"
)

func printArray(arr *[12]string) {
	for _, letter := range arr {
		fmt.Printf("%s", letter)
	}
	fmt.Println()
}

type person struct {
	name string
	age  int
}

func main() {
	//array
	cloudacademyArr := [12]string{"c", "l", "o", "u", "d", "a", "c", "a", "d", "e", "m", "y"}
	fmt.Printf("type: %T\n", cloudacademyArr)
	var anotherArr = [...]int{3, 1, 4, 1, 5, 9} //defaults to the number of elements
	fmt.Printf("type: %T\n", anotherArr)

	//slice
	cloudacademySlice := []string{"c", "l", "o", "u", "d", "a", "c", "a", "d", "e", "m", "y"}
	fmt.Printf("type: %T\n", cloudacademySlice)

	for idx, letter := range cloudacademyArr {
		fmt.Printf("index %d, %s\n", idx, letter)
	}

	printArray(&cloudacademyArr)
	//printArray(anotherArr) //compile error - type is wrong length

	fmt.Printf("array length %d\n", len(cloudacademySlice))
	cloudacademySlice = append(cloudacademySlice, "y")
	cloudacademySlice = append(cloudacademySlice, "y")
	fmt.Printf("array length %d\n", len(cloudacademySlice))

	emptyArr := [10]int{}
	emptyArr[0] = 1
	emptyArr[1] = 2
	//emptyArr[2] = "blah" //compile error - wrong type
	//emptyArr[10] = 2  //compile error - idx out of bounds

	allSortsArr := [10]interface{}{}
	allSortsArr[0] = "blah"
	allSortsArr[1] = 100
	allSortsArr[2] = true
	allSortsArr[3] = &person{name: "Isaac Newton", age: 25}

	for idx, val := range allSortsArr {
		fmt.Printf("index %d, type %T\n", idx, val)

		switch v := val.(type) {
		case int:
			fmt.Printf("item %T is type int\n", v)
		case string:
			fmt.Printf("item %T is type string\n", v)
		case bool:
			fmt.Printf("item %T is type boolean\n", v)
		case *person:
			fmt.Printf("item %T is type pointer to person\n", v)
		default:
			fmt.Println("unkown type")
		}
	}

	scientists := [2]*person{}
	scientists[0] = &person{name: "Isaac Newton", age: 25}
	scientists[1] = &person{name: "Leonhard Euler", age: 30}

	for idx, s := range scientists {
		fmt.Printf("index %d, person %T\n", idx, *s)
		fmt.Printf("index %d, person %s\n", idx, *&s.name)
		fmt.Printf("index %d, person %s\n", idx, (*s).name)
	}
}
