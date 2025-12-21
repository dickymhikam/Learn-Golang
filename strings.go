package main

import (
	"fmt"
	"strings"
)

// package strings adalah package yang berisikan function function untuk memanipulasi type data string
func main()  {
	fmt.Println(strings.Contains("Dicky M","M")) // untuk mengecek apakah pada value pertama terdapat nama tersebuts
	fmt.Println(strings.Split("Dicky M"," ")) // untuk memotong string bedasarkan separator
	fmt.Println(strings.ToLower("Dicky M")) //untuk mengubah menjadi huruf kecil semua
	fmt.Println(strings.ToUpper("Dicky M")) //untuk mengubah menjadi huruf besar semua
	fmt.Println(strings.Trim("      Dicky M      ", " ")) //memotong cutset di awal dan di akhir string
	fmt.Println(strings.ReplaceAll("Dicky M  Dicky Hikam","Dicky", "Nudi")) //mengganti nama dicky menjadi nudi
}