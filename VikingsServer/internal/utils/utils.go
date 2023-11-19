package utils

import (
	"VikingsServer/internal/app/ds"
	"fmt"
	"github.com/rs/xid"
	"strings"
)

// MARK: - Генератор имени для фото

func GenerateUniqueName(imageName *string) error {
	parts := strings.Split(*imageName, ".")
	if len(parts) > 1 {
		fileExt := parts[len(parts)-1]
		uniqueID := xid.New()
		*imageName = fmt.Sprintf("%s.%s", uniqueID.String(), fileExt)
		return nil
	}
	return fmt.Errorf("uncorrect file name. not fount image extension")
}

// Contains Функция для проверки наличия элемента в срезе
func Contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

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
