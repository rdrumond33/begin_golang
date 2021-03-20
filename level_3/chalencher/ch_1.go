package main

import "fmt"

func main() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("Decimal: %d, Hex: %x, Unicode: %#U​,\t\n​", i, i, i)
	}
}
