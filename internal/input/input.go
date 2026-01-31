package input

import (
	"bufio"
	"fmt"
)

//get user input
func GetUserInput(prompt string, r *bufio.Reader) {
	fmt.Print(prompt)

	input, err := r.ReadString('\n')

	if err != nil {
		panic(err)
	}

	switch input {
	case "A":
		fmt.Println("You picked A")
	case "B":
		fmt.Println("You picked B")
	case "C":
		fmt.Println("You picked C")
	case "D":
		fmt.Println("You picked D")
	default:
		fmt.Println("I don't know what that is, sorry...")
		GetUserInput(prompt, r)
	}
}