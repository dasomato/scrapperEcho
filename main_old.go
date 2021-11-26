package main

import (
	"fmt"
	"scrapper_echo/modules/hello"
	"strings"
)

func multiply(a int, b int) int {
	return a * b
}

func deferfunc(name string) string {
	return "ENDENDEND" + name
}
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(name ...string) {
	defer fmt.Println(deferfunc("abcEnd"))
	for index, nm := range name {
		fmt.Println(index, nm)
		fmt.Println(name[index])
	}
	fmt.Println(name)
}

func lenAndUpper2(name string) (legnth int, uppercase string) {
	legnth = 3
	uppercase = "nico"
	return
}

func forTest(start int, end int) {
	for i := start; i <= end; i++ {
		for x := 1; x <= 9; x++ {
			fmt.Printf("%d x %d = %d\n", i, x, i*x)
		}
	}
}

func canIDrink(plusAgeFlag bool, age int) bool {
	if plusAge := age; plusAgeFlag {
		plusAge += 2
		if plusAge < 18 {
			return false
		} else {
			return true
		}
	} else {
		if plusAge < 18 {
			return false
		} else {
			return true
		}
	}
	return true
}

func main_old() {
	fmt.Println(canIDrink(false, 16))
	_, name := lenAndUpper("lee kilhun")
	var world2 string = "asdfjlasjdflk"
	const world string = "DDDDD"
	forTest(2, 9)
	fmt.Println(world2)
	fmt.Println(world)
	length, name2 := lenAndUpper2("lee")
	fmt.Println(length, name2)
	fmt.Printf("name : %s\n", name)
	repeatMe("nico", "abc")
	fmt.Printf("multiply : %d\n", multiply(2, 2))
	fmt.Println("Hello World")
	hello.SayHello()
	hello.SayBye()
	hello.SayGoodBye()
}
