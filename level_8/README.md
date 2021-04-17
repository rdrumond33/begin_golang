# Level_8

## Ex_1

- Partindo do código abaixo, utilize marshal para transformar []user em JSON.

```go
package main

import (
  "fmt"
)

type user struct {
  first string
  age   int
}

func main() {
  u1 := user{
    first: "James",
    age:   32,
 }

  u2 := user{
    first: "Moneypenny",
    age:   27,
   }

  u3 := user{
    first: "M",
    age:   54,
  }

 users := []user{u1, u2, u3}

  fmt.Println(users)

  // your code goes here
}
```

## EX_2

- Partindo do código abaixo, utilize unmarshal e demonstre os valores.

```go
package main

import (
  "fmt"
)

func main() {
  s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
  fmt.Println(s)
}
```

## Ex_3

- Partindo do código abaixo, utilize NewEncoder() e Encode() para enviar as informações como JSON para Stdout.

```go
package main

import (
  "fmt"
)

type user struct {
  First   string
  Last    string
  Age     int
  Sayings []string
}

func main() {
  u1 := user{
    First: "James",
    Last:  "Bond",
    Age:   32,
    Sayings: []string{
      "Shaken, not stirred",
      "Youth is no guarantee of innovation",
      "In his majesty's royal service",
    },
  }

  u2 := user{
    First: "Miss",
    Last:  "Moneypenny",
    Age:   27,
    Sayings: []string{
      "James, it is soo good to see you",
      "Would you like me to take care of that for you, James?",
      "I would really prefer to be a secret agent myself.",
    },
  }

  u3 := user{
    First: "M",
    Last:  "Hmmmm",
    Age:   54,
    Sayings: []string{
      "Oh, James. You didn't.",
      "Dear God, what has James done now?",
      "Can someone please tell me where James Bond is?",
    },
  }

  users := []user{u1, u2, u3}

  fmt.Println(users)

  // your code goes here

}
```

## Ex_4

- Partindo do código abaixo, ordene a []int e a []string.

```golang
package main

import (
  "fmt"
)

func main() {
  xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
  xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

  fmt.Println(xi)
  // sort xi
  fmt.Println(xi)

  fmt.Println(xs)
  // sort xs
  fmt.Println(xs)
}
```

## Ex_5

- Partindo do código abaixo, ordene os []user por idade e sobrenome.
- Os valores no campo Sayings devem ser ordenados tambem, e demonstrados de maneira esteticamente harmoniosa.

```golang
package main

import (
  "fmt"
)

type user struct {
  First   string
  Last    string
  Age     int
  Sayings []string
}

func main() {
  u1 := user{
    First: "James",
    Last:  "Bond",
    Age:   32,
    Sayings: []string{
      "Shaken, not stirred",
      "Youth is no guarantee of innovation",
      "In his majesty's royal service",
    },
  }

  u2 := user{
    First: "Miss",
    Last:  "Moneypenny",
    Age:   27,
    Sayings: []string{
      "James, it is soo good to see you",
      "Would you like me to take care of that for you, James?",
      "I would really prefer to be a secret agent myself.",
    },
  }

  u3 := user{
    First: "M",
    Last:  "Hmmmm",
    Age:   54,
    Sayings: []string{
      "Oh, James. You didn't.",
      "Dear God, what has James done now?",
      "Can someone please tell me where James Bond is?",
    },
  }

  users := []user{u1, u2, u3}

  fmt.Println(users)

  // your code goes here

}

```
