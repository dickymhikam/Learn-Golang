package main

import (
	"bufio"
	"io"
	"os"
)

func createNewFile (name string, message string) error  {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY,0666)
	if err != nil {
		return  err
	}
	defer file.Close()
	file.WriteString(message)
	return nil 
}

func addToFile (name string, message string) error  {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND,0666)
	if err != nil {
		return  err
	}
	defer file.Close()
	file.WriteString(message)
	return nil 
}

func readFile (name string) (string, error)  {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	// untuk memastikan file ditutup 
	defer file.Close()
	
	reader := bufio.NewReader(file)
	// pada var message ini untuk menampung isi filenya 
	var message string
	for {
		line, _ ,err := reader.ReadLine()
		// pengecekan kalo misalnya datanya habis maka diberhentikan 
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil
}

func main()  {
	// createNewFile("sample log", "this is sample message")
	// result, _ := readFile("sample log")
	// fmt.Println(result)
	addToFile("sample log", "\nThis New Message")
}