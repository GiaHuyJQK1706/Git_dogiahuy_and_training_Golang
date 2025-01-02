package validator

import (
	"errors"
	"regexp"
	"time"
)

// ValidatePassword check mat khau
func ValidatePassword(password string) error {
	if len(password) < 8 || len(password) > 20 {
		return errors.New("password must be between 8 and 20 characters")
	}

	// Regex to check for uppercase, lowercase, and digit
	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLowercase := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)

	if !hasUppercase || !hasLowercase || !hasDigit {
		return errors.New("password must contain at least one uppercase letter, one lowercase letter, and one number")
	}
	
	return nil
}

// ValidateCategory check category hop le
var validCategories = map[string]bool{
	"client":       true,
	"non-billable": true,
	"system":       true,
}

func ValidateCategory(category string) error {
	if _, exists := validCategories[category]; !exists {
		return errors.New("category must be one of 'client', 'non-billable', 'system'")
	}

	return nil
}

// ValidateProjectDates check thoi gian bat dau va ket thuc cua project
func ValidateProjectDates(start, end time.Time) error {
	now := time.Now()

	if start.Before(now) {
		return errors.New("project_started_at must be greater than the current time")
	}

	if !end.IsZero() && end.Before(start) {
		return errors.New("project_ended_at must be greater than project_started_at")
	}

	return nil
}
