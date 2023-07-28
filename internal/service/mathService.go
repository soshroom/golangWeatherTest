package service

import (
	"errors"
	"strconv"
	"strings"
)

type mathService struct {
}

func (m mathService) ProcessExpression(expression string) (int, error) {
	return calculateExpression(expression)
}

func NewMathService() *mathService {
	return &mathService{}
}

func calculateExpression(expression string) (int, error) {
	tokens := strings.Split(expression, "")
	stack := make([]int, 0)
	operator := '+'
	num := 0

	for i := 0; i < len(tokens); i++ {
		if tokens[i] == " " {
			continue
		}

		if isDigit(tokens[i]) {
			num = num*10 + parseInt(tokens[i])
		}

		if !isDigit(tokens[i]) || i == len(tokens)-1 {
			switch operator {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			default:
				return 0, errors.New("некорректный оператор")
			}
			operator = rune(tokens[i][0])
			num = 0
		}
	}

	result := 0
	for _, val := range stack {
		result += val
	}

	return result, nil
}

func isDigit(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func parseInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
