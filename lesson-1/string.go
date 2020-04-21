package main

import "fmt"

func main() {
	var str1, str2 string

	str1 = "Hello"

	// мы получить длину строки в байтах с помощью операции len()
	fmt.Println(len(str1))

	str1 = "Привет"
	// длина строки не всегда равна количеству символов
	fmt.Println(len(str1))

	// также мы можем получить отдельный байт в строке используя индексную нотацию
	byteElem := str1[2]
	fmt.Println(byteElem)

	str2 = " Мир"

	// строки можно конкатенировать
	fmt.Println(str1 + str2)

	// поддержка управляющих последовательности
	str1 = "Привет\nМир"
	fmt.Println(str1)
}
