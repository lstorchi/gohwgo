package main

import (
    "fmt"
    "math/rand"
)

func simpleadd (input int) int {
  if input < 5 {
    return input
  } else {
    return input * 2
  }
}

func main() {

  for i := 0; i < 10; i++ {
    gn := rand.Intn(10)
    fmt.Println("generated  number is", gn)
    uv := simpleadd(gn)
    fmt.Println("elaborated number is", uv)
  }

}
