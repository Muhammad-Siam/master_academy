package main

import "fmt"
//Example -1
/*func add(x int, y int) int{
	
	//body inside braces
	r := x+y
	return r
}*/



//Example-2
/*func add(x, y int) int{
	
	//body inside braces


	r := x+y
	return r

}*/

//Example-3 = name return value, 4
//(x, y int) = parameter
//return r = return

func add(x, y int) (r int){
	
	//body inside braces


	r = x+y
	return 
}

//returning multiple
func rectangle(l, b int) (area, parameter int){

	parameter = 2 * (l + b)

	area = l * b

	return  //return statement without specify variable name

}

//Pointer = store memory addresses of variable

func update(a *int, t *string) {

	*a = *a + 5        //defrencing pinter address
	*t = *t + "Siam"

	return


}





func main(){

	
		//x:= add(10, 30)

	//a,p := rectangle(10, 10)

	//fmt.Println(a,p)

	//a:=10

	//t:= "Siam"	
	//update(&number, &name)

	//fmt.Println(number, name)
	//anonymous function = যখন কোনো ফাংশনের নাম না দিইয়ে ভিন্ন উপায়ে ফাংশন ডিক্লেয়ার করা।
	add := func(l, b int) (r int){

	r = l + b

	return
	}(10, 30)
	
	fmt.Println(add)
	//fmt.Println(add(10, 20))


}