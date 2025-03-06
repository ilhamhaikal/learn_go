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

func (us UserSlice) Len() int {
	return len(us)
}

func (us UserSlice) Less(i, j int) bool {
	return us[i].Age < us[j].Age
}

func (us UserSlice) Swap(i, j int) {
	us[i], us[j] = us[j], us[i]
}

func main() {
	User := []User{
		{"Ilham", 27},
		{"Haikal", 26},
		{"Budiman", 22},
		{"Marko", 20},
	}
	sort.Sort(UserSlice(User))

	fmt.Println(User)
}
