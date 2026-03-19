package api

import "errors"

//  В данном классе мы будем хранить api-ошибки

var ErrNotFound = errors.New("not found")

var ErrNotSupported = errors.New("not supported")

var ErrIllegalArgument = errors.New("illegal argument provided")
