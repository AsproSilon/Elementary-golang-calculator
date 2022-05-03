package main

import "fmt"

func main() {
	var operator, operator1, next, enter string
	var number1, number2, number3 int

	fmt.Println("Enter number 1")
	fmt.Scanln(&number1)

	fmt.Println("Enter number 2")
	fmt.Scanln(&number2)

	fmt.Println("Choose operator: +,-,/,*")
	fmt.Scanln(&operator)

	output := 0

	switch operator {
	case "+":
		output = number1 + number2

	case "-":
		output = number1 - number2

	case "/":
		output = number1 % number2

	case "*":
		output = number1 * number2

	}
	fmt.Println(output)

	//loop for infinite calculator
	for enter != "Enter" {
		fmt.Println(" Choose next action ")
		fmt.Println("If you want continue, write C, or if you want stop, write Enter")
		fmt.Scanln(&next)

		// if you want to stop the calculator , write Enter and the program will stop
		if next == "Enter" {
			break
			// if you want continue , write C and calculator ask you for enter operator
		} else if next == "C" {
			fmt.Println("Enter number 3")
			fmt.Scanln(&number3)

			fmt.Println("Choose operator1: +,-,/,*")
			fmt.Scanln(&operator1)

			answer := 0

			switch operator1 {
			case "+":
				answer = output + number3
			case "-":
				answer = output - number3
			case "/":
				answer = output % number3
			case "*":
				answer = output * number3

			}
			fmt.Println(answer)
		}
	}

}
