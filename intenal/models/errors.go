package models

import "errors"

var (
	ErrObjectNotFound           = errors.New("not found")
	ErrValidationFailed         = errors.New("validation failed")
	ErrNotExpired               = errors.New("not expired")
	ErrAlreadyIssued            = errors.New("already issued")
	ErrForbidden                = errors.New("forbidden")
	ErrAlreadyExpired           = errors.New("already expired")
	ErrNotIssued                = errors.New("not issued")
	ErrCursorParse              = errors.New("invalid cursor")
	ErrLimitParse               = errors.New("invalid limit")
	ErrUserIDParse              = errors.New("invalid user id")
	ErrBlankOrderIDs            = errors.New("order ids cannot be blank")
	ErrEncryptionFailed         = errors.New("password encryption failed")
	ErrEmailAlreadyExists       = errors.New("such email already exists")
	ErrLoginAlreadyExists       = errors.New("such login already exists")
	ErrPhoneNumberAlreadyExists = errors.New("such phone number already exists")
	ErrAlreadyAvailable         = errors.New("physical book already available")
	ErrAlreadyUnavailable       = errors.New("physical book already unavailable")
	ErrLibCardExpired           = errors.New("library card expired")
	ErrIncorrectOrderStatus     = errors.New("incorrect order status")
	ErrNoRows                   = errors.New("now rows")
)
