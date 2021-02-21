package  main                            //Simple BMI project

import "fmt"


func bmi(m *float32, h *float32) float32{

	*m = *m
	*h =  *h * *h
	
	return *m / *h


}

func main(){

	fmt.Println("Enter your Weight(kg) & Height(m):")

	var m float32
	var h float32
	
	fmt.Scanln(&m, &h)

	fmt.Println(bmi(&m, &h))

	fmt.Println("Underweight = <18.5")

	fmt.Println("Perfect = 18.5 - 24.9")

	fmt.Println("Overweight= 24.9<")



}