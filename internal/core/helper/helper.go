package helper

import (
	"github.com/pkg/errors"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// Arabic digits
var arabicTemplate = []string{
	"\u0660",
	"\u0661",
	"\u0662",
	"\u0663",
	"\u0664",
	"\u0665",
	"\u0666",
	"\u0667",
	"\u0668",
	"\u0669",
}

var persianTemplate = []string{
	"\u06f0",
	"\u06f1",
	"\u06f2",
	"\u06f3",
	"\u06f4",
	"\u06f5",
	"\u06f6",
	"\u06f7",
	"\u06f8",
	"\u06f9",
}

func GO_UNUSED(sth ...any) {}

func CheckCardNumber(cardNumberStr string) bool {
	if len(cardNumberStr) != 16 {
		return false
	}
	var cardTotal int64 = 0
	for i, ch := range cardNumberStr {
		c, err := strconv.ParseInt(string(ch), 10, 8)
		if err != nil {
			return false
		}
		if i%2 == 0 {
			if c*2 > 9 {
				cardTotal = cardTotal + (c * 2) - 9
			} else {
				cardTotal = cardTotal + (c * 2)
			}
		} else {
			cardTotal += c
		}
	}

	return cardTotal%10 == 0
}

func FixCardNumberString(cardNumber string) string {
	tmpString := cardNumber
	//tmpString := "۱۰۳۶۷۵۱"

	for i, _ := range persianTemplate {
		tmpString = strings.Replace(tmpString, persianTemplate[i], strconv.Itoa(i), -1)
		tmpString = strings.Replace(tmpString, arabicTemplate[i], strconv.Itoa(i), -1)

		//fmt.Printf("%s => %s or %s \n", strconv.Itoa(i), persianTemplate[i], arabicTemplate[i])
	}
	//fmt.Printf("%s", tmpString)
	return tmpString
}

func Filename() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("unable to get the current filename")
	}
	return filename, nil
}

func Dirname() (string, error) {
	filename, err := Filename()
	if err != nil {
		return "", err
	}
	return filepath.Dir(filename), nil
}
