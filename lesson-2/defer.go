package main

import "fmt"

func main() {
	// Будет вызвана последней
	defer fmt.Println("Первая строка")

	// Будет вызвана третьей
	defer fmt.Println("Вторая строка")

	// Будет вызвана первой
	fmt.Println("Третья строка")

	// Будет вызвана второй
	fmt.Println("Четвертая строка")
}
