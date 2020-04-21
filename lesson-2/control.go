package main

import "fmt"

func main() {

	// разделение значений счетчика на четные и нечетные
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "четное")
		} else {
			fmt.Println(i, "нечетное")
		}
	}

	// разделение значений счетчика на четные и нечетные
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "четное")
			continue
		}
		fmt.Println(i, "нечетное")
	}

	// вариант с множественным условием
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("Число %v делится нацело на 2 \n", i)
		} else if i%3 == 0 {
			fmt.Printf("Число %v делится нацело на 3 \n", i)
		} else if i%5 == 0 {
			fmt.Printf("Число %v делится нацело на 5 \n", i)
		}
	}

	// switch как альтернатива if-else
	for i := 1; i <= 10; i++ {
		switch i {
		case 0:
			fmt.Println("Ноль")
		case 1:
			fmt.Println("Один")
		case 2:
			fmt.Println("Два")
		case 3:
			fmt.Println("Три")
		case 4:
			fmt.Println("Четыре")
		case 5:
			fmt.Println("Пять")
		default:
			fmt.Println("Неизвестный номер")
		}
	}
}
