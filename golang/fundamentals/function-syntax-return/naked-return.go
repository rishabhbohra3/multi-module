package main

import "fmt"

//normal function with return values 
func swap(x, y string) (string, string) {
	return y, x
}

//naked return function: variable name is specified in the return statement
func concat(x, y string) (concat_str string) {
	concat_str = x + " " + y
	return
}

func main() {
	a, b := swap("Rishabh", "Bohra")
	fmt.Println(a, b)
	fmt.Println(concat("Rishabh", "Bohra"))
}