package main

import "fmt"

func main() {

	//declare variable (default value is assigned)
	var str string
	fmt.Println("str: " + str)

	//declare and assign value
	var init int = 1
	fmt.Println(init)

	//short declaration operator ":=". it allows to declare and assign the value to a variable 
	a:= 42
	fmt.Println(a)

	//multiple assignments
	b, _, d := 1, 0, "Content"
	fmt.Println(b, d)
	// '_' its a blank identifier
	fmt.Println("_ cannot print this value ")

	non_used_field := "Throw exception if the value is not used"
	fmt.Println("try commenting this line to check: " + non_used_field)
}