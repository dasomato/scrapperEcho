package hello

import "fmt"

func SayBye2() {
	fmt.Println("SayBye")
}

func SayGoodBye() {
	greeting := true
	fmt.Println(greeting) //function
	greeting = false
	fmt.Println(greeting)
}
