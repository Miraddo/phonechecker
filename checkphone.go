package checkphone

import (
	"regexp"
)

var (
	IranPhoneChecker = regexp.MustCompile(`(?m)^(?:98|\+98|0098|0)?9[0-9]{9}$`)
)

// Iran checking Iran Phone numbers.
func Iran(n string) bool {
	return IranPhoneChecker.MatchString(n)
}

// CheckIranPhone checking Iran Phone numbers.
// func CheckIranPhone(n string) bool {
// 	match, _ := regexp.MatchString(`(?m)^(?:98|\+98|0098|0)?9[0-9]{9}$`, n)
// 	return match
// }
