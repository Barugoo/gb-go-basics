package main

import (
	"fmt"
)

func main() {
	// значение по умолчанию для целочисленных переменных
	var defaultInt int
	fmt.Println("int type default value:", defaultInt)

	// значение по умолчанию для числовых переменных с плавующей точкой
	var defaultFloat float32
	fmt.Println("float type default value:", defaultFloat)

	// значение по умолчанию для булевых
	var defaultBool bool
	fmt.Println("bool type default value:", defaultBool)

	// значение по умолчанию для строк
	var defaultString string
	fmt.Println("string type default value:", defaultString)

	// значение по умолчанию для указателей
	var defaultPointer *int
	fmt.Println("pointer type default value:", defaultPointer)
}
