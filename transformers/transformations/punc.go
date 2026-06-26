package transformations

import (
	"regexp"
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
	var result []string
	for _, v := range newStr {
		if v != "" {
			result = append(result, v)
		}
	}
	return strings.Join(result, " ")
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

func MPunc(s string) string {
	re := regexp.MustCompile(` {2,}`)

	processQuotes := func(input string, quote byte) string {
		parts := strings.Split(input, string(quote))
		for i := 1; i < len(parts); i += 2 {
			content := strings.TrimSpace(parts[i])
			parts[i] = re.ReplaceAllString(content, " ")

			if i > 0 && parts[i-1] != "" && !strings.HasSuffix(parts[i-1], " ") {
				parts[i-1] += " "
			}
			if i+1 < len(parts) && parts[i+1] != "" && !strings.HasPrefix(parts[i+1], " ") {
				parts[i+1] = " " + parts[i+1]
			}
		}
		return strings.Join(parts, string(quote))
	}

	s1, s2 := strings.IndexByte(s, '\''), strings.IndexByte(s, '"')
	if (s1 > s2 && s2 != -1) || s1 == -1 {
		s = processQuotes(s, '\'')
		s = processQuotes(s, '"')
	} else {
		s = processQuotes(s, '"')
		s = processQuotes(s, '\'')
	}
	return strings.TrimSpace(re.ReplaceAllString(s, " "))
}
