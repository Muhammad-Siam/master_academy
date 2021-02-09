package main

import "fmt"
import "reflect"


func main(){
         //primitive datatypes
	//rune, byte= alias
	//int, float32, string, bool= orginal primitive

	//composite datatypes
	//array
	//slice
	//map
	//struct
	//bracket = []
	//braces/ curly braces = {}
	//paranthesis = ()

        //var students [3]string
	//students[0]="Siam"
	//students[1]="Rahat"

	/*fmt.Println=(students)
	fmt.Println(len(students))
	fmt.Println(students[0])*/
	//implicit= ...
	//students := [...]string{"Siam, Rahat, Sajid"}
	//fmt.Println(students)
	
	/*x:=students[0:2]
	fmt.Println(x)*/

	//x := make([]string, 0)
	
	//b := reflect.TypeOf(x).Kind().String()
	//fmt.Println(b)
	

//  var x [5]int
  //x [4] = 100
 // fmt.Println(x[4])

	//var y [5]float64
	//y [0] = 100
	//y [1] = 200
	//y [3] = 300
	//y [4] = 400
	//fmt.Println(y)
//a := [5]float64 {12, 23, 56, 67, 56}
//fmt.Println(a)

//var friends [3]string

//friends[0]="Nafis,"
//friends[1]="Saifur,"
//friends[2]="Rafi,"

 //x := friends[0:3]
//fmt.Println(x)
//slices without array
//x := make([]string, 0)
//fmt.Println(x)
//append

var fruits []string
fruits=append(fruits, "APPLE", "BANANA")
//fmt.Println(fruits, len(fruits))
//fmt.Printf("%T\n", fruits)
//fmt.Printf("%T", friends) 

a := reflect.TypeOf(fruits).Kind().String()
fmt.Println(a)

var x map[string]int
fmt.Println(x)



}