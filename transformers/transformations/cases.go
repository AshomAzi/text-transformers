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

func CaseN(input string) string {

	newStr := strings.Fields(input)
	for i, v := range newStr {
		switch v {
		case "(up,":
			nextStr := strings.TrimRight(newStr[i+1], ")")
			nextNum, err := strconv.ParseInt(nextStr, 10, 64)
			if err == nil {
				for j := i-1;j >= int(nextNum); j--{
					newStr[j] = strings.ToUpper(newStr[j])
					newStr = append(newStr[:i], newStr[i+1:]... )
				} 
			} else {
				fmt.Println(err)
				return ""
			}
		}
	}

	return strings.Join(newStr, " ")
}

