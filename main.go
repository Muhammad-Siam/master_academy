package main

import "fmt"

func main(){
         //primitive datatypes
	//rune, byte= alias
	//imt, float32, string, bool= orginal primitive

	//composite datatypes
	//array
	//slice
	//map
	//struct
	//bracket = []
	//braces/ curly braces = {}
	//paranthesis = ()

         /* var students [3]string
	students[0]="Siam"
	students[1]="Rahat"
	fmt.Println=(students)
	fmt.Println(len(students))
	fmt.Println(students[0])*/
	//implicit= ...
	students := [...]string{"Siam, Rahat, Sajid"}
	fmt.Println(students)
	
	

}