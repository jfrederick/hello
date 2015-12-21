package main

import "fmt"

type Mind struct {
	name string
	age int
	aptitude string
}

func main() {
    john := Mind{"John", 31, "Artifical Intelligence"}
    fmt.Println(john.aptitude)
}
