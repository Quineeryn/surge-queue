package model

import "myAPI/pkg/shared/common"

type User struct {
	common.CommonModel
	Name    string
	Email   string
	Balance int
}
