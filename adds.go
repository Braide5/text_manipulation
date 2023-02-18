package main

import "fmt"

func main() {
	var name string 
	var year int 
	fmt.Println("please enter your name")
	fmt.Scanln(&name)
	fmt.Println("please enter your year of birth")
	fmt.Scanln(&year)
	fmt.Println(name, "is", 2023-year, "old")
}