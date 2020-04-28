package main

import (
	"errors"
	"fmt"
)

func main() {
	var firstArg, secondArg float64
	var operator string

	fmt.Println("Введите первый аргумент:")
	fmt.Scanln(&firstArg)

	fmt.Println("Введите оператор")
	fmt.Scanln(&operator)

	fmt.Println("Введите второй аргумент:")
	fmt.Scanln(&secondArg)

	result, err := calculate(firstArg, operator, secondArg)
	if err != nil && errors.As(err, ErrorCalc{}) {
		fmt.Printf("Произошла ошибка: %s", err.Error())
		return
	}
	fmt.Println("Результат:", result)
}
