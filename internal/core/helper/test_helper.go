package helper

import "testing"

func Test_check_card_number_should_return_fail_when_card_number_pattern_is_not_valid(t *testing.T) {

	cardNumber := "1957202120304154"
	if CheckCardNumber(cardNumber) {
		t.Error("Failed To Check " + cardNumber)
	}

	cardNumber = "1234123412341234"
	if CheckCardNumber(cardNumber) {
		t.Error("Failed To Check " + cardNumber)
	}
}

func Test_should_check_arabic_or_persian_number_conversion(t *testing.T) {

	cardNumber := "۱۰۳۶۷۵۱"
	cardNumber = FixCardNumberString(cardNumber)
	if cardNumber != "1036751" {
		t.Error("Failed To Check Out" + cardNumber)
	}
}
