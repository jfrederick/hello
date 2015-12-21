package main

import "fmt"
import "strconv"

type Mind struct {
	name string
	age int
	aptitude string
}

func main() {
	// Initialize and manipulate a pointer to a struct.
    john := Mind{"John", 31, "Artifical Intelligence"}
    john_clone := &john
    john_clone.aptitude = "Artifical Emotionality"
    // Examine the memory usage of the Println() function.
    bytes, _ := fmt.Println(john.aptitude)
    fmt.Println(bytes)
    // Toy with strings.
    strBytes := strconv.Itoa(bytes)
    fmt.Println("The previous number of bytes allocated was " + strBytes)
}
