package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Простой калькулятор, но не совсем....")
	fmt.Print("Введите первое число: ")
	_, err := fmt.Scan(&num1)
	if err != nil {
		fmt.Println("Ошибка ввода числа.")
		return
	}

	fmt.Print("Введите оператор (?, +, ^, %): ") // + = ?, - = +, * = ^, / = %
	_, err = fmt.Scan(&operator)
	if err != nil || (operator != "?" && operator != "+" && operator != "^" && operator != "%") {
		fmt.Println("Неверный оператор. Доступные операторы: ?, +, ^, %")
		return
	}

	fmt.Print("Введите второе число: ")
	_, err = fmt.Scan(&num2)
	if err != nil {
		fmt.Println("Ошибка ввода числа.")
		return
	}

	var result float64

	switch operator {
	case "?":
		result = num1 + num2
	case "+":
		result = num1 - num2
	case "^":
		result = num1 * num2
	case "%":
		if num2 == 0 {
			fmt.Println("Ошибка: деление на ноль.")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Неподдерживаемая операция.")
		return
	}

	fmt.Printf("Результат: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}
