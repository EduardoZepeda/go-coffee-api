package validator

import (
	"net/mail"

	"github.com/EduardoZepeda/go-coffee-api/models"
)

func ValidateCoffeeShop(v *Validator, coffeeShop *models.CoffeeShop) {
	v.Validate(len(coffeeShop.Name) >= 3 && len(coffeeShop.Name) <= 100, "Name", "Must be greater than 3 characters and less than 100 chars")
	v.Validate(len(coffeeShop.Address) >= 5 && len(coffeeShop.Address) <= 100, "Address", "Must be greater than 3 chars and less than 100 chars")
	v.Validate(coffeeShop.Rating >= 0 && coffeeShop.Rating <= 5.0, "Rating", "Rating must be a floating number between 0 and 5.0")
}

func ValidateUserSignup(v *Validator, user *models.SignUpRequest) {
	_, emailError := mail.ParseAddress(user.Email)
	v.Validate(emailError == nil, "Email", "Please enter a valid email address")
	v.Validate(user.Password == user.PasswordConfirmation, "Password confirmation", "Password and password confirmation didn't match")
}

func ValidateUserUpdate(v *Validator, user *models.UpdateUserRequest) {
	v.Validate(len(user.Bio) <= 250, "Bio", "Your profile Bio can't be greater than 250 chars")
}
