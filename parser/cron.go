package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type cronType int

const (
	invalid cronType = iota
	rangeLoHi
	star
	number
	stepRange
	stepStar
	stepNumber
)

var (
	cronStepRegEx   = regexp.MustCompile("^(.*?)/(\\d+)$") //STAR or NUMBER or RANGE/NUMBER
	cronStarRegEx   = regexp.MustCompile("^\\*$")
	cronRangeRegEx  = regexp.MustCompile("^(\\d+)-(\\d+)$")
	cronNumberRegEx = regexp.MustCompile("^\\d+$")
)

type cronFn func(string, int) bool

var cronErr = func(string, int) bool {
	//todo log as error/warning
	return false
}

//matches everything
var cronStar = func(string, int) bool { return true }

//only if exact match
var cronNumber = func(token string, i int) bool {
	x, err := strconv.Atoi(token) //todo handle error
	return err == nil && i == x
}

//only within range given
var cronRange = func(token string, i int) bool {
	rLo, rHi, err := rangeToLoHi(token)
	return err == nil && i >= rLo && i <= rHi
}

//only if divisible by step
var cronStepStar = func(token string, i int) bool {
	_, step, err := stepToParts(token)
	return err == nil && i%step == 0
}

//not exactly sure if this is valid but https://crontab.guru/#5/55_4_*_*_* - shows number-hi/step
var cronStepNumber = func(token string, i int) bool {
	numberVal, step, err := stepToParts(token)
	num, err := strconv.Atoi(numberVal)
	return err == nil && i >= num && i%step == 0
}

//only if within range and divisible by step
var cronStepRange = func(token string, i int) bool {
	rangeVal, step, err := stepToParts(token)
	rLo, rHi, err := rangeToLoHi(rangeVal)
	return err == nil && i >= rLo && i <= rHi && i%step == 0
}

//returns a space delimited string on when the cron will run between lo and hi
func getWhenCronWillRun(lo, hi int, ids []cronIdentifier) (runs string) {
	for i := lo; i <= hi; i++ {
		for _, e := range ids {
			if e.fn(e.token, i) {
				runs += fmt.Sprintf("%d ", i)
				break
			}
		}
	}
	if len(runs) > 0 {
		runs = runs[:len(runs)-1]
	}
	return
}

//convert cron range e.g. "1-10" to lo & hi ints: "1, 10"
func rangeToLoHi(r string) (lo, hi int, err error) {
	pos1 := strings.Index(r, "-")
	left, right := r[0:pos1], r[pos1+1:]
	lo, err = strconv.Atoi(left)
	hi, err = strconv.Atoi(right)
	return
}

//convert cron step e.g. "1-10/40" to cron expr: "1-10" and step value "40"
func stepToParts(s string) (id string, step int, err error) {
	pos1 := strings.Index(s, "/")
	id, right := s[0:pos1], s[pos1+1:]
	step, err = strconv.Atoi(right)
	return
}
