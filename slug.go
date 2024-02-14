package goslug

import (
	"net/url"
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

// URL-encoded slug
func Slug(in string) string {
	// 1) string operation without any translation
	replacer := strings.NewReplacer("ß", "ss", "tak zwany", "tzw") // TODO when???
	res := []byte(strings.Map(reRule, replacer.Replace(strings.ToLower(string(in)))))
	// % only for escaping encoded characters - only before valid encoded symbols!
	// replace % with %-encoded '%'
	reSpace := regexp.MustCompile(`[\s\\\/?&#.,;:*!%]+`)
	res = reSpace.ReplaceAll([]byte(res), []byte("-"))
	// TODO encode matched characters
	// percent encoding "other" characters
	reOther := regexp.MustCompile(`[^a-z0-9\-\+]`)
	res = reOther.ReplaceAll(res, []byte("-"))

	reTrim := regexp.MustCompile(`^-+|-+$`)
	reReduce := regexp.MustCompile(`-{2,}`)
	res = reReduce.ReplaceAll(reTrim.ReplaceAll(res, nil), []byte("-"))

	return url.QueryEscape(string(res))
}

// C, Bash, PHP, GO - compatible variable name
func Var(in string) string {
	sb := strings.Builder{}
	begin := true
	for _, c := range in {
		if isVarChar(c) {
			if !(begin && '0' <= c && c <= '9') {
				sb.WriteRune(c)
				begin = false
			}
		}
	}
	return sb.String()
}

func isVarChar(c rune) bool {
	if ('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') || c == '_' {
		return true
	}
	return false
}

var notAllowedWin = []rune{'<', '>', ':', '/', '\\', '|', '?', '*'}

// Windows 32 filename
func FileNameWin(in string) string {
	sb := strings.Builder{}
	for _, c := range in {
		if c >= 32 && sliceIndex(notAllowedWin, c) == -1 {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}

// POSIX compatible filename
func FileNamePosix(in string) string {
	sb := strings.Builder{}
	for _, c := range in {
		if ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || c == '.' || c == '-' || c == '_' {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}

func sliceIndex(s []rune, c rune) int {
	for i, r := range s {
		if r == c {
			return i
		}
	}
	return -1
}
