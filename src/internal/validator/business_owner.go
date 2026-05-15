package validator

import (
	"backend/src/internal/domain"
	"errors"
	"regexp"
	"strings"
	"time"
)

func ValidateBusinessOwner(owner *domain.BusinessOwner) error {
	if err := validateName(owner.FirstName, "first_name"); err != nil {
		return err
	}

	if err := validateName(owner.LastName, "last_name"); err != nil {
		return err
	}

	if owner.MiddleName != nil && *owner.MiddleName != "" {
		if err := validateName(*owner.MiddleName, "middle_name"); err != nil {
			return err
		}
	}

	if err := validateINN(owner.INN); err != nil {
		return err
	}

	if err := validatePhone(owner.PhoneNumber); err != nil {
		return err
	}

	if err := validateEmail(owner.Email); err != nil {
		return err
	}

	if err := validatePassword(owner.Password); err != nil {
		return err
	}

	if err := validateBirthday(owner.Birthday); err != nil {
		return err
	}

	return nil
}

func validateName(name, field string) error {
	if name == "" {
		return errors.New(field + " is required")
	}

	if len(name) < 2 {
		return errors.New(field + " must be at least 2 characters")
	}

	if len(name) > 50 {
		return errors.New(field + " must be less than 50 characters")
	}

	matched, _ := regexp.MatchString(`^[a-zA-Zа-яА-ЯёЁ\-\s]+$`, name)
	if !matched {
		return errors.New(field + " must contain only letters, hyphens and spaces")
	}
	return nil
}

func validateINN(inn string) error {
	if inn == "" {
		return errors.New("inn is required")
	}

	inn = strings.TrimSpace(inn)

	if len(inn) != 10 && len(inn) != 12 {
		return errors.New("inn must be 10 or 12 digits")
	}

	matched, _ := regexp.MatchString(`^\d+$`, inn)
	if !matched {
		return errors.New("inn must contain only digits")
	}

	return nil
}

func validatePhone(phone string) error {
	if phone == "" {
		return errors.New("phone_number is required")
	}
	phone = strings.TrimSpace(phone)

	matched, _ := regexp.MatchString(`^(\+7|7|8)?\d{10}$`, phone)
	if !matched {
		return errors.New("phone_number must be valid Russian phone number")
	}

	return nil
}

func validateEmail(email string) error {
	if email == "" {
		return errors.New("email is required")
	}

	matched, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, email)
	if !matched {
		return errors.New("email is invalid")
	}

	return nil
}

func validatePassword(password string) error {
	if password == "" {
		return errors.New("password is required")
	}

	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	if len(password) > 30 {
		return errors.New("password must be less than 30 characters")
	}

	return nil
}

func validateBirthday(birthday time.Time) error {
	if birthday.IsZero() {
		return errors.New("birthday is required")
	}

	now := time.Now()
	age := now.Year() - birthday.Year()

	if birthday.After(now) {
		return errors.New("birthday cannot be in the future")
	}

	if age < 18 {
		return errors.New("owner must be at least 18 years old")
	}

	if age > 120 {
		return errors.New("invalid birthday")
	}

	return nil
}
