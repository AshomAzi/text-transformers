package transformations

import "strings"

func Vowels(input string) string {
	vowels := "aeiouAEIOU"
	newStr := strings.Fields(input)
	for i, v := range newStr {
		for _, char := range vowels {
			if i < len(newStr)-1 && v == "a" && string(newStr[i+1][0]) == string(char) {
				newStr[i] = "an"
			}
			if i < len(newStr)-1 && v == "A" && string(newStr[i+1][0]) == string(char) {
				newStr[i] = "An"
			}
		}
	}
	return strings.Join(newStr, " ")
}
