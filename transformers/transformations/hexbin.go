package transformations

import (
	"fmt"
	"strconv"
	"strings"
)

func HexBin(text string) string {
	newStr := strings.Fields(text)
	for i := 0; i < len(newStr); i++ {
		switch newStr[i] {
		case "(hex)":
			newStrConv, err := strconv.ParseInt(newStr[i-1], 16, 64)
			if err == nil {
				newStr[i-1] = strconv.FormatInt(newStrConv, 10)
				newStr = append(newStr[:i], newStr[i+1:]...)
			} else {
				fmt.Println(err)
			}
		case "(bin)":
			newStrConv, err := strconv.ParseInt(newStr[i-1], 2, 64)
			if err == nil {
				newStr[i-1] = strconv.FormatInt(newStrConv, 10)
				newStr = append(newStr[:i], newStr[i+1:]...)
			} else {
				fmt.Println(err)
			}
		}
	}
	return strings.Join(newStr, " ")
}
