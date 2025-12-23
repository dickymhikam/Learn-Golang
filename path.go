package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main()  {
	// Dir = Direktori
	fmt.Println(path.Dir("hello/world.go"))
	// base = lokasi filenya
	fmt.Println(path.Base("hello/world.go"))
	// ext = extensionnya atau bahasanya
	fmt.Println(path.Ext("hello/world.go"))
	// join = untuk menggabugkan atau berhubungan 
	fmt.Println(path.Join("hello", "world", "main.go"))

	fmt.Println("=========== filepath ===========")

	fmt.Println(filepath.Dir("hello/world.go"))
	fmt.Println(filepath.Base("hello/world.go"))
	fmt.Println(filepath.Ext("hello/world.go"))
	// is absolute = ketika path dimulai dari awal seperti dari drive mana C / D
	// sedangkan isLocal kebalikannya
	fmt.Println(filepath.IsAbs("hello/world.go"))
	fmt.Println(filepath.IsLocal("hello/world.go"))
	fmt.Println(filepath.Join("hello", "world", "main.go"))

}