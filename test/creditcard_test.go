package test

import (
	"cc-validation/service"
	"testing"
)

func TestCheckCardValidity(t *testing.T) {

	var (
		cardNumber string
		isValid    bool
	)

	cardNumber = "4111111111111111"
	isValid = service.CheckCardValidity(cardNumber)

	if !isValid {
		t.Errorf("validate cardnumber: %v using luhn algorithm, got: %v, want: true", cardNumber, isValid)
	}

	cardNumber = "41111111111111182"
	isValid = service.CheckCardValidity(cardNumber)

	if isValid {
		t.Errorf("validate cardnumber: %v using luhn algorithm, got: %v, want: false", cardNumber, isValid)
	}

}

func TestCheckCardType(t *testing.T) {

	var (
		cardNumber string
		cardType   string
	)

	//validate visa
	cardNumber = "4111111111111111"
	cardType = service.CheckCardType(cardNumber)

	if cardType != "Visa" {
		t.Errorf("validate cardnumber: %v, got: %v, want:%v", cardNumber, cardType, "Visa")
	}

	//validate mastercard
	cardNumber = "5105105105105100"
	cardType = service.CheckCardType(cardNumber)

	if cardType != "MasterCard" {
		t.Errorf("validate cardnumber: %v, got: %v, want:%v", cardNumber, cardType, "MasterCard")
	}

	//validate amex
	cardNumber = "378282246310005"
	cardType = service.CheckCardType(cardNumber)

	if cardType != "AMEX" {
		t.Errorf("validate cardnumber: %v, got: %v, want:%v", cardNumber, cardType, "AMEX")
	}

	//validate discover
	cardNumber = "6011111111111117"
	cardType = service.CheckCardType(cardNumber)

	if cardType != "Discover" {
		t.Errorf("validate cardnumber: %v, got: %v, want:%v", cardNumber, cardType, "Discover")
	}

}
