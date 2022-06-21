package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func subtract(x int, y int) int {
	return x - y
}

func multiply(x int, y int) int {
	return x * y
}

func divide(x int, y int) (float64, error) {
	if y == 0 || x == 0 {
		return 0, fmt.Errorf("You cannot division by 0")
	}
	return float64(x) / float64(y), nil
}

func chooseOperation() {
	selectOperation := "Select operation."
	fmt.Printf("Please %s\n", selectOperation)
	// Two figures structs is empty struct
	operation := map[string]struct{}{"1 .add": {}, "2 .subtract": {}, "3 .multiply": {}, "4 .divide": {}}
	for k := range operation {
		fmt.Println(k)
	}
}

func getTwoValues() (int, int, error) {
	var enterFirstValue int
	var enterSecondValue int
	var err error

	_, err = fmt.Scan(&enterFirstValue)
	if err != nil {
		return 0, 0, err
	}
	_, err = fmt.Scan(&enterSecondValue)
	if err != nil {
		return 0, 0, err
	}
	return enterFirstValue, enterSecondValue, err
}

func processingCalculation() {
processing:
	for {
		var result int
		var resultDivision float64
		choice := ""
		fmt.Println("Please enter your choice(1/2/3/4): ")
		fmt.Scan(&choice)

		switch choice {
		case "1":
			fmt.Println("Enter two values to add:")
			firstValue, secondValue, err := getTwoValues()
			if err != nil {
				fmt.Println(err)
				continue
			}
			result = add(firstValue, secondValue)
			fmt.Printf("Result %d\n", result)
		case "2":
			fmt.Println("Enter two values to subtract:")
			firstValue, secondValue, err := getTwoValues()
			if err != nil {
				fmt.Println(err)
				continue
			}
			result = subtract(firstValue, secondValue)
			fmt.Printf("Result %d\n", result)
		case "3":
			fmt.Println("Enter two values to multiply:")
			firstValue, secondValue, err := getTwoValues()
			if err != nil {
				fmt.Println(err)
				continue
			}
			result = multiply(firstValue, secondValue)
			fmt.Printf("Result %d\n", result)
		case "4":
			fmt.Println("Enter two values to divide:")
			firstValue, secondValue, err := getTwoValues()
			if err != nil {
				fmt.Println(err)
				continue
			}
			resultDivision, err = divide(firstValue, secondValue)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Result %f\n", resultDivision)
		default:
			fmt.Println("The program has been terminated.")
			break processing
		}
	}
}

func main() {
	chooseOperation()
	processingCalculation()
}
