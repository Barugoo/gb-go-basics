package main

import "fmt"

func main() {
	// создание
	var slice []int // len=0, cap=0

	// обращение к элементам
	someInt := slice[0]

	// ошибка при выполнении
	// panic: runtime error: index out of range
	// someOtherInt := buf2[1]

	fmt.Println(someInt)

	// добавление элементов
	var buf []int            // len=0, cap=0
	buf = append(buf, 9, 10) // len=2, cap=2
	buf = append(buf, 12)    // len=3, cap=4

	// добавление друго слайса
	otherBuf := make([]int, 3)     // [0,0,0]
	buf = append(buf, otherBuf...) // len=6, cap=8

	fmt.Println(buf, otherBuf)

	// просмотр информации о слайсе
	var bufLen, bufCap int = len(buf), cap(buf)

	fmt.Println(bufLen, bufCap)

	vv := [][]int{
		[]int{1},
		[]int{3, 4},
	}

	fmt.Println(vv)
}
