package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func isNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func print(result string) {
	if len(result) > 40 {
		result = result[0:40] + "..."
		fmt.Println("Введите выражение: ")
	} else {
		result = strconv.Quote(result)
	}
	fmt.Println(result)
	fmt.Println("Введите выражение: ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите выражение:")
	for scanner.Scan() {

		originalInput := scanner.Text()
		input := originalInput

		if strings.Contains(input, "\" - \"") {
			reg := regexp.MustCompile(`^"([^"]+)"`)
			match := reg.FindString(input)
			if match != "" {
				input = strings.ReplaceAll(input, match, strings.ReplaceAll(match, " ", ""))
			}
		}

		parts := strings.Split(input, " ")
		if len(parts) != 3 {
			panic("Неверный формат ввода !")
		}

		aStr, operator, bStr := parts[0], parts[1], parts[2]

		if strings.HasPrefix(aStr, `"`) && strings.HasSuffix(aStr, `"`) {
			if (len(aStr) > 12) || (len(bStr) > 12) {
				panic("Длина строки не должна превышать 10 символов !")
			}

			switch operator {
			case "+":
				if (strings.HasPrefix(aStr, `"`) && strings.HasSuffix(aStr, `"`)) && (strings.HasPrefix(bStr, `"`) && strings.HasSuffix(bStr, `"`)) {
					result := strings.Trim(aStr, "\"") + strings.Trim(bStr, "\"")
					print(result)
				} else {
					panic("Калькулятор умеет рассчитывать сумму только между строками !")
				}

			case "-":
				if (strings.HasPrefix(aStr, `"`) && strings.HasSuffix(aStr, `"`)) && (strings.HasPrefix(bStr, `"`) && strings.HasSuffix(bStr, `"`)) {

					if originalInput != input {

						result := strings.Replace(strings.Trim(aStr, "\""), strings.Trim(bStr, "\""), "", -1)
						print(result + " ")

					} else {
						result := strings.Replace(strings.Trim(aStr, "\""), strings.Trim(bStr, "\""), "", -1)
						print(result)
					}

				} else {
					panic("Калькулятор умеет рассчитывать разность только между строками !")
				}

			case "*":

				if isNumber(bStr) {
					bInt, _ := strconv.Atoi(bStr)
					if bInt >= 1 && bInt <= 10 {
						result := strings.Repeat(strings.Trim(aStr, "\""), bInt)
						print(result)
					} else {
						panic("Диапозон чисел должен быть от 1 до 10!")
					}
				} else {
					panic("Калькулятор умеет рассчитывать произведение между строкой и целым числом!")
				}

			case "/":

				if isNumber(bStr) {
					bInt, _ := strconv.Atoi(bStr)
					if bInt >= 1 && bInt <= 10 {
						a := strings.Trim(aStr, "\"")
						result := a[0 : len(a)/bInt]
						print(result)
					} else {
						panic("Диапозон чисел должен быть от 1 до 10!")
					}
				} else {
					panic("Калькулятор умеет расчитывать частное только между строкой и целым числом !")
				}
			default:
				panic("Неверный оператор !")
			}
		} else {
			panic("Первая строка не должна быть числом !")
		}
	}
}
