package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

// Len = untuk menghitung panjang dari value itu sendiri
func (s UserSlice) Len() int {
	return len(s)
}

// Less = untuk mengurutkan / sorting valuenya
func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	// var temp = s[i]
	// temp = s[j]
	// s[j] = temp
	s[i], s[j] = s[j], s[i]
	fmt.Println("swap", i,j)
}

func main() {
	var users = []User{
		{"Budi", 22},
		{"Alex", 21},
		{"Rudi", 23},
		{"audi", 25},
		{"rudi", 24},
	}
	sort.Sort(UserSlice(users))
	fmt.Println(users)
}