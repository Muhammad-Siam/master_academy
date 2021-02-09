package main

import "fmt"



func main(){
//key = value
//other language = null
//golang = nil
       x := make(map[string]string)
	x["Name"]="SIAM"
	x["age"]="5.7"
	delete(x, "age")
	fmt.Println(x)



}