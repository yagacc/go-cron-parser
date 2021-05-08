package parser

import "testing"

func TestLexer(t *testing.T) {
	validElems := []element{
		{typ: elemText, val: "*/15"},
		{typ: elemText, val: "0"},
		{typ: elemText, val: "1,15"},
		{typ: elemText, val: "*"},
		{typ: elemText, val: "1-5"},
		{typ: elemText, val: "/usr/bin/find"},
		{typ: elemEOF},
	}

	expressionTests := map[string][]element{
		"\"*/15 0 1,15 * 1-5 /usr/bin/find\"":                          validElems,
		"*/15 0 1,15 * 1-5 /usr/bin/find":                              validElems,
		"     */15     0     1,15    *   1-5    /usr/bin/find        ": validElems,
	}
	for expr, expected := range expressionTests {
		l := lex(expr)
		if !equal(l.elements, expected) {
			t.Errorf("%s: got\t%v, expected\t%v", expr, l.elements, expected)
		}
	}
}

func equal(i1, i2 []element) bool {
	if len(i1) != len(i2) {
		return false
	}
	for k := range i1 {
		if i1[k].typ != i2[k].typ {
			return false
		}
		if i1[k].val != i2[k].val {
			return false
		}
	}
	return true
}
