package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main22() {
	names := []string{"abc", "def", "ghi"}
	names = append(names, "abcdefg")
	fmt.Println(names)

	nico := map[string]string{"abc": "12", "abcd": "13"}
	for key, value := range nico {
		fmt.Println(key, value)
		fmt.Println(nico[key])
	}
	if val := nico["abcjsdalfj"]; val == "" {
		fmt.Println("11111111")
	} else {
		nico["abcjsdalfj"] = "1234567890"
		fmt.Println(nico["abcjsdalfj"])
	}
	fmt.Println(nico["abc"])
	fmt.Println(nico["abcdefdfd"])
	fmt.Println(nico)
	a := 2
	b := &a
	fmt.Println(a)
	fmt.Println(*b)

	favFood := []string{"kimchi", "ramen"}
	mer2 := []person{}
	mer := person{name: "merkuree", age: 33, favFood: favFood}
	mer2 = append(mer2, mer)
	mer2 = append(mer2, person{"kimnana", 28, favFood})
	fmt.Println(mer)
	fmt.Println(mer.name)
	fmt.Println(mer2)
	fmt.Println(mer2[1])
}
