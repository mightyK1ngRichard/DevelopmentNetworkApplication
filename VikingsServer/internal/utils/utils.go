package utils

import (
	"VikingsServer/internal/app/ds"
	"fmt"
)

// MARK: - Поиск элемента

func FindElement(slice []ds.City, target ds.City) int {
	for i, val := range slice {
		if val == target {
			return i
		}
	}

	return -1
}

func Max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func Min(num1 int, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

// MARK: - Просто поигрался

type KingURL string

const (
	Cities KingURL = "/cities"
	Hikes  KingURL = "/hikes"
)

func (king KingURL) GenerateEndpoint() (error, string) {
	baseURL := "api/v3"
	switch king {
	case Cities:
		return nil, baseURL + string(Cities)
	case Hikes:
		return nil, baseURL + string(Hikes)
	}

	return fmt.Errorf("unknown endpoint"), ""
}

// MARK: - Фанюсь

func generateSqlCommand(params ...interface{}) (string, string) {
	resultParams := ""
	resultValues := ""
	for i, param := range params {
		if i == 0 {
			switch param.(type) {
			case string:
				value := param.(string)
				if value != "" {
					resultParams += value
					resultValues += value
				}
			case int:
				value := param.(int)
				if value != -1 {
					resultParams += fmt.Sprintf("%d", value)
					resultValues += fmt.Sprintf("'%d'", value)
				}
			}
			continue
		}

		switch param.(type) {
		case string:
			value := param.(string)
			if value != "" {
				resultParams += "," + value
				resultValues += fmt.Sprintf(", '%s'", value)
			}
		case int:
			value := param.(int)
			if value != -1 {
				resultParams += fmt.Sprintf(", %d", value)
				resultValues += fmt.Sprintf(", '%d'", value)
			}
		}
	}

	return resultParams, resultValues
}
