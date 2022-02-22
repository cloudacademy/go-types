package main

import (
	"fmt"
	"sort"
)

type scientist struct {
	field string
	iq    int
}

type scientists map[string]*scientist

type doers func(...interface{})

func func1(p ...interface{}) {
	fmt.Printf("func1 called with param: %s\n", p[0])
}

func func2(p ...interface{}) {
	fmt.Printf("func2 called with params: %d, %t\n", p[0], p[1])
}

func func3(name string) {
	fmt.Printf("func3 called with name param: %s\n", name)
}

func func4(name string, year int) {
	fmt.Printf("func4 called with name: %s, year: %d\n", name, year)
}

func main() {
	var map1 map[string]int
	map1 = make(map[string]int)

	map1["bob"] = 31
	map1["alice"] = 25
	map1["joe"] = 55
	map1["mike"] = 42

	fmt.Println(map1["fast"])

	fmt.Printf("type: %T\n", map1)

	for key, val := range map1 {
		fmt.Printf("%s: %d\n", key, val)
	}

	fmt.Println()

	if val, ok := map1["joe"]; ok {
		fmt.Println(val)
	}

	fmt.Println()

	var map1Keys []string
	for k := range map1 {
		map1Keys = append(map1Keys, k)
	}
	sort.Strings(map1Keys)
	for _, k := range map1Keys {
		fmt.Printf("%s: %d\n", k, map1[k])
	}

	fmt.Println()

	fmt.Printf("length before: %d\n", len(map1))
	delete(map1, "joe")
	fmt.Printf("length after: %d\n", len(map1))

	fmt.Println()

	scientists := make(scientists)
	scientists["Isaac Newton"] = &scientist{field: "Physics", iq: 190}
	scientists["Leonhard Euler"] = &scientist{field: "Mathematics", iq: 195}

	for key, val := range scientists {
		fmt.Printf("Name: %s: IQ: %d\n", key, (*val).iq)
	}

	fmt.Println()

	mapOfFuncs1 := map[string]doers{
		"f1": func1,
		"f2": func2,
	}

	mapOfFuncs1["f1"]("blah")
	mapOfFuncs1["f2"](1, false)

	fmt.Println()

	mapOfFuncs2 := map[string]interface{}{
		"f1": func3,
		"f2": func4,
	}

	for k, v := range mapOfFuncs2 {
		switch k {
		case "f1":
			v.(func(string))("cloudacademy")
		case "f2":
			v.(func(string, int))("cloudacademy", 2022)
		}
	}
}
