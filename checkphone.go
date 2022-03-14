package checkphone

import (
	"regexp"
)

var (
	IranPhoneChecker   = regexp.MustCompile(`(?m)^(?:98|\+98|0098|0)?9[0-9]{9}$`)
	GermanPhoneChecker = regexp.MustCompile(`(\(?([\d \-\)\–\+\/\(]+){6,}\)?([ .\-–\/]?)([\d]+))`)
)

// Iran checking Iran Phone numbers.
func Iran(n string) bool {
	return IranPhoneChecker.MatchString(n)
}

// German checking Iran Phone numbers.
func German(n string) bool {
	return GermanPhoneChecker.MatchString(n)
}

// CheckIranPhone checking Iran Phone numbers.
// func CheckIranPhone(n string) bool {
// 	match, _ := regexp.MatchString(`(?m)^(?:98|\+98|0098|0)?9[0-9]{9}$`, n)
// 	return match
// }
