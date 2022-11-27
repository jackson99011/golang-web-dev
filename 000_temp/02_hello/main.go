package main

import "fmt"

var x int = 42
var y string = "john"
var z bool = true

func main() {
	//s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)
	//fmt.Println(s)
	//for i := 33; i < 122; i++ {
	//	fmt.Printf("%v\t%#x\t%#U\n", i, i, i)
	//}
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
