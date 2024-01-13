package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Can you give rating for out pizza ")

	// comma ok syntax below to get input
	// its also called as comma err syntax

	input, _ := reader.ReadString('\n')

	/*
      inplace of underscore mostly some variable is used
	  which is used to handle errors while taking input 
	  we will look into it in future 
	  
	*/

	fmt.Println("Thanks for Rating: ", input)

}
