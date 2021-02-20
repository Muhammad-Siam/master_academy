package main

import "fmt"

func main(){
         //pointer's 0 value is nill 
	var a *int
	var b int =100
	a = &b
	
	fmt.Println(b, a, *a)

	//b:= 10

	//c:= 11
	

	//a=&b
	//a=&c
	//*a=5
	//*a=10
	//*a=1000
	//fmt.Println(*a)
}

//value print korle dereferencing
//adreess print asake referencing bole

//Pointer:

//অন্য কোনো ভ্যারিয়েবল এর মেমোরি এড্রেস স্টোর করতে ব্যাবহার করা হয় ।
//জিরো ভ্যালু হচ্ছে নিল।