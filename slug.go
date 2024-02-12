package goslug

import (
	"regexp"
	"strings"
)

var reMap = map[rune]rune{
	'ą': 'a', 'ć': 'c', 'ę': 'e', 'ł': 'l', 'ń': 'n', 'ó': 'o', 'ś': 's', 'ź': 'z', 'ż': 'z',
}

func reRule(r rune) rune {
	if repl, ok := reMap[r]; ok {
		return repl
	}
	return r
}

func slug(in string) string {
	// 1) string operation without any translation
	replacer := strings.NewReplacer("ß", "ss", "tak zwany", "tzw") // TODO when???
	res := []byte(strings.Map(reRule, replacer.Replace(strings.ToLower(string(in)))))
	// % only for escaping encoded characters - only before valid encoded symbols!
	// replace % with %-encoded '%'
	reSpace := regexp.MustCompile(`[\s\\\/?&#.,;:*!%]+`)
	res = reSpace.ReplaceAll([]byte(res), []byte("-"))
	// TODO encode matched characters
	reOther := regexp.MustCompile(`[^a-z0-9\-\+]`)
	res = reOther.ReplaceAll(res, []byte("-"))

	reTrim := regexp.MustCompile(`^-+|-+$`)
	reReduce := regexp.MustCompile(`-{2,}`)
	res = reReduce.ReplaceAll(reTrim.ReplaceAll(res, nil), []byte("-"))

	return string(res)
}
