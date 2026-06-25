package transformations

import (
	"strconv"
	"strings"
)

func Cases(s string) string {
	new := strings.Fields(s)
	for i, v := range new {
		switch v {
		case "(cap)":
			new[i-1] = strings.ToUpper(string(new[i-1][0])) + strings.ToLower(string(new[i-1][1:]))
			new = append(new[:i], new[i+1:]...)
		case "(up)":
			new[i-1] = strings.ToUpper(new[i-1])
			new = append(new[:i], new[i+1:]...)
		case "(low)":
			new[i-1] = strings.ToLower(new[i-1])
			new = append(new[:i], new[i+1:]...)
		case "(up,":
			if strings.HasPrefix(new[i+1], ")") && i < 0 {
				snum := strings.TrimRight(new[i+1], ")")
				num, _ := strconv.Atoi(snum)
				for j := 1; j <= num && i-j >= 0; j++ {
					new[i-j] = strings.ToUpper(new[i-j])
				}
				new = append(new[:i], new[i+2:]...)
			}
		}
	}
	return strings.Join(new, " ")
}

func CaseN(s string) string {
	new := strings.Fields(s)
	for i := 0; i < len(new); i++ {
		if new[i] == "(up," && i > 0 {
			snum := strings.TrimRight(new[i+1], ")")
			num, _ := strconv.Atoi(snum)
			for j := 1; j <= num && i-j >= 0; j++ {
				new[i-j] = strings.ToUpper(new[i-j])
			}
			new = append(new[:i], new[i+2:]...)
		} else if new[i] == "(cap," && i > 0 {
			snum := strings.TrimRight(new[i+1], ")")
			num, _ := strconv.Atoi(snum)
			for j := 1; j <= num && i-j >= 0; j++ {
				new[i-j] = strings.ToUpper(string(new[i-j])[:1]) + strings.ToLower(string(new[i-j][1:]))
			}
			new = append(new[:i], new[i+2:]...)
		} else if new[i] == "(low," && i > 0 {
			snum := strings.TrimRight(new[i+1], ")")
			num, _ := strconv.Atoi(snum)
			for j := 1; j <= num && i-j >= 0; j++ {
				new[i-j] = strings.ToLower(new[i-j])
			}
			new = append(new[:i], new[i+2:]...)
		}
	}
	return strings.Join(new, " ")
}
