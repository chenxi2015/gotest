package main

import (
	"errors"
	"fmt"
	"math"
)

//type error interface {
//	Error() string
//}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: error")
	}
	return math.Sqrt(f), nil
}

func main() {
	result, err := Sqrt(-1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
