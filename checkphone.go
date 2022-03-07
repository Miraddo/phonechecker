package checkphone

import (
	"regexp"
)

// CheckIranPhone checking Iran Phone numbers.
func CheckIranPhone(n string) bool {
	match, _ := regexp.MatchString(`(?m)^(?:98|\+98|0098|0)?9[0-9]{9}$`, n)
	return match
}
