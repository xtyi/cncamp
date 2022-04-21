package main

import "fmt"

func main() {
	lst := []string{"I", "am", "stupid", "and", "weak"}
	for i, str := range lst {
		if str == "stupid" {
			lst[i] = "smart"
		}
		if str == "weak" {
			lst[i] = "strong"
		}
	}
	fmt.Println(lst)
}
