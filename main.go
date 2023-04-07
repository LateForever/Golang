package main

import (
	"fmt"
	"math/big"
)

// Syntax 1 - you always need to specify the type of the variable or value you are declaring or both
var example string = "Hello World"

// Syntax 2 - you need to always add value to the variable
func getName() string {
	name := "John"
	surname := "Morgan"

	fmt.Printf("Name: %v and Type is %T \n" , name, name)
	fmt.Printf("Surname: %v and Type is %T" , surname, surname)

	return name + " " + surname
}

// Differences scope between Syntax 1 and Syntax 2
func fibonnaci()  {
	a := big.NewInt(0)
	b := big.NewInt(1)

	n := big.NewInt(0)

	for i := 0; i < 10000; i++ {

		n.Add(a, b)
		a.Set(b)
		b.Set(n)

		fmt.Println(n, " ", i)
	}
}


func getCircumference(radius float64) float64 {
	const PI float64 = 3.14

	return 2 * PI * radius
}



var arr = [5]int{1,2,3,4,5}
var slice = []int{1,2,3,4,5}

func adder(list []int) []int  {

	for i := 0; i < len(list); i++ {
		list[i] += 1
	}

	fmt.Println(list)

	return list
}


func main() {

	fibonnaci()
}
