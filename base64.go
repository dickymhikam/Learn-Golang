package main

import (
	"encoding/base64"
	"fmt"
)

func main()  {
	value := "Dicky M Hikam"

	encode := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encode)

	var decode, err = base64.StdEncoding.DecodeString(encode)
	if err != nil {
		fmt.Println("error", err.Error())
	}else{
		fmt.Println(string(decode))
	}
}