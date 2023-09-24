package handler

import "errors"

var (
	idNotFound          = errors.New("param `id` not found")
	idMustBeEmpty       = errors.New("param `id` must be empty")
	cityCannotBeEmpty   = errors.New("city name cannot be empty")
	headerNotFound      = errors.New("no file uploaded")
	vikingCannotBeEmpty = errors.New("viking name cannot be empty")
)
