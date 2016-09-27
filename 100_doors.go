package main

import "fmt"

func main() {
    var doorArray [100]bool
    doors := doorArray[:]
    for i := 1; i <= len(doors); i++ {
        for y := i - 1; y < len(doors); y += i {
            doors[y] = !doors[y]
        }
    }
    closed := Filter(doors, func(t bool) bool {
        return !t
    })
    open := Filter(doors, func(t bool) bool {
        return t
    })

    fmt.Printf("open: ")
    for _, v := range open[:] {
        fmt.Printf("%d ", v + 1)
    }
    fmt.Printf("\nclosed: ")
    defer fmt.Printf("\n")
    for _, v := range closed[:] {
        fmt.Printf("%d ", v + 1)
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
