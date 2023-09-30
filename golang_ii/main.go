package main

import "fmt"

func main() {
	/*
		var name = "Muhibullah "
		var lastname string
		var age uint
		fmt.Print("Enter the your last  name : ")
		fmt.Scan(&lastname)
		fmt.Print("Enter your first name : ")
		fmt.Scan(&age)
		fmt.Printf("Welcome Mr %v  %v  your age is %v \n ", name, lastname, age)
		fmt.Printf("Print Type   %T  %T  your age is %T \n", name, lastname, age)
		fmt.Printf("Pointer address  Mr %v  %v  your age is %v \n", &name, &lastname, &age)
	*/
	var students = [50]string{"Sajad", "Ahmad ", "Ali", "Sakhi"}
	fmt.Println(students)

	var empyts [23]string

	fmt.Printf("this length of firstv: %v  array and second array %v \n", len(students), len(empyts))
}
