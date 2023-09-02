package domainerrors

import "errors"

var (
	ErrImageProcessEmptyResult = errors.New("the provided result is empty")
	ErrImageProcessNotFound    = errors.New("image process not found")
)
