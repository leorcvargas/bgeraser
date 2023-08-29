package domainerrors

import "errors"

var ErrImageProcessEmptyResultID = errors.New("the provided result ID is empty")
