package main

import "fmt"

func main() {
	a := 1 == 2
	b := 1 != 2
	c := 1 > 2
	d := 1 < 2
	e := 1 <= 2
	f := 1 >= 2

	fmt.Printf("a:(%v),b:(%v),c:(%v),d:(%v),e:(%v),f:(%v)", a, b, c, d, e, f)
}
