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
