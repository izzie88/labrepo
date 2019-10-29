package main

import "fmt"
type person struct{
    Name string,
    Age  int,
    Sex  string
}

func main(){
    p := person{Name:"David", Age:30, Sex: "Male"}
    fmt.Println(p)
}

