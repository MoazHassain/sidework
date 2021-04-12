package main

import (
	"fmt"
	"reflect"
)

func main() {
	// array
	var student [3]string
	student[0] = "moaz"
	student[1] = "miraz"
	student[2] = "mitul"

	fmt.Println(student)
	fmt.Println(len(student))
	fmt.Println(student[2])
	// short hand way array declaration
	teacher := [2]string{"mostain", "mubir"}
	fmt.Println(teacher)

	// slice
	group01 := student[0:1]
	fmt.Println(group01)
	// declaring elements on slice
	var fruit []string
	fruit = append(fruit, "apple", "grape")
	fmt.Println(fruit)
	fmt.Printf("%T", fruit) //show the data type

	// reflect: to see exactly what type of data it is
	a := reflect.TypeOf(student).Kind().String()
	fmt.Println(a)
	b := reflect.TypeOf(fruit).Kind().String()
	fmt.Println(b)

	// map
	x := make(map[string]string)
	x["name"] = "moaz"
	x["age"] = "22"
	x["address"] = "dhaka"
	x["height"] = "5.4"

	fmt.Println(x)
	fmt.Println(x["address"])
	delete(x, "height")
	fmt.Println(x)
}
