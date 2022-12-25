package example

import "fmt"

func thisIsSilly() { // want "should not contain silly"
	fmt.Println("Hi")
}

func normal() {
	fmt.Println("Hi normal")
}
