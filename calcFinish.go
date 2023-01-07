package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func int_roma(num int) string {

	if num < 1 {
		return "Вывод ошибки, так как в римской системе нет отрицательных чисел."
	}

	r := [][]string{
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}}
	n := []int{100, 10, 1}
	result := ""

	for k, v := range n {
		result += r[k][num/v]
		num = num % v
	}

	return result
}

func search(key string, arr []string) int {

	for index, value := range arr {
		if value == key {
			return index
		}
	}

	return -1
}

func calculeted(a int, b int, operation string, no_roma bool) string {

	if operation == "+" {
		if no_roma {
			return int_roma(a + b)
		} else {
			return strconv.Itoa(a + b)
		}
	}

	if operation == "-" {
		if no_roma {
			return int_roma(a - b)
		} else {
			return strconv.Itoa(a - b)
		}
	}

	if operation == "*" {
		if no_roma {
			return int_roma(a * b)
		} else {
			return strconv.Itoa(a * b)
		}
	}

	if operation == "/" {
		if no_roma {
			return int_roma(a / b)
		} else {
			return strconv.Itoa(a / b)
		}
	}

	return fmt.Sprint("Знак: ", operation, " не верный, введите (+, -, /, *) ")
}

func main() {
	no_roma := false
	true_numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	true_numbers_rimpal := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите пример (1 + 1 или I + I): ")
	equation, _ := reader.ReadString('\n')
	equation_list := strings.Fields(equation)

	if len(equation_list) != 3 {
		fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		return
	}

	index_a, index_b := search(equation_list[0], true_numbers), search(equation_list[2], true_numbers)
	if index_a == -1 || index_b == -1 {
		no_roma = true
		index_a, index_b = search(equation_list[0], true_numbers_rimpal), search(equation_list[2], true_numbers_rimpal)
		if index_a == -1 || index_b == -1 {
			fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
			return
		}
	}

	fmt.Println(calculeted(index_a+1, index_b+1, equation_list[1], no_roma))
}
