package main

import "fmt"

type payload struct {
	TopSecret bool
	Weight    int
}

type rocket struct {
	Name    string
	Engine  string
	Thrust  float32
	Length  int
	Payload payload
}

func main() {
	r1 := rocket{
		Name:   "Saturn V",
		Engine: "RocketDyne F1",
		Thrust: 25000.00,
		Length: 110,
	}

	fmt.Printf("%v\n", r1)
	fmt.Printf("%+v\n", r1)
	fmt.Printf("%T\n", r1)
	fmt.Println()

	r2 := rocket{
		Name:   "Electron",
		Engine: "Electron",
		Thrust: 15000.00,
		Length: 60,
		Payload: payload{
			TopSecret: true,
			Weight:    100,
		},
	}

	fmt.Printf("%v\n", r2)
	fmt.Printf("%+v\n", r2)
	fmt.Printf("%T\n", r2)
	fmt.Println()
}
