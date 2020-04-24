package stack

var x []string

// Push добавит новый элемент в стек
func Push(str string) {
	x = append(x, str)
}

// Pop вернет последний добавленный в стек элемент
func Pop() string {
	numOfElements := len(x)
	// Когда стек будет пустым, он вернет пустую строку
	if numOfElements == 0 {
		return ""
	}
	popElem := x[numOfElements-1]
	x = x[:numOfElements-1]
	return popElem
}
