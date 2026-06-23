package transformations

import (
	"fmt"
	"strconv"
	"strings"
)

func Cases(input string) string {

	newStr := strings.Fields(input)
	for i, v := range newStr {
		switch v {
		case "(up)":
			newStr[i-1] = strings.ToUpper(newStr[i-1])
			newStr = append(newStr[:i], newStr[i+1:]...)
		case "(low)":
			newStr[i-1] = strings.ToLower(newStr[i-1])
			newStr = append(newStr[:i], newStr[i+1:]...)
		case "(cap)":
			newStr[i-1] = strings.ToUpper(newStr[i-1][:1]) + strings.ToLower(newStr[i-1][1:])
			newStr = append(newStr[:i], newStr[i+1:]...)
		}
	}

	return strings.Join(newStr, " ")
}


