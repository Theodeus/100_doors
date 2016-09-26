package main

import "fmt"

func main() {
    var doorArray [101]bool
    doors := doorArray[:]
    for i := 1; i <= len(doors); i++ {
        for y := 0; y < len(doors); y = y + i {
            doors[y] = !doors[y]
        }
    }
    closed := Filter(doors, func(t bool) bool {
        return !t
    })
    open := Filter(doors, func(t bool) bool {
        return t
    })

    fmt.Printf("closed: ")
    for _, v := range closed[1:] {
        fmt.Printf("%d ", v)
    }
    fmt.Printf("\nopen: ")
    defer fmt.Printf("\n")
    for _, v := range open[1:] {
        fmt.Printf("%d ", v)
    }
}


func Filter(vs []bool, f func(bool) bool) []int {
    vsf := make([]int, 0)
    for i, v := range vs {
        if f(v) {
            vsf = append(vsf, i)
        }
    }
    return vsf
}
