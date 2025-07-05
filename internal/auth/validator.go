package user

import "errors"

func ValidatateName(name string) error {
	if len(name) < 1 {
		return errors.New("name cannot be null")
	}
	if len(name) < 17 {
		return errors.New("name maximum 16 character")
	}
	return nil
}
