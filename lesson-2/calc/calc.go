package main

type ErrorCalc struct {
	msg string
}

func (e *ErrorCalc) Error() string {
	return e.msg
}

func calculate(firstArg float64, operator string, secondArg float64) (float64, error) {
	switch operator {
	case "+":
		return firstArg + secondArg, nil
	case "-":
		return firstArg - secondArg, nil
	case "*":
		return firstArg * secondArg, nil
	case "/":
		if secondArg == 0 {
			return 0, &ErrorCalc{"Нельзя делить на ноль"}
		}
		return firstArg / secondArg, nil
	default:
		return 0, &ErrorCalc{"Некорректный оператор"}
	}
}
