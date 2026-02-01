package operations

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)


	//create a slice
	var numbers = []float64 {}

func GetNumbers(prompt string, r *bufio.Reader) []float64 {
	//ask user question
	fmt.Print(prompt)

	//read the number the user initially gives
	num, err := r.ReadString('\n')

	//classic error handling
	if err != nil {
		panic(err)
	}

	//remove white space
	trimmedNum := strings.TrimSpace(num)

	//convert the input from the user from "string" to "float64" to allow arithmetic operations
	n, err := strconv.ParseFloat(trimmedNum, 64)

	//classic error handling incase something goes wrong with conversion
	if err != nil {
		panic(err)
	}

	//store the values gotten from the user in the earlier created slice
	numbers = append(numbers, n)

	//Confirm if user wants to use other numbers in this operation
	fmt.Print("Any other numbers? yes/no ")

	//get response "yes" or "no"
	reply, err := r.ReadString('\n')

	//classic error handling
	if err != nil {
		panic(err)
	}

	trimmedReply := strings.TrimSpace(reply)

	//take action based on reply
	reader := bufio.NewReader(os.Stdin)


	switch trimmedReply {

	case "yes":
		GetNumbers("What other number? ", reader)

	case "no":
		fmt.Println("🧮 Calculating...")
		time.Sleep(1 * time.Second)

	default:
		fmt.Println("❌ That is not an option...")
		GetNumbers("Give me a number ", reader)
	}

	//return slice of all numbers
	return numbers
}


func ArithmeticOperation(operator string) float64 {
	reader := bufio.NewReader(os.Stdin)

	numbers := GetNumbers("Give me a number ", reader)

		switch operator {
		case "+":
			answer := 0.0
			for index := 0; index < len(numbers); index++ {
				answer += numbers[index]
			}
			fmt.Println("✅ Your answer is ", answer)

			
			case "-":
				answer := 0.0
				for index := 0; index < len(numbers) - 1; index++ {
					answer = numbers[index] - numbers[index+1]
				}
				fmt.Println("✅ Your answer is ", answer)

				
			case "*":
				answer := 1.0
				for index := 0; index < len(numbers); index++ {
					answer *= numbers[index]
				}
				fmt.Println("✅ Your answer is ", answer)


			case "/":
				answer := 0.0
				for index := 0; index < len(numbers); index++ {
					if numbers[index] == 0 {
						fmt.Println("❌ Division by zero is not allowed.")
						return 0
					}
					if index == 0 {
						answer = numbers[index]
					} else {
						answer /= numbers[index]
					}
				}
				fmt.Println("✅ Your answer is ", answer)
				default:
					fmt.Println("❌ Unknown operator.")
		}
	return 0
}
		