package main

import "fmt"

func main() {
	// fmt.Print("Enter your age: ")
	// var age int
	// fmt.Scanf("%d", &age)

	// if age < 3 {
	// 	fmt.Println("infant")
	// } else if age > 2 && age < 12 {
	// 	fmt.Println("children")
	// } else if age > 11 && age < 18 {
	// 	fmt.Println("teenage")
	// } else if age > 17 && age < 35 {
	// 	fmt.Println("youth")
	// } else if age > 34 {
	// 	fmt.Println("adult")
	// }

	// fmt.Println("enter age:", age)35

	// by using switch case
	// switch age {
	// case 1, 2:
	// 	fmt.Println("infant")
	// 	fallthrough /*to execute the next case also*/

	// case 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
	// 	fmt.Println("children")
	// case 13, 14, 15, 16, 17, 18:
	// 	fmt.Println("teenage")
	// default:
	// 	fmt.Println("adult")
	// }

	// for loop
	// for i := 20; i >= 5; i-- {
	// 	fmt.Println(i)
	// }

	// student := []string{"soummo", "ashfak", "nabil"}

	// for i, std := range student {
	// 	fmt.Println(i, std)
	// }

	i := 0
	for i < 10 {
		fmt.Println(i, "hello")
		i++
	}
}
