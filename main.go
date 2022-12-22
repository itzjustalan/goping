package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func main() {
	fmt.Println("ha")
	validate = validator.New()

	user := &User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "test@g.m",
	}

	errs := validate.Var(user.Email, "required,email")
	if errs != nil {
		fmt.Println(errs)
	}
}

// User contains user information
type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	FavouriteColor string     `validate:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*Address `validate:"required,dive,required"` // a person can have a home and cottage...
}

// Address houses a users address information
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate
