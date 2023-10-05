package main

import "fmt"

// lang := "GoLang" // This method llowed only inside method

// Constants
const lang string = "GoLang" // or const lang = "GoLang"

// Public -> First Letter Capital
const Version = "go version go1.21.1 windows/amd64"

func main() {
	fmt.Println("Variables!!")
	// String
	var name string = "Jainam"
	fmt.Println(name)
	fmt.Printf("Variable Type: %T \n", name)

	// boolean
	var isTrue bool = false
	// fmt.Println(isTrue)
	fmt.Printf("Variable Type: %T Value: %v\n", isTrue, isTrue)

	// Int
	var age int = 23
	fmt.Println(age)
	fmt.Printf("Variable Type: %T \n", age)

	var uint8Var uint8 = 255
	fmt.Println(uint8Var)
	fmt.Printf("Variable Type: %T \n", uint8Var)

	// float
	var loan32 float32 = 7890.8790876487690654839
	fmt.Println(loan32)
	fmt.Printf("Variable Type: %T \n", loan32)

	var loan64 float64 = 7890.8790876487690654839
	fmt.Println(loan64)
	fmt.Printf("Variable Type: %T \n", loan64)

	// default values
	var anotherVar string
	fmt.Println(anotherVar)
	fmt.Println(anotherVar == "")
	fmt.Printf("Variable Type: %T \n", anotherVar)
	var anotherVar2 int
	fmt.Println(anotherVar2)
	fmt.Printf("Variable Type: %T \n", anotherVar2)

	// implicit type
	var github = "https://github.com/JainamZobaliya/"
	fmt.Println(github)
	fmt.Printf("Variable Type: %T \n", github)
	// github = 5 // COMPILE TIME ERROR: cannot use 5 (type int) as type string in assignment

	// no var style
	fullName := "Jainam Zobaliya"
	fmt.Println(fullName)
	fmt.Printf("Variable Type: %T \n", fullName)

	// Using contant Variables
	fmt.Println(lang)
	fmt.Printf("Variable Type: %T \n", lang)

	// Using Public Variables
	fmt.Println(Version)
	fmt.Printf("Variable Type: %T \n", Version)

}
