package service

import (
	"cc-validation/model"
	"regexp"
	"strconv"
	"strings"

	"github.com/kataras/iris/v12"
)

var _ccMap = map[string]string{
	"34":   "AMEX",
	"37":   "AMEX",
	"4":    "Visa",
	"51":   "MasterCard",
	"52":   "MasterCard",
	"53":   "MasterCard",
	"54":   "MasterCard",
	"55":   "MasterCard",
	"6011": "Discover",
}

// ValidateCreditCard validation route service
func ValidateCreditCard(ctx iris.Context) {

	var inputParam model.InputParameter

	_ = ctx.ReadJSON(&inputParam)

	if inputParam.CardNumber == "" {
		ctx.StatusCode(400)
		ctx.JSON(model.Response{
			CardNumber:   inputParam.CardNumber,
			Status:       "Invalid",
			ErrorMessage: "Card number is required",
		})
		return
	}

	cardType := CheckCardType(inputParam.CardNumber)

	if cardType == "" || !CheckCardValidity(inputParam.CardNumber) {
		ctx.StatusCode(200)
		ctx.JSON(model.Response{
			CardNumber: inputParam.CardNumber,
			CardType:   "unknown",
			Status:     "Invalid",
		})
		return
	}

	ctx.StatusCode(200)
	ctx.JSON(model.Response{
		CardNumber: inputParam.CardNumber,
		CardType:   cardType,
		Status:     "Valid",
	})

}

/*
CheckCardType -
	check if card type meets cc type requirements
					begins with		length
		Amex  			34 or 37	15
		Discover		6011		16
		MasterCard		51-55		16
		Visa 			4			13 or 6

*/
func CheckCardType(cardNum string) string {

	if len(cardNum) > 16 || len(cardNum) < 13 {
		return ""
	}

	re := regexp.MustCompile(`^(34|37|[5-5][1-5]|4|6011)`)

	return _ccMap[re.FindString(cardNum)]
}

//CheckCardValidity - using luhn algrorithm check card validity
func CheckCardValidity(cardNum string) bool {

	if len(cardNum) > 16 || len(cardNum) < 13 {
		return false
	}

	var (
		cardSlice = strings.Split(cardNum, "")
		doubleVal = false
		total     = 0
		val       = 0
	)

	for i := len(cardSlice) - 1; i >= 0; i-- {

		val, _ = strconv.Atoi(cardSlice[i])

		if doubleVal {
			val = val * 2
		}

		doubleVal = !doubleVal

		if val > 9 {
			val = val - 9
		}

		total += val

	}

	if total%10 != 0 {
		return false
	}

	return true

}
