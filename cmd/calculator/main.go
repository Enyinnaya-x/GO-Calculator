package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"github.com/Enyinnaya-x/GO-Calculator/internal/input"
)

func main() {

	//welcome user
	fmt.Println("👋 Hello there, welcome to the Go calculator")

	// delay for 2 seconds
	time.Sleep(2 * time.Second)

	//ask what operation user wants to carry out from options
	fmt.Println("😁 What operation would you like to carry out?")

	//define reader
	reader := bufio.NewReader(os.Stdin)

	//get user input
	input.GetUserInput("A - Addition, B - Subtraction, C - Multiplication, D - Division : ", reader)
}
