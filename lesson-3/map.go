package main

import "fmt"

func main() {
	// инициализация при создании
	var user map[string]string

	short := map[string]string{
		"name":     "Dmitry",
		"lastName": "Shelamov",
	}

	//  сразу с нужной ёмкостью
	profile := make(map[string]string, 10)

	// количество элементов
	mapLength := len(user)

	fmt.Printf("%d %+v\n", mapLength, user)

	// если ключа нет - вернёт значение по умолчанию для типа
	kk := "middleName"
	mName := user[kk]
	fmt.Println("mName:", mName)

	// проверка на существование ключа
	mName, mNameExist := user["middleName"]
	fmt.Println("mName:", mName, "mNameExist:", mNameExist)

	// пустая переменная - только проверяем что ключ есть
	_, mNameExist2 := user["middleName"]
	fmt.Println("mNameExist2", mNameExist2)

	// удаление ключа
	delete(user, "lastName")
	fmt.Printf("%#v\n", user)

	// итеррация по мапе
	for key, value := range user {
		fmt.Println(key, value)
	}
}
