package simpleMaths

import (
	"errors"
	"fmt"
	"unicode"
)

func mul(a, b int64) (int64, error) {
	return a * b, nil
}
func add(a, b int64) (int64, error) {
	return a + b, nil
}
func div(a, b int64) (int64, error) {
	if b == 0 {
		return 0, errors.New("zero is not divisible")
	}
	return a / b, nil
}
func sub(a, b int64) (int64, error) {
	return a - b, nil
}

func Exparser(ex string, window ...int64) {
	var i, j int64

	if len(window) < 2 {
		i, j = 0, int64(len([]rune(ex))-1)
	} else {
		i, j = window[0], window[1]
	}

	runes := []rune(ex) // handles Unicode safely
	if unicode.IsDigit(runes[i]) {
		fmt.Println("It's a digit")
	} else if unicode.IsLetter(runes[i]) {
		fmt.Println("It's a letter")
	} else {
		fmt.Println("It's something else")
	}

	fmt.Println(runes[j])
	if unicode.IsDigit(runes[j]) {
		fmt.Println("It's a digit")
	} else if unicode.IsLetter(runes[j]) {
		fmt.Println("It's a letter")
	} else {
		fmt.Println("It's something else")
	}
}
