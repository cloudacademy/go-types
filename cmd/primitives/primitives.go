package main

import "fmt"

func main() {
	var enabled bool = true
	fmt.Printf("%t\n", enabled)

	var name string = "CloudAcademy"
	fmt.Printf("%s\n", name)
	fmt.Println()

	var (
		testInt1 int   = 5
		testInt2 int8  = 1<<7 - 1  //signed  8-bit integers (-128 to 127)
		testInt3 int16 = 1<<15 - 1 //signed 16-bit integers (-32768 to 32767)
		testInt4 int32 = 1<<31 - 1 //signed 32-bit integers (-2147483648 to 2147483647)
		testInt5 int64 = -1 << 63  //signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
	)

	fmt.Printf("%d\n", testInt1)
	fmt.Printf("%d\n", testInt2)
	fmt.Printf("%d\n", testInt3)
	fmt.Printf("%d\n", testInt4)
	fmt.Printf("%d\n", testInt5)
	fmt.Println()

	var (
		unsignedInt1 uint
		unsignedInt2 uint8  = 1<<8 - 1  //unsigned  8-bit integers (0 to 255)
		unsignedInt3 uint16 = 1<<16 - 1 //unsigned 16-bit integers (0 to 65535)
		unsignedInt4 uint32 = 1<<32 - 1 //unsigned 16-bit integers (0 to 4294967295)
		unsignedInt5 uint64 = 1<<64 - 1 //unsigned 16-bit integers (0 to 18446744073709551615)
	)

	fmt.Printf("%d\n", unsignedInt1)
	fmt.Printf("%d\n", unsignedInt2)
	fmt.Printf("%d\n", unsignedInt3)
	fmt.Printf("%d\n", unsignedInt4)
	fmt.Printf("%d\n", unsignedInt5)
	fmt.Println()

	heart := '\u2764'

	var (
		testByte byte = 99    //alias for uint8
		testRune rune = 10084 //alias for uint32
	)

	heart2 := string([]byte{226, 157, 164})

	fmt.Printf("heart %c\n", heart)
	fmt.Printf("heart2 %s\n", heart2)
	fmt.Printf("%c\n", testByte)
	fmt.Printf("%c\n", testRune)
	fmt.Println([]byte("❤ golang"))
	fmt.Println([]rune("❤ golang"))
	fmt.Println()

	msg := "cloudacademy ❤ golang"

	for idx, e := range msg {
		fmt.Printf("%c pos: %d\n", e, idx)
	}
	fmt.Println()

	var (
		pi     float32 = 3.14159  //IEEE-754 32-bit floating-point numbers
		float1 float64 = 1.7e+308 //IEEE-754 64-bit floating-point numbers
	)

	fmt.Printf("Type: %T, value: %v\n", pi, pi)
	fmt.Printf("Type: %T, value: %v\n", float1, float1)
	fmt.Println()
}
