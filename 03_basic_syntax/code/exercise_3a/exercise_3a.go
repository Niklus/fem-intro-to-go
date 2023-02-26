package main

import "fmt"

func main () {
	
	motto := "Golang is amaizing!"

	// Print odd letters
	for index, element := range motto {
		if index % 2 == 1 {
			fmt.Println(string(element))
		}
	}
	
  // Print even letters
	for _, element := range motto {
		fmt.Println(string(element))	
	}
}