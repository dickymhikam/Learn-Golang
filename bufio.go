package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main()  {
	input := strings.NewReader("isi dari input ini panjang\nline baru lagi")
	reader := bufio.NewReader(input);

	for {
		line, _ , err := reader.ReadLine()
			if err == io.EOF {
				break
			}

			fmt.Println(string(line))
	}

	fmt.Println("========= writer =========")

	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("Hello world \n")
	_, _ = writer.WriteString("Welcome ")

	writer.Flush()
}