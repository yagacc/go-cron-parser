package parser

import (
	"regexp"
	"strings"
)

var cronExpressionRegex = regexp.MustCompile("^\"?(.*?)\"?$")

type elementType int

const (
	elemErr elementType = iota
	elemEOF
	elemText
)

type element struct {
	typ elementType
	val string
}

type lexer struct {
	elements []element
}

func lex(input string) *lexer {
	l := &lexer{
		elements: []element{},
	}

	match := cronExpressionRegex.FindStringSubmatch(input)
	if match == nil {
		l.elements = append(l.elements, element{elemErr, "invalid cron expression"})
	} else {
		tokens := strings.Split(match[1], " ")
		for _, t := range tokens {
			if strings.Trim(t, "") == "" {
				continue
			}
			l.elements = append(l.elements, element{elemText, t})
		}
		l.elements = append(l.elements, element{elemEOF, ""})
	}

	return l
}
