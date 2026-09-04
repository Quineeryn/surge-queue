package entity

import "errors"

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrEmailDuplicate      = errors.New("email already exists")
	ErrBalanceInsufficient = errors.New("balance is insufficient")
	ErrInvalidTransfer     = errors.New("invalid transfer parameters")
)
