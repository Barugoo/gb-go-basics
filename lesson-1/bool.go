package main

import "fmt"

func main() {
	// создаются булевые переменные точно так же, как и любые другие
	var a, b bool

	a = true
	b = false

	fmt.Println("Оператор отрицания:", !a)
	fmt.Println("Оператор 'И':", a && b)
	fmt.Println("Оператор 'ИЛИ':", a || b)

	// результат сравнения любых переменных - булевый
	num1, num2 := 1, 2
	fmt.Println("Результат сравнения переменных:", num1 == num2)
}
