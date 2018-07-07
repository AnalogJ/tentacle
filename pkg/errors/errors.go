package errors

import (
	"fmt"
)

// Raised when config file is missing
type ConfigFileMissingError string

func (str ConfigFileMissingError) Error() string {
	return fmt.Sprintf("ConfigFileMissingError: %q", string(str))
}

// Raised when the config file doesnt match schema
type ConfigValidationError string

func (str ConfigValidationError) Error() string {
	return fmt.Sprintf("ConfigValidationError: %q", string(str))
}

// Raised when a dependency (like keychain) is missing
type DependencyMissingError string

func (str DependencyMissingError) Error() string {
	return fmt.Sprintf("DependencyMissingError: %q", string(str))
}

type InvalidArgumentsError string

func (str InvalidArgumentsError) Error() string {
	return fmt.Sprintf("InvalidArgumentsError: %q", string(str))
}