package main

import "fmt"

func main () {

 var degrees int

 var idk string

fmt.Println("Enter the temperature here!")

fmt.Scanf("%d", &degrees)

fmt.Println("Press C for Celsious and F for Farenheight")

fmt.Scanf("%s", &idk)

if idk == "C" {
fmt.Println(degrees * 9 / 5 + 32)
} else if idk == "F" {
	fmt.Println(((degrees - 32) * 5) / 9)
	}

}
