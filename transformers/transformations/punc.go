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

func Puncs(input string) string {
	puncs := ",.:;'!?"
	newStr := strings.Fields(input)
	var out []string
	for _, v := range newStr {
		if len(v) > 1 && len(out) > 0 {
			allPunc := true
			for _, r := range v {
				if !strings.ContainsRune(puncs, r) {
					allPunc = false
					break
				}
			}
			if allPunc {
				out[len(out)-1] += v
				continue
			}
		}
		out = append(out, v)
	}

	return strings.Join(out, " ")
}
