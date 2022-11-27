package test

import (
	"fmt"
	"math/rand"
	"time"
)

var z string
var a int

type hotdog int

var b hotdog

func try() {
	a = 42
	b = 43
	fmt.Println(z)
	fmt.Println("printing the value of z", z, "ending")
	fmt.Printf("%T", z)
	a = 42
	fmt.Println(a)
}

func main() {
	rand.Seed(time.Now().Unix())
	x := rand.Intn(100)
	fmt.Println(x)
	var y = 100
	fmt.Println(y)
	fmt.Println("printing the value of y", y, "ending")
	fmt.Printf("%T", z)
	a = 42
	fmt.Println(a)

}
