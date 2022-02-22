package main

import "fmt"

type person struct {
	name string
	age  int
}

type robot struct {
	model string
}

func (p person) speak(data string) {
	fmt.Printf("person speaking...'%s'\n", data)
}

func (r robot) communicate(data string) {
	fmt.Printf("robot communicating... '%s'\n", data)
}

func process(object interface{}, data string) {
	if p, ok := object.(person); ok {
		p.speak(data)
	}

	if r, ok := object.(robot); ok {
		r.communicate(data)
	}

	switch obj := object.(type) {
	case nil:
		fmt.Println("obj is nil")
	case person:
		fmt.Printf("person obj: %T\n", obj)
	case robot:
		fmt.Printf("robot obj: %T\n", obj)
	default:
		fmt.Println("obj type unknown")
	}

	fmt.Println()
}

func main() {
	fmt.Println()

	p1 := person{name: "John Doe", age: 20}
	r1 := robot{model: "T-800"}

	process(p1, "type assertion performed")
	process(r1, "type assertion performed")
}
