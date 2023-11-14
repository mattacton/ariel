package termio

import (
	"fmt"
)

func PromptRead(prompt string, isValidFn func(string) bool) string {
	var input string
	fmt.Print(prompt)
	fmt.Scanf("%s", &input)

	if !isValidFn(input) {
		return PromptRead(prompt, isValidFn)
	}
	return input
}
