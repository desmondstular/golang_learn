package main

import "fmt"


func main() {
	// int types = int int64 int32 int16 int8
	// unsigned int types = uint64 uint32 uint16 uint
	var num int = 10

	//float types: float64 float32

	fmt.Println(num)

	// Strings can use "" or ``. If you use ``, it will look exactly how you type when you print to console.
	var myString string = `My name is
Desmond`

	fmt.Println(myString)

	// To get the number of characters, can use len(string) function.
	// It is different as it returns the number of bytes
	// So if you expect fancy chars like gamma symbol (returns 2 bytes length per character)
	// use a package like "unicode/utf8" and use the function utf8.RuneCountInString() function but it is for runes.


	// Default values for ints, floats, runes is 0 upon intialization
	// 					for strings is ""
	//					for boolean is false

	// const declares a constant value and cannot be changed after initialization
	// consts cannot be declared empty like a var
}