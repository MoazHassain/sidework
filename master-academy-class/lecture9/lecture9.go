package main

import "fmt"

func main() {

	fmt.Println("Enter name & age:")
	var name string
	var age int

	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("Your name is %s & age is %d", name, age)
}
