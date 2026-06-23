package transformations

import (
	"strings"
)

func Punc(input string) string {
	puncs := ",.;:'?!"
	newStr := strings.Fields(input)
	for i, v := range newStr {
		for _, p := range puncs {
			if strings.HasPrefix(v, string(p)) && i > 0 {
				newStr[i-1] = newStr[i-1] + string(p)
				newStr[i] = strings.Trim(newStr[i], string(p))
			}
		}
	}
	return strings.Join(newStr, " ")
}

// func main() {
// 	fmt.Println(Punc("Punctuation tests are ... kinda boring ,what do you think ?"))
// }
