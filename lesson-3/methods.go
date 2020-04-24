package main

import "fmt"

type Person struct {
	Id   int
	Name string
}

// не изменит оригинальной структуры, для который вызван
func (p Person) UpdateName(name string) {
	p.Name = name
}

// изменяет оригинальную структуру
func (p *Person) SetName(name string) {
	p.Name = name
}

type Account struct {
	Id   int
	Name string
	Person
}

// func (p *Account) SetName(name string) {
// 	p.Name = name
// }

type MySlice []int

func (sl *MySlice) Add(val int) {
	*sl = append(*sl, val)
}

func (sl *MySlice) Count() int {
	return len(*sl)
}

func main() {
	// pers := &Person{1, "Vasily"}
	pers := new(Person)
	pers.SetName("Dmitry Shelamov")
	// (&pers).SetName("Vasily Romanov")
	// fmt.Printf("updated person: %#v\n", pers)

	var acc Account = Account{
		Id:   1,
		Name: "dshelamov",
		Person: Person{
			Id:   2,
			Name: "Dmitry Shelamov",
		},
	}

	acc.SetName("dmitry.shelamov")
	acc.Person.SetName("Test")

	fmt.Printf("%#v \n", acc)

	return

	sl := MySlice([]int{1, 2})
	sl.Add(5)
	fmt.Println(sl.Count(), sl)
}
