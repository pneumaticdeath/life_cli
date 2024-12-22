package main

import (
  "fmt"
  "github.com/pneumaticdeath/golife"
  "os"
  "strings"
)

func display(g *golife.Game) {
    min_cell, max_cell := g.BoundingBox()
    width := max_cell.X - min_cell.X + 1

    fmt.Printf("Generation %d\n", g.Generation)
    if max_cell.X >= min_cell.X {
        fmt.Println("Bounding Box",min_cell,"to",max_cell)
        var x, y int64
        fmt.Printf("+%s+\n", strings.Repeat("-", int(width * 2 - 1)))
        for y = min_cell.Y; y <= max_cell.Y; y++ {
            cell_line := make([]string, width)
            for x = min_cell.X; x <= max_cell.X; x++ {
                if g.Population[golife.Cell{x, y}] {
                    cell_line[int(x - min_cell.X)] = "*"
                } else {
                    cell_line[int(x - min_cell.X)] = " "
                }
            }
            fmt.Printf("|%s|\n", strings.Join(cell_line, " "))
        }
        fmt.Printf("+%s+\n", strings.Repeat("-", int(width * 2 - 1)))
    }
}

func main() {
    var g *golife.Game

    if len(os.Args) > 1 {
        g = golife.Load(os.Args[1])
    } else {
        cells := []golife.Cell{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {2, 1}}
        g = &golife.Game{}
        g.Init()
        g.Population.Add(cells)
    }

    for i := 0; i < 100 && len(g.Population) > 0; i++ {
        display(g)
        g.Next()
    }
    display(g)
}
