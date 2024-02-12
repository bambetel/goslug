package goslug

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"testing"
)

var reValidSlug *regexp.Regexp

type TestPair struct {
	In  string `json:"in"`
	Out string `json:"out"`
}
type TestPairs struct {
	Pairs []TestPair `json:"pairs"`
}

var pairs TestPairs

func init() {
	reValidSlug = regexp.MustCompile(`^([A-Za-z0-9\-\+.]*|\%[0-9A-F]{2,2})*$`)
	fJsonPairs, err := os.ReadFile("testdata/pairs.json")
	if err != nil {
		_ = fmt.Errorf("Cannot load test data: pairs.json")
	}
	json.Unmarshal(fJsonPairs, &pairs)

	// fmt.Printf("test file data:  %q", pairs)
}

func TestSlug(t *testing.T) {
	for i := 0; i < len(pairs.Pairs); i++ {
		got := slug(pairs.Pairs[i].In)
		if got != pairs.Pairs[i].Out {
			t.Errorf("result invalid TODO where, why it might be:\n%s\n%s\n", got, pairs.Pairs[i].Out)
		}
	}
}

func TestSlugAuto(t *testing.T) {
	data := []string{"testowe dane", "Łączność?!"}
	for _, s := range data {
		got := slug(s)
		if !IsValidSlug(got) {
			t.Errorf("result invalid TODO where, why it might be:\n%s\n", got)
		}
	}
}

var validPercentEncoded = []string{"%20", "%20Tu%C5%BC%20po%20%C5%9Bwi%C4%99tach%20rz%C4%85d%20zebra%C5%82%20si%C4%99%2C%20by%20kontynuowa%C4%87%20prace%20nad%20bud%C5%BCetem.%20Kamery%20telewizyjne%20wy%C5%82apa%C5%82y%20pierwsze%20s%C5%82owa%20premiera%20Donalda%20Tuska%2C%20kt%C3%B3ry%20t%C5%82umaczy%C5%82%20op%C3%B3%C5%BAnienie%20rozpocz%C4%99cia%20obrad.%20S%C5%82owa%20by%C5%82y%20skierowane%20ju%C5%BC%20tylko%20do%20cz%C5%82onk%C3%B3w%20rz%C4%85du%2C%20ale%20dziennikarze%20wy%C5%82apali%20je%2C%20zanim%20opu%C5%9Bcili%20sal%C4%99.%20Posiedzenie%20rz%C4%85du%20przed%C5%82u%C5%BCa%20si%C4%99.%20Konferencj%C4%99%20prasow%C4%85%20premiera%20zaplanowano%20dopiero%20na%20godz.%2015%3A00.%20R%C3%B3wnie%C5%BC%20w%20%C5%9Brod%C4%99%20do%20Sejmu%20trafi%C5%82%20prezydencki%20projekt%20dotycz%C4%85cy%20rozwi%C4%85za%C5%84%20bud%C5%BCetowych%2C%20zak%C5%82adaj%C4%85cy%20podwy%C5%BCki%20dla%20nauczycieli%2C%20jednak%20nieuwzgl%C4%99dniaj%C4%85cy%203%20mld%20z%C5%82%20dla%20medi%C3"}

var invalidPercentEncoded = []string{"%#nieprawidlowe", "Tu%5gSK"}

func TestSlugPercentEncoding(t *testing.T) {
	for _, s := range validPercentEncoded {
		if !IsValidSlug(s) {
			t.Errorf("invalid percent encoded slug!: %s\n", s)
		}
	}
	for _, s := range invalidPercentEncoded {
		if IsValidSlug(s) {
			t.Errorf("invalid percent encoded slug!: %s\n", s)
		}
	}
}

func BenchmarkSlug(b *testing.B) {
	l := len(validPercentEncoded)
	for i := 0; i < b.N; i++ {
		slug(validPercentEncoded[i%l])
	}
}

func FuzzSlug(f *testing.F) {
	f.Add("test")
	f.Fuzz(func(t *testing.T, s string) {
		slug(s)
	})
}

func IsValidSlug(s string) bool {
	return reValidSlug.Match(([]byte)(s))
}

func IsPrettySlug(s string) bool {
	if !IsValidSlug(s) {
		return false
	}
	return reValidSlug.Match(([]byte)(s))
}
