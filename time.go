package main

import (
	"fmt"
	"time"
)

func main()  {
	var timeNow time.Time = time.Now()
	fmt.Println(timeNow.Local())

	var utc time.Time = time.Date(2025, time.December, 12, 7, 0, 0, 0, time.UTC)
	fmt.Println(utc)

	format := "2006-10-10 15:04:05"
	value := "2020-12-12 10:00:00"
	valueTime ,err := time.Parse(format, value)
	if err != nil {
		fmt.Println("error",err.Error())
	}else{
		fmt.Println(valueTime)
	}
	
}