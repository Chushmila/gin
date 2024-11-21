// validator/validator.go
package validator

import (
	"fmt"
	"regexp"
)

// Custom error types for invalid inputs.
type InvalidEmailError struct {
	Email string
}

func (e *InvalidEmailError) Error() string {
	return fmt.Sprintf("Invalid email format: %s", e.Email)
}

type InvalidPhoneNumberError struct {
	PhoneNumber string
}

func (e *InvalidPhoneNumberError) Error() string {
	return fmt.Sprintf("Invalid phone number format: %s", e.PhoneNumber)
}

// ValidateEmail checks if the provided email is in a valid format.
func ValidateEmail(email string) error {
	// Basic email validation regex
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(emailRegex, email)
	if !matched {
		return &InvalidEmailError{Email: email}
	}
	return nil
}

// ValidatePhoneNumber checks if the provided phone number is in a valid format.
func ValidatePhoneNumber(phoneNumber string) error {
	// Basic phone number validation regex (allows digits and basic formats)
	phoneRegex := `^\+?[0-9\s()-]{7,15}$`
	matched, _ := regexp.MatchString(phoneRegex, phoneNumber)
	if !matched {
		return &InvalidPhoneNumberError{PhoneNumber: phoneNumber}
	}
	return nil
}
