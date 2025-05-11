package models

import "errors"

var (
	ErrObjectNotFound   = errors.New("not found")
	ErrValidationFailed = errors.New("validation failed")
	ErrNotExpired       = errors.New("not expired")
	ErrAlreadyIssued    = errors.New("already issued")
	ErrForbidden        = errors.New("forbidden")
	ErrAlreadyExpired   = errors.New("already expired")
	ErrNotIssued        = errors.New("not issued")
	ErrInvalidRole      = errors.New("invalid role")

	ErrTimeParse = errors.New("incorrect time")
	ErrIntParse  = errors.New("incorrect int")
	ErrBoolParse = errors.New("incorrect bool")

	ErrBlankOrderIDs            = errors.New("order ids cannot be blank")
	ErrBlankUserID              = errors.New("user id cannot be blank")
	ErrEncryptionFailed         = errors.New("password encryption failed")
	ErrEmailAlreadyExists       = errors.New("such email already exists")
	ErrLoginAlreadyExists       = errors.New("such login already exists")
	ErrPhoneNumberAlreadyExists = errors.New("such phone number already exists")
	ErrAlreadyAvailable         = errors.New("physical book already available")
	ErrAlreadyUnavailable       = errors.New("physical book already unavailable")
	ErrLibCardExpired           = errors.New("library card expired")
	ErrIncorrectOrderStatus     = errors.New("incorrect order status")
	ErrNoRows                   = errors.New("now rows")
	ErrPendingOrders            = errors.New("there are pending orders")
	ErrRepeatedPassword         = errors.New("password must be different from previous")
	ErrCardNotExpired           = errors.New("card have not expired yet")

	ErrInternal = errors.New("internal error")
)
