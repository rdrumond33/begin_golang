package main

import "fmt"

const (
	_ = iota + 2021
	year1
	year2
	year3
	year4
)

func main() {
	fmt.Printf("years: %v,%v,%v,%v", year1, year2, year3, year4)
}
