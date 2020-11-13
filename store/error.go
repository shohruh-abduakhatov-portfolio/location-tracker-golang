package store

import "errors"

type ValidationError error

var (
	errIncorrectData = ValidationError(errors.New("Incorrect Data"))
	errLng           = ValidationError(errors.New("Cannot parse Longitude"))
	errLat           = ValidationError(errors.New("Cannot parse Latitude"))
)

func IsValidationError(err error) bool {
	_, ok := err.(ValidationError)
	return ok
}
