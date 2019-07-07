package main

import (
	"fmt"
)

type DivideError struct {
	dividee int
	divider int
}

func (de *DivideError) Error() string {
	strFormat := `
		Cannot proceed, the devider is zero.
		devidee: %d
		devider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

func Divide(varDevidee int, varDevider int) (result int, errorMsg string) {
	if varDevider == 0 {
		dData := DivideError{
			dividee: varDevidee,
			divider: varDevider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDevidee / varDevider, ""
	}
}

func main() {
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}

	// 当被除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}
}
