package parser

import "testing"

func TestParseCronElement(t *testing.T) {
	type expected struct {
		val string
		typ cronType
	}
	testTokens := map[string][]expected{
		//errors
		"****": {{"****", invalid}},
		"*/x":  {{"*/x", invalid}},
		"x/9":  {{"x/9", invalid}},
		//valid
		"*/15":         {{"*/15", stepStar}},
		"5/15":         {{"5/15", stepNumber}},
		"1-5/15":       {{"1-5/15", stepRange}},
		"*":            {{"*", star}},
		"5":            {{"5", number}},
		"1-5":          {{"1-5", rangeLoHi}},
		"*/15,*/15":    {{"*/15", stepStar}, {"*/15", stepStar}},
		"*/15,,,,*/15": {{"*/15", stepStar}, {"*/15", stepStar}},
	}
	for expr, x := range testTokens {
		crons := parseCronElement(expr)
		for idx, val := range x {
			if val.val != crons.ids[idx].token {
				t.Errorf("%s: got\t%v, expected\t%v", expr, crons.ids[idx].token, val.val)
			}
			if val.typ != crons.ids[idx].typ {
				t.Errorf("%s: got\t%v, expected\t%v", expr, crons.ids[idx].typ, val.typ)
			}
		}
	}
}

func TestParse(t *testing.T) {
	type expected struct {
		min, hr, dom, mth, dow, cmd string
	}
	testTokens := map[string]*expected{
		//errors
		"":                                  nil,
		"*/avds 0 1,15 * 1-5 /usr/bin/find": nil,
		//valid
		"*/15 0 1,15 * 1-5 /usr/bin/find": {"0 15 30 45", "0", "1 15", "1 2 3 4 5 6 7 8 9 10 11 12", "1 2 3 4 5", "/usr/bin/find"},
	}
	for expr, x := range testTokens {
		p := CronExpressionParser{Expression: expr}
		cron, err := p.Parse()
		if x != nil && err != nil {
			t.Errorf("%s: got\t%v, expected\t%v", expr, err, x)
		}
		if err == nil {
			if cron.Minute.runs != x.min {
				t.Errorf("%s: Minute got\t%v, expected\t%v", expr, cron.Minute.runs, x.min)
			}
			if cron.Hour.runs != x.hr {
				t.Errorf("%s: Hour got\t%v, expected\t%v", expr, cron.Hour.runs, x.hr)
			}
			if cron.DayOfMonth.runs != x.dom {
				t.Errorf("%s: DayOfMonth got\t%v, expected\t%v", expr, cron.DayOfMonth.runs, x.dom)
			}
			if cron.Month.runs != x.mth {
				t.Errorf("%s: Month got\t%v, expected\t%v", expr, cron.Month.runs, x.mth)
			}
			if cron.DayOfWeek.runs != x.dow {
				t.Errorf("%s: DayOfWeek got\t%v, expected\t%v", expr, cron.DayOfWeek.runs, x.dow)
			}
			if cron.Command.runs != x.cmd {
				t.Errorf("%s: Command got\t%v, expected\t%v", expr, cron.Command.runs, x.cmd)
			}
		}
	}
}
