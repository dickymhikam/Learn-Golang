package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation Error")
	NotFoundError = errors.New("Not Found Error")
)

func GetById(id string) error {
	if id == "" {
		return  ValidationError
	}

	if id != "Dicky" {
		return NotFoundError
	}

	return nil
}

// menggunakan is untuk mendefinisasikan sebuah pengecekan error 

func main()  {
	err := GetById("Dicky")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation Error:", ValidationError.Error())
		} else if errors.Is(err, NotFoundError){
			fmt.Println("Not Found Error:", NotFoundError.Error())
		}else{
			fmt.Println("unknown error")
		}

	}else{
		fmt.Println("Sukses")
	}
}