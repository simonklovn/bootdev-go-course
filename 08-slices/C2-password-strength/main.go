package main

import (
	"fmt"
	"unicode"
)
func isValidPassword(password string) bool {
	hasUppercase, hasDigit := false, false

	// Verify length
	if len(password) < 5 || len(password) > 12 {
		fmt.Println("Password length must be between 5 and 12 characters.")
		return false
	}

	// Check for at least one uppercase letter and one digit
	for _,char := range password {
		if unicode.IsUpper(char) {
			hasUppercase = true
		}
		if unicode.IsDigit(char){
			hasDigit = true
		}
		if hasUppercase && hasDigit {
			break
		}
	}
	if !hasUppercase || !hasDigit {
		fmt.Println("Password must contain at least one uppercase letter and one digit.")
		return false
	}
	return true
}
