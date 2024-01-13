package main

import "fmt"

const LoginToken string = "1eff4rwfss#f"

// this first letter as capital has some significance as
// it is public for all the files in same program and can be access any file in
// in same folder as of now

func main() {
	var username string= "Prasad Patil";

	fmt.Println(username) 
	fmt.Printf("Variable is of type: %T \n", username)

	// below is how we go for boolean data type

	var isLoggedIn bool = true
	fmt.Printf("Logged? %T \n", isLoggedIn)

	/*
		there are many data types:
		-int
		-float32
		-float64
	*/

	fmt.Println(LoginToken)
}
