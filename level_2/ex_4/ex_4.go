package main

import "fmt"

func main() {
	x := 10
	fmt.Printf("%v:(Dec: %d, Bin: %b, hex: %x)\n", x, x, x, x)
	y := x << 1
	fmt.Printf("%v:(Dec: %d, Bin: %b, hex: %x)", y, y, y, y)

}
