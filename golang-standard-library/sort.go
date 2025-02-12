package main

import (
	"fmt"
	"sort"
)

type User struct {
		Name string
		Age  int
}

type Userslice []User

func (s Userslice) Len() int {
	return len(s)
}

func (s Userslice) Less(i, j int) bool {
	return s[i].Age < s[j].Age

}

func (s Userslice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	user := []User{
		{"Jonathan", 20},
		{"Budi", 30},
		{"Bata", 10},
	}

	sort.Sort(Userslice(user))
	fmt.Println(user)
}
