package input

import (
	"bufio"
	"fmt"
	"strings"
	"github.com/Enyinnaya-x/GO-Calculator/internal/operations"
)

// get user input
func GetUserInput(prompt string, r *bufio.Reader) {
	fmt.Print(prompt)

	input, err := r.ReadString('\n')

	if err != nil {
		panic(err)
	}

	//set user selection to uppercase
	input = strings.ToUpper(input)

	//trim white space from user input
	input = strings.TrimSpace(input)

	switch input {
	case "A":
		operations.ArithmeticOperation("+")
	case "B":
		operations.ArithmeticOperation("-")
	case "C":
		operations.ArithmeticOperation("*")
	case "D":
		fmt.Println("You picked D")
		operations.ArithmeticOperation("/")
	default:
		fmt.Println("I don't know what that is, sorry...")
		GetUserInput(prompt, r)
	}
}
