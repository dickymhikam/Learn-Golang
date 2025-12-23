package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

	func main()  {
		csvString := "Dicky, M, Hikam \n" +
					"Joko, Anwar, Jo \n" +
					"Selamet, Udin"

		reader := csv.NewReader(strings.NewReader(csvString))
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			fmt.Println(record)
		} 
	}