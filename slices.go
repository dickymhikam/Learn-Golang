package main

import (
	"fmt"
	"slices"
)

func main()  {
	Name := []string{"dicky", "m", "hikam"}
	Values := []int {1,2,3,4,5}

	fmt.Println(slices.Max(Values))
	fmt.Println(slices.Min(Values))
	fmt.Println(slices.Contains(Name, "dicky"))
	fmt.Println(slices.Index(Name, "hikam"))

}