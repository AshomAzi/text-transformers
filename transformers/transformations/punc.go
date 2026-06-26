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

func MPunc(input string) string {
	isContractionSuffix := func(s string) bool {
		suffixes := []string{"t", "s", "m", "d", "re", "ll", "ve"}
		lower := strings.ToLower(s)
		for _, suffix := range suffixes {
			if lower == suffix {
				return true
			}
		}
		return false
	}

	re := regexp.MustCompile(`"|'|[^\s"']+`)
	tokens := re.FindAllString(input, -1)

	var merged []string
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "'" && i > 0 && i+1 < len(tokens) &&
			tokens[i-1] != `"` && tokens[i-1] != `'` &&
			tokens[i+1] != `"` && tokens[i+1] != `'` &&
			isContractionSuffix(tokens[i+1]) {
			merged[len(merged)-1] += "'" + tokens[i+1]
			i++
			continue
		}
		merged = append(merged, tokens[i])
	}
	tokens = merged

	isSolePunct := func(s string) bool {
		if len(s) == 0 {
			return false
		}
		for _, c := range s {
			if !strings.ContainsRune(",.;:!?", c) {
				return false
			}
		}
		return true
	}

	openQuotes := map[string]bool{`"`: false, `'`: false}

	for i := 0; i < len(tokens); i++ {
		t := tokens[i]
		if t != `"` && t != `'` {
			continue
		}

		if !openQuotes[t] {
			j := i + 1
			for j < len(tokens) && (tokens[j] == `"` || tokens[j] == `'`) {
				j++
			}
			if j < len(tokens) {
				tokens[j] = t + tokens[j]
				openQuotes[t] = true
			}
		} else {
			j := i - 1
			for j >= 0 && (tokens[j] == `"` || tokens[j] == `'`) {
				j--
			}
			if j >= 0 {
				tokens[j] += t
				openQuotes[t] = false
			}
		}
	}

	var filtered []string
	for _, t := range tokens {
		if t != `"` && t != `'` {
			filtered = append(filtered, t)
		}
	}
	tokens = filtered

	merged = []string{}
	for i := 0; i < len(tokens); i++ {
		if i > 0 && isSolePunct(tokens[i]) && tokens[i-1] != `"` && tokens[i-1] != `'` {
			merged[len(merged)-1] += tokens[i]
			continue
		}
		merged = append(merged, tokens[i])
	}
	tokens = merged

	var result []string
	for _, t := range tokens {
		if t != `"` && t != `'` {
			result = append(result, t)
		}
	}
	return strings.Join(result, " ")
}
