package main

import "fmt"
import "math"

// 接口定义。
type geometry interface {
    area() float64
    perim() float64
}

// 结构体。
type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

// 结构体中有接口所以对应的所有方法，它就实现了接口。
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

// 结构体圆的接口实现
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

// 使用接口类型来测试对象。
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    measure(r)
    measure(c)
}
