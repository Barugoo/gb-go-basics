package main

import "fmt"

type name string
type phones []int

func (p *phones) Print() {
	for _, phone := range *p {
		fmt.Println(phone)
	}
}

func main() {
	phones := phones{892894439, 83928942}
	phoneBook := map[name]phones
}
