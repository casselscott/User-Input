package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	//Create a new reader from standard input (os.Stdin)

	reader:= bufio.NewReader(os.Stdin)

	//Prompt the user to enter their city
	fmt.Print("What City do you Currently live in? ")


	//Read the input from the user

	city, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	//print the city entered by the user
	fmt.Println("You live in:", city)
	
}