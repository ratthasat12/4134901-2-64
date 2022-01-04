package main

import (
	"fmt"

	"ratthasat/lab7/structs"
)

func main (){
	const rname = "Ratthasat"
	var age = 21

	fmt.Printf("Hello..%s : %v \n",rname,age)

	if age > 21 {
		fmt.Printf(" > 21")
	}else{
		fmt.Printf(" < 21")
	}

	switch age{
	case 21:
		fmt.Printf("\n age is 21")
	case 25:
		fmt.Printf("\n age is 25")
	}

	var count = 0
	for count < 5{
		fmt.Printf("%v",count)
		count++
	}

	result := add(10,20)
	fmt.Printf("\n%v",result)

	var ptr *int 
	fmt.Printf("\n%v",&ptr)

	city := "Buriram"
	var cityPointer *string
	cityPointer = &city
	fmt.Println(*cityPointer)

	structs.Learn()
}

func add(a int,b int) int{
	c := a+b
	return c
}

	