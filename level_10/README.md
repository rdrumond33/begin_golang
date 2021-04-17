
# level 10

## Ex_1

- Faça esse código funcionar
  - Usando uma função anônima auto-executável
  - Usando buffer
  
```golang
package main

import (
  "fmt"
)

func main() {
  c := make(chan int)

  c <- 42

  fmt.Println(<-c)
}
```

Ex_2
- Faça esse código funcionar

```golang
package main

import (
  "fmt"
)

func main() {
  cs := make(chan<- int)

  go func() {
    cs <- 42
  }()
  fmt.Println(<-cs)

  fmt.Printf("------\n")
  fmt.Printf("cs\t%T\n", cs)
}

``` 

## Ex_4

- use um select statement para demonstrar os valores do canal.
- no codigo abaixo

```golang
package main

import (
  "fmt"
)

func main() {
  q := make(chan int)
  c := gen(q)

  receive(c, q)

  fmt.Println("about to exit")
}

func gen(q <-chan int) <-chan int {
  c := make(chan int)

  for i := 0; i < 100; i++ {
    c <- i
  }

  return c
}
```