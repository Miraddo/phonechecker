package checkphone_test

import (
	"testing"

	"github.com/miraddo/checkphone"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestCheckPhone(t *testing.T) {
	// a list of wrong phone numbers!
	list := []string{"+98937538285", "98937538285", "0937538285", "937538285", "093753828522", "08375382852"}

	// starting the test
	t.Log("Given the need to test New functionality.")
	{
		t.Logf("\tTest 0:\tWhen adding a list of wrong Iran Phone numbers.")
		{
			for _, p := range list {
				if checkphone.Iran(p) != false {
					t.Logf("\t%s\tTest 0:\tShould not be able to return true for %s items.", failed, p)
				}
			}
			t.Logf("\t%s\tTest 0:\tEverything looks fine.", succeed)
		}
	}
}
