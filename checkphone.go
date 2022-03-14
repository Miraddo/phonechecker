package checkphone

import (
	"regexp"
)

var (
	iranPhoneChecker   = regexp.MustCompile(`(?m)^(?:98|\+98|0098|0)?9[0-9]{9}$`)
	germanPhoneChecker = regexp.MustCompile(`(?m)^(([+]{1}[1-9]{1}[0-9]{0,2}[ ]{1}([1-9]{1}[0-9]{1,4}){1}[ ]{1}([1-9]{1}[0-9]{2,6}){1}([ -][0-9]{1,5})?)|([0]{1}[1-9]{1}[0-9]{1,4}[ ]{1}[0-9]{1,8}([ -][0-9]{1,8})?)?)`)
)

// Iran checking Iran Phone numbers.
func Iran(n string) bool {
	return iranPhoneChecker.MatchString(n)
}

// German checking Iran Phone numbers.
func German(n string) bool {
	return germanPhoneChecker.MatchString(n)
}

// CheckIranPhone checking Iran Phone numbers.
// func CheckIranPhone(n string) bool {
// 	match, _ := regexp.MatchString(`(?m)^(?:98|\+98|0098|0)?9[0-9]{9}$`, n)
// 	return match
// }
