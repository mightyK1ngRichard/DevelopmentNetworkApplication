package utils

import (
	"fmt"
)

/*
Просто игрался с enum )
Жаль как в swift не получится
*/

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
