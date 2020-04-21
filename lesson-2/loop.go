package main

import "fmt"

func main() {

	// цикл без условия (бесконечный)
	for {
		fmt.Println("Да, я все еще бесконечный")
	}

	// цикл с условием
	i := 0
	for i <= 10 {
		i++
		fmt.Println(i)
	}

	// классический нотация for-цикла
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// цикл по сущности, хранящей несколько элементов
	str := "Hello World"
	for _, ch := range str {
		fmt.Println("Код символа", ch)
	}

	// continue
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "четное")
			continue
		}
		fmt.Println(i, "нечетное")
	}

	// break
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "четное")
			break
		}
	}
}
