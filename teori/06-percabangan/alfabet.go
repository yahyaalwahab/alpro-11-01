package main

import "fmt"

func main() {
	var karakter string
	fmt.Scan(&karakter)

	if karakter >= "A" && <= "Z" || karakter >= "a"	&& <= "z" {
		fmt.Println("Alphabet")
	} else {
		fmt.Println("Bukan Alphabet")
	}
}