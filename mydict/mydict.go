package mydict

import (
	"errors"
	"fmt"
	"scrapper_echo/accounts"
)

//Dictionay type
type Dictionary map[string]string

var flag = accounts.NewBoolean(true)

//Not Found
var (
	errNoFound    = errors.New("word Not Found")
	errWordExists = errors.New("That word is already exists")
	errCantUpdate = errors.New("Can't update non-exists")
)

//Search a word to the Dictionary
func (d Dictionary) Search(word string) (string, error) {
	if value, exists := d[word]; exists {
		return value, nil
	}
	return "", errNoFound

}

//Add a word to the Dictionary
func (d Dictionary) Add(key string, word string) error {
	_, err := d.Search(key)
	switch err {
	case errNoFound:
		d[key] = word
	case nil:
		return errWordExists
	}

	return nil
}

//Update a word to the Dictionary
func (d Dictionary) Update(key, word string) {
	if _, err := d.Search(key); flag.SystemFlag() {
		if err == errNoFound {
			d.Add(key, word)
		} else {
			d[key] = word
		}
	} else {
		switch err {
		case errNoFound:
			fmt.Println(key, errCantUpdate)
		case nil:
			d[key] = word
		}
	}
}

//Delete a word to the Dictionary
func (d Dictionary) Delete(key string) {
	_, err := d.Search(key)
	switch err {
	case nil:
		delete(d, key)
	case errNoFound:
		fmt.Println(key, errNoFound)
	}
}
