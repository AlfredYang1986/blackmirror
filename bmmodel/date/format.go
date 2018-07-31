package date

import (
	"time"
)

type TimeFormat int

const (
	TimeYearMonth TimeFormat = 1
	TimeMonthDay  TimeFormat = 2
	TimeDayOfWeek TimeFormat = 3
)

type DDTime struct {
	Timespan time.Time
}

func dateFromSpan(span int64) DDTime { // NOTE: span is second
	return DDTime{
		Timespan: time.Unix(span, 0)}
}

func (t *DDTime) Format(f TimeFormat) string {
	switch f {
	case TimeYearMonth:
		return t.Timespan.Format("2018-07")
	case TimeMonthDay:
		return t.Timespan.Format("07-31")
	case TimeDayOfWeek:
		return t.Timespan.Format("Mon")
	}

	return ""
}
