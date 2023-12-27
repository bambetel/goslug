package main

import "fmt"

func main() {
	test := []string{"Czy miał/a Pan/Pani dylemat moralny?", "Tak zwany Marszałek Pisłsudski - sługą szatana?", "Wykonanie artykułu 5 w Przemyślu (będzie wojna) 14 kwietnia 2024?", "<div><p>paragraf w HTML. Czy on <strong>jest</strong>dobrze wyescapowany?", "# Lista w markdown\n\nParagraf w markdown wraz z listą poniżej.\n* first item\n* second item\n* third item\n\n Second paragraph in markdown.", "Zażółć gęślą jaźń", "Scheißen"}

	for _, s := range test {
		fmt.Println(slug(s))
	}
}
