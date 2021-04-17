package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Seu sitema operacional e:", runtime.GOOS)
	fmt.Println("Sua arctetura e:", runtime.GOARCH)
}
