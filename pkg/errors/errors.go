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

// Raised when the config file has incorrect values
type ConfigInvalidError string

func (str ConfigInvalidError) Error() string {
	return fmt.Sprintf("ConfigInvalidError: %q", string(str))
}

// Raised when a dependency (like keychain) is missing
type DependencyMissingError string

func (str DependencyMissingError) Error() string {
	return fmt.Sprintf("DependencyMissingError: %q", string(str))
}
// Raised when CLI arguments are missing or incorrect
type InvalidArgumentsError string

func (str InvalidArgumentsError) Error() string {
	return fmt.Sprintf("InvalidArgumentsError: %q", string(str))
}

//Raised when functionality is not implmented.
type NotImplementedError string

func (str NotImplementedError) Error() string {
	return fmt.Sprintf("NotImplementedError: %q", string(str))
}

// Raised when the Provider returns an error
type ProviderError string
func (str ProviderError) Error() string {
	return fmt.Sprintf("ProviderError: %q", string(str))
}