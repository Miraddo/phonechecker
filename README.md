# Checking Phone Numbers

this is a simple package just to check phone number input to be sure that the number is correct or not.

## Installation

```
go get github.com/miraddo/checkphone
```


## Using

```go
package main

import (
	"fmt"

	"github.com/miraddo/checkphone"
)

func main() {
	input := "+989372222222"

	fmt.Println(checkphone.CheckIranPhone(input)) // true

	input2 := "+9893722222222"

	fmt.Println(checkphone.CheckIranPhone(input2)) // false

}
```
