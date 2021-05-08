package parser

import (
	"fmt"
	"strings"
)

type cronIdentifier struct {
	typ   cronType
	fn    cronFn
	token string
}

type cronElement struct {
	ids  []cronIdentifier
	runs string
}

func (c *cronElement) valid() bool {
	for _, id := range c.ids {
		if id.typ == invalid {
			return false
		}
	}
	return true
}

type CronExpression struct {
	Minute     cronElement
	Hour       cronElement
	DayOfMonth cronElement
	Month      cronElement
	DayOfWeek  cronElement
	Command    cronElement
}

func (c *CronExpression) Valid() bool {
	return c.Minute.valid() && c.Hour.valid() && c.DayOfMonth.valid() && c.Month.valid() && c.DayOfWeek.valid()
}

func (c *CronExpression) String() string {
	return fmt.Sprintf(
		"%-14s%v\n%-14s%v\n%-14s%v\n%-14s%v\n%-14s%v\n%-14s%v",
		"minute", c.Minute.runs,
		"hour", c.Hour.runs,
		"day of month", c.DayOfMonth.runs,
		"month", c.Month.runs,
		"day of week", c.DayOfWeek.runs,
		"command", c.Command.runs,
	)
}

type CronExpressionParser struct {
	Expression string
}

func (p *CronExpressionParser) Parse() (cron *CronExpression, err error) {
	tokens := lex(p.Expression)

	if len(tokens.elements) != 7 { //incl EOF
		return nil, fmt.Errorf("invalid expression")
	}

	c := &CronExpression{
		Minute:     parseCronElement(tokens.elements[0].val),
		Hour:       parseCronElement(tokens.elements[1].val),
		DayOfMonth: parseCronElement(tokens.elements[2].val),
		Month:      parseCronElement(tokens.elements[3].val),
		DayOfWeek:  parseCronElement(tokens.elements[4].val),
		Command:    cronElement{runs: tokens.elements[5].val},
	}
	if !c.Valid() {
		return nil, fmt.Errorf("an invalid element or elements within expression")
	}

	for f := cronMinute; f != nil; {
		f = f(c)
	}

	return c, nil
}

func parseCronElement(token string) cronElement {
	var tokens []string
	//multi with comma
	multi := strings.Split(token, ",")
	for _, m := range multi {
		if strings.Trim(m, "") == "" {
			continue
		}
		tokens = append(tokens, m)
	}
	//identify type
	cronElem := cronElement{}
	for _, t := range tokens {
		if cronStarRegEx.MatchString(t) {
			cronElem.ids = append(cronElem.ids, cronIdentifier{star, cronStar, t})
		} else if cronNumberRegEx.MatchString(t) {
			cronElem.ids = append(cronElem.ids, cronIdentifier{number, cronNumber, t})
		} else if cronRangeRegEx.MatchString(t) {
			cronElem.ids = append(cronElem.ids, cronIdentifier{rangeLoHi, cronRange, t})
		} else if cronStepRegEx.MatchString(t) {
			step := cronStepRegEx.FindStringSubmatch(t)
			if cronStarRegEx.MatchString(step[1]) {
				cronElem.ids = append(cronElem.ids, cronIdentifier{stepStar, cronStepStar, t})
			} else if cronNumberRegEx.MatchString(step[1]) {
				cronElem.ids = append(cronElem.ids, cronIdentifier{stepNumber, cronStepNumber, t})
			} else if cronRangeRegEx.MatchString(step[1]) {
				cronElem.ids = append(cronElem.ids, cronIdentifier{stepRange, cronStepRange, t})
			} else {
				cronElem.ids = append(cronElem.ids, cronIdentifier{invalid, cronErr, t})
			}
		} else {
			cronElem.ids = append(cronElem.ids, cronIdentifier{invalid, cronErr, t})
		}
	}
	return cronElem
}

type timeFn func(*CronExpression) timeFn

func cronMinute(c *CronExpression) timeFn {
	c.Minute.runs = getWhenCronWillRun(0, 59, c.Minute.ids)
	return cronHour
}

func cronHour(c *CronExpression) timeFn {
	c.Hour.runs = getWhenCronWillRun(0, 23, c.Hour.ids)
	return cronDayOfMonth
}

func cronDayOfMonth(c *CronExpression) timeFn {
	c.DayOfMonth.runs = getWhenCronWillRun(1, 31, c.DayOfMonth.ids)
	return cronMonth
}

func cronMonth(c *CronExpression) timeFn {
	c.Month.runs = getWhenCronWillRun(1, 12, c.Month.ids)
	return cronDayOfWeek
}

func cronDayOfWeek(c *CronExpression) timeFn {
	c.DayOfWeek.runs = getWhenCronWillRun(0, 6, c.DayOfWeek.ids)
	return nil
}
