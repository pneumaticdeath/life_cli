package main

import (
  "fmt"
  "github.com/pneumaticdeath/golife"
)

func main() {

   pop := make(golife.Population)
   cells := []golife.Cell{{0, 0}, {0, 1}, {0, 2}}

   pop.Add(cells)

   for i := 0; i < 10; i++ {
      fmt.Println(pop)
      pop = pop.Step()
    }
    fmt.Println(pop)
}
