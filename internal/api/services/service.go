package services

import (
	"errors"
	"sprint1_finalTask/pkg/calc"
	"strings"
)

func Calculate(expression string) (float64, error) {
	res, err := calc.Calc(strings.ReplaceAll(expression, " ", ""))
	if err != nil {
		return 0, errors.New("internal calculator error")
	}
	return res, nil
}
