package main

import (
	"fmt"
	"scrapper_echo/accounts"
	"scrapper_echo/mydict"
)

func main_test() {
	dictionary := mydict.Dictionary{"first": "First Word"}
	hello := "hello"
	greeting := "Greeting"

	if err := dictionary.Add(hello, greeting); err != nil {
		fmt.Println(err)
	}
	fmt.Println("ADD")
	if err := dictionary.Add(hello, greeting); err != nil {
		fmt.Println(err)
	}

	if val, err2 := dictionary.Search("llo"); err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(val)
	}

	dictionary.Update(hello, "Say Hello")
	dictionary.Update("2hello", "Say Hello")
	// dictionary.Delete("first")
	// dictionary.Delete("2hello")
	fmt.Println("key list : ")
	for key, value := range dictionary {
		fmt.Println(key)
		fmt.Println(value)
	}
}

func main2222() {
	merkuree := accounts.NewAccount("merkuree")
	merkuree.Deposit(10)
	fmt.Println(merkuree.Balance())
	nico := accounts.NewAccount("nico")
	nico.Deposit(102)
	fmt.Println(nico)
	nico2 := accounts.NewAccount("merkuree")
	nico2.Deposit(10)
	err := nico2.ChangeWithdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(nico2)
	fmt.Println(merkuree)

	// bank := accounts.NewBank("bbb", "ccc", "ddd", 1)
	// fmt.Println("bank", bank)
	// bank = accounts.NewBank("bbb", "ccc", "ddd", 2)
	// fmt.Println("bank", bank)
}
