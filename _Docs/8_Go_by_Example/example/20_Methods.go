package main

import "fmt"

type rectt struct {
    width, height int
}

func (r *rectt) area() int {
    return r.width * r.height
}

func (r rectt) perim() int {
    return 2*r.width + 2*r.height
}

func methods() {
    r := rectt{width: 10, height: 5}

    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}