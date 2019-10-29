package main

import "fmt"
type person struct{
    Name string,
    Age  int,
    Sex  string
}

type Vertex struct {
    X, Y int
}

var (
    p = Vertex{1, 2}  // has type Vertex
    q = &Vertex{1, 2} // has type *Vertex
    r = Vertex{X: 1}  // Y:0 is implicit
    s = Vertex{}      // X:0 and Y:0
)

func main(){
    m := person{Name:"David", Age:30, Sex: "Male"}
    fmt.Println(p)
    fmt.Println(p, q, r, s)
}

