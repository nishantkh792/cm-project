package main

import("fmt")

func main(){
	calculator:=Calculator{Radius:10}
	fmt.Println(calculator.Perimeter())
	fmt.Println(calculator.Area())





	circle:=Circle{Radius:10}
	fmt.Println(circle.Perimeter())
	fmt.Println(circle.Area())

	rectangle:=Rectangle{Radius:10,Height:20}
	fmt.Println(rectangle.Perimeter())
	fmt.Println(rectangle.Area())

	
}