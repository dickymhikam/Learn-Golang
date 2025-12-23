package main

import (
	"encoding/csv"
	"os"
)

func main()  {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([] string{"Dicky", "M", "Hikam"})
	_ = writer.Write([] string{"Jaun", "bat", "ron"})
	_ = writer.Write([] string{"Selamet", "Jun", "ni"})

	writer.Flush()
	// flush itu mengirim semua perubahan 
}