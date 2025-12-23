package main

import (
	"fmt"
	"reflect"
)

type Person struct{
	Name string `required:"true" max:"10"`	
}

type Sample struct{
	Name string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
}

// reflect ini digunakan untuk mengentahui seluruh type struct
// numField untuk mengetahui jumlah fieldnya
func readField(value any)  {
	var valueType reflect.Type = reflect.TypeOf(value)  
	fmt.Println("type name: ", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		// var valueField reflect.StructField = valueType.Field(i)
		var valueSturctField reflect.StructField = valueType.Field(i)
		fmt.Println(valueSturctField.Name, "With type" , valueSturctField.Type)
		fmt.Println(valueSturctField.Tag.Get("required"))
		fmt.Println(valueSturctField.Tag.Get("max"))
		// fmt.Println(valueField);
	}
}

func isValid(value any) (result bool) {
	result = true
	var t = reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		var f = t.Field(i)
		if f.Tag.Get("required") == "true" {
			var data = reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}
	return result
}

// func isValidMax(value any) (result bool){
// 	result = true
// 	m := reflect.TypeOf(value)
// 	for i := 0; i < m.NumField(); i++ {
// 		t := m.Field(i)
// 		if t.Tag.Get("required") == "max" {
// 			data := reflect.ValueOf(value).Field(i).Interface()
// 			result = data != 10
// 			if result == false {
// 				return result
// 			}
// 		}
// 	}
// 	return result
// }

func main()  {
	readField(Sample{"Agus", "Inkopad"})
	readField(Person{"gusa",})

	sample := Sample{
		Name: "Balik",
		Address: "isi",
	}
	fmt.Println(isValid(sample))
}