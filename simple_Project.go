package main

import "fmt"

func main(){
	
	fmt.Println("Enter your name & age: ")
	var name string

	var age int

	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("Your name %s & age is %d", name, age)
}