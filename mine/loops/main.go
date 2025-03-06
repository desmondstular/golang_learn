package main

import "fmt"

func main() {
	
	for i, letter := range "dude" {
		fmt.Println(i, string(letter))
	}

	aMap := map[string]int {
		"tatum": 27,
		"desmond": 27,
		"eva": 6,
	}

	for k, t := range aMap {
		fmt.Printf("%s : %v\n", k, t)
	}

	var a, b int = 4, 10

	result, error := test(a, b)

	fmt.Println(result, error)
}

func test(a int, b int) (int, error) {
	return (a + b), nil
}