package tracker

import "errors"

//  В данном классе мы будем хранить бизнес-ошибки

var ErrNotFound = errors.New("not found")

var ErrNotSupported = errors.New("not supported")

var ErrIllegalArgument = errors.New("illegal argument")

var ErrItemNotFound = errors.New("item not found")

var ErrAlreadyExists = errors.New("already exists")
