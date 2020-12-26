package repository

import "errors"

// ErrNoSuchEntity ...
var ErrNoSuchEntity = errors.New("no such entity")

// ErrBadRequest ...
var ErrBadRequest = errors.New("bad request")
