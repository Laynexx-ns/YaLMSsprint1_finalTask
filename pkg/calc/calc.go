package calc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func isNum(r rune) bool {
	return unicode.IsDigit(r)
}

var priority = map[rune]int{
	'+': 1,
	'-': 1,
	'/': 2,
	'*': 2,
}

func isOperator(ch rune) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/'
}

func applyOperator(a, b float64, operator rune) float64 {
	switch operator {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':

		return a * b
	case '/':

		return a / b
	}
	return 0
}

func Parse(expression string) ([]string, error) {
	var result []string
	var operators []rune
	var num strings.Builder

	for _, ch := range expression {
		if isNum(ch) || ch == '.' {
			num.WriteRune(ch)
		} else {
			if num.Len() > 0 {
				result = append(result, num.String())
				num.Reset()
			}
			if isOperator(ch) {
				for len(operators) > 0 && priority[operators[len(operators)-1]] >= priority[ch] {
					result = append(result, string(operators[len(operators)-1]))
					operators = operators[:len(operators)-1]
				}
				operators = append(operators, ch)
			} else if ch == '(' {
				operators = append(operators, ch)
			} else if ch == ')' {
				for len(operators) > 0 && operators[len(operators)-1] != '(' {
					result = append(result, string(operators[len(operators)-1]))
					operators = operators[:len(operators)-1]
				}
				if len(operators) == 0 {
					return nil, fmt.Errorf("wrong((")
				}
				operators = operators[:len(operators)-1]
			}
		}
	}
	if num.Len() > 0 {
		result = append(result, num.String())
	}
	for len(operators) > 0 {
		result = append(result, string(operators[len(operators)-1]))
		operators = operators[:len(operators)-1]
	}

	return result, nil
}

func evaluate(parsedExpression []string) (float64, error) {
	var stack []float64

	for _, token := range parsedExpression {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			stack = append(stack, num)
		} else if len(stack) >= 2 {
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			result := applyOperator(a, b, rune(token[0]))
			stack = append(stack, result)
		} else {
			return 0, fmt.Errorf("calculator died")
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("it's all your fault")
	}

	return stack[0], nil

}

func Calc(expression string) (float64, error) {
	a, err := Parse(expression)
	if err != nil {
		return 0, err
	}
	return evaluate(a)
}
