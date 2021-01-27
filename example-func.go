package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func foo() {
	fmt.Println("The square root of 4 is", math.Sqrt(4))
}

func main() {
	foo()
	fmt.Println("It is currently:", time.Now())
	fmt.Println("A number from 1-100", rand.Intn(100))

}
