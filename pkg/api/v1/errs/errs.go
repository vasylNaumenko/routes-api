/*
 * telegram: @VasylNaumenko
 */

package errs

import (
	"errors"
)

// List of errors, used across services
var (
	ErrNotFound = errors.New("not found")
	ErrNotValid = errors.New("") // used as an error flag
)
