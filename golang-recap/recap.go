package main

import "fmt"

// type person struct {
// 	name   string
// 	height float32
// 	age    int
// }

func main() {

	fmt.Println("hello")

	// part-one

	// var a int = 16
	// fmt.Println(a)

	// var b int
	// b = 15
	// fmt.Println(b)

	// c := "moaz"
	// fmt.Println(c)

	// var d [5]int
	// d[0] = 1
	// d[1] = 2
	// d[2] = 3
	// d[3] = 4
	// d[4] = 5
	// fmt.Println(d)
	// fmt.Println(d[3])

	// e := [2]string{"moaz", "mostain"}
	// fmt.Println(e)

	// fmt.Println(e[1:2])

	// var f *int = &b
	// fmt.Println(f)

	// g := &c
	// fmt.Println(*g)

	// *g = "mitul"
	// fmt.Println(c)

	// h := [...]int{1, 2, 3}
	// fmt.Println(h)

	// i := d[0:3]
	// fmt.Println(i)

	// pet := [...]string{"cat", "dog"}
	// fmt.Println(pet)
	// i = append(i, 3, 4)
	// fmt.Printf("%T\n", i)
	// fmt.Printf("%T", pet)

	// j := reflect.TypeOf(c).Kind().String()
	// fmt.Println("\n", j)

	// k := make(map[string]string)
	// k["name"] = "moaz"
	// k["height"] = "5'3"
	// k["age"] = "22"
	// fmt.Println(k, "\n", k["age"])
	// delete(k, "age")
	// fmt.Println(k)

	// part-two

	// var p1 person
	// p1.name = "moaz"
	// p1.height = 5.4
	// p1.age = 22

	// fmt.Println(p1)
	// fmt.Println(p1.age)

	// fruit := struct {
	// 	name    string
	// 	color   string
	// 	ammount int
	// }{
	// 	name:    "apple",
	// 	color:   "red",
	// 	ammount: 3,
	// }
	// fmt.Println(fruit)

	// part-three

	// fmt.Print("enter your name: ")
	// var name string
	// fmt.Scanf("%s", &name)
	// fmt.Println(name)

	// fmt.Print("enter age: ")
	// var age int
	// fmt.Scanf("%d", &age)
	// if age < 3 {
	// 	fmt.Println("infant")
	// } else if age > 2 && age < 10 {
	// 	fmt.Println("child")
	// } else if age > 9 && age < 25 {
	// 	fmt.Println("young boy")
	// } else if age > 24 {
	// 	fmt.Println("adult")
	// } else {
	// 	fmt.Println("old")
	// }

	// var i int
	// for i = 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// product := [...]string{"ram", "monitor", "mouse"}
	// for i, value := range product {
	// 	fmt.Println(i, value)
	// }

	j := 1
	for j < 10 {
		fmt.Println(j, "happy")
		j++
	}
}
