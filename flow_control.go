package main

import (
	"fmt"


)

func main(){
	
	fmt.Println("Enter your age:")
	var age int
	fmt.Scanf("%d", &age)
	fmt.Println("Your age is:", age)

	//if boolean_expression {
	//logic or statement here
	//}
	
	if age  < 3 {
	fmt.Println("You are an Infant")
 	}else if age > 2 && age <= 12 { 
		//2 to 12
		fmt.Println(" you are a children")
		}else if age>12 && age <= 19{
	fmt.Println("You are a teen age")
	}else{
	fmt.Println("You are an adult")
	}
	//fixed value
	fmt.Println("Enter your age:")
	var age1 int
	fmt.Scanf("%d", &age1)
	fmt.Println("Your age is:", age1)

	switch age1 {
	case 1,2:
		fmt.Println("Infant")
		fallthrough

	case 3,4,5,6,7,8,9,10,11,12:
		fmt.Println("Children")

	case 13,14,15,16,17,18:
		fmt.Println("teen age")

	default:
		fmt.Println("adult")
	}

	//for loop
	//1,2,3,4,5,6,7,8,9

	//i=i+1 == i++

	for i:=1; i<=9; i++ {

		fmt.Println(i)
}
     //array string literals
	//for range loop

	students := []string{"Siam", "Rahat", "Sajid"}

	for i, std := range students {

		fmt.Println(i, std)
}

//While loop
 i:=0
 for i<10{

	fmt.Println("Hello")
 i++
}
	

	
	

}