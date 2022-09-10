package main

import "fmt"

func test2(){
    var l,d,g,f = "f","dd","ggg","ghtre"
    var r,y,i,o int = 5,9,3,1
    fmt.Println("I'm be there")
    fmt.Printf("Thread this")
    fmt.Printf(l)
    fmt.Printf(d)
    fmt.Printf(g)
    fmt.Printf(f)
    fmt.Println(r)
    fmt.Println(y)
    fmt.Println(i)
    fmt.Println(o)
    if(r != 5){
            fmt.Println("Ig")
        }
}

func test(){
    fmt.Println("I'm be there")
    fmt.Printf("Thread this")
}

func main(){
    /* This is my first sample program. */
    fmt.Println("Hello, World!")
    fmt.Printf(" This is Different")
    fmt.Println(" This is Different")
    fmt.Printf(" This is Different")
    test()
    test2()
}