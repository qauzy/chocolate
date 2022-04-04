package xtime

import (
	"time"
)

// StartOfCentury returns a Time instance for start of the century.
// 本世纪开始时间
func (c Time) StartOfCentury() Time {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year()/YearsPerCentury*YearsPerCentury, 1, 1, 0, 0, 0, 0, c.loc)
	return c
}

// EndOfCentury returns a Time instance for end of the century.
// 本世纪结束时间
func (c Time) EndOfCentury() Time {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year()/YearsPerCentury*YearsPerCentury+99, 12, 31, 23, 59, 59, 999999999, c.loc)
	return c
}

// StartOfDecade returns a Time instance for start of the decade.
// 本年代开始时间
func (c Time) StartOfDecade() Time {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year()/YearsPerDecade*YearsPerDecade, 1, 1, 0, 0, 0, 0, c.loc)
	return c
}

// EndOfDecade returns a Time instance for end of the decade.
// 本年代结束时间
func (c Time) EndOfDecade() Time {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year()/YearsPerDecade*YearsPerDecade+9, 12, 31, 23, 59, 59, 999999999, c.loc)
	return c
}

// StartOfYear returns a Time instance for start of the year.
// 本年开始时间
func (c Time) StartOfYear() Time {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year(), 1, 1, 0, 0, 0, 0, c.loc)
	return c
}

// EndOfYear returns a Time instance for end of the year.
// 本年结束时间
func (c Time) EndOfYear() Time {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year(), 12, 31, 23, 59, 59, 999999999, c.loc)
	return c
}

// StartOfQuarter returns a Time instance for start of the quarter.
// 本季度开始时间
func (c Time) StartOfQuarter() Time {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year(), time.Month(3*c.Quarter()-2), 1, 0, 0, 0, 0, c.loc)
	return c
}

// EndOfQuarter returns a Time instance for end of the quarter.
// 本季度结束时间
func (c Time) EndOfQuarter() Time {
	if c.IsInvalid() {
		return c
	}
	quarter, day := c.Quarter(), 30
	switch quarter {
	case 1, 4:
		day = 31
	case 2, 3:
		day = 30
	}
	c.time = time.Date(c.Year(), time.Month(3*quarter), day, 23, 59, 59, 999999999, c.loc)
	return c
}

// StartOfMonth returns a Time instance for start of the month.
// 本月开始时间
func (c Time) StartOfMonth() Time {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year(), time.Month(c.Month()), 1, 0, 0, 0, 0, c.loc)
	return c
}

// EndOfMonth returns a Time instance for end of the month.
// 本月结束时间
func (c Time) EndOfMonth() Time {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year(), time.Month(c.Month()), 1, 23, 59, 59, 999999999, c.loc).AddDate(0, 1, -1)
	return c
}

// StartOfWeek returns a Time instance for start of the week.
// 本周开始时间
func (c Time) StartOfWeek() Time {
	if c.IsInvalid() {
		return c
	}
	dayOfWeek, weekStartsAt := c.DayOfWeek(), int(c.weekStartsAt)
	return c.SubDays((DaysPerWeek + dayOfWeek - weekStartsAt) % DaysPerWeek).StartOfDay()
}

// EndOfWeek returns a Time instance for end of the week.
// 本周结束时间
func (c Time) EndOfWeek() Time {
	if c.IsInvalid() {
		return c
	}
	dayOfWeek, weekEndsAt := c.DayOfWeek(), int(c.weekStartsAt)+DaysPerWeek-1
	return c.AddDays((DaysPerWeek - dayOfWeek + weekEndsAt) % DaysPerWeek).EndOfDay()
}

// StartOfDay returns a Time instance for start of the day.
// 本日开始时间
func (c Time) StartOfDay() Time {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), 0, 0, 0, 0, c.loc)
	return c
}

// EndOfDay returns a Time instance for end of the day.
// 本日结束时间
func (c Time) EndOfDay() Time {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), 23, 59, 59, 999999999, c.loc)
	return c
}

// StartOfHour returns a Time instance for start of the hour.
// 小时开始时间
func (c Time) StartOfHour() Time {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), 0, 0, 0, c.loc)
	return c
}

// EndOfHour returns a Time instance for end of the hour.
// 小时结束时间
func (c Time) EndOfHour() Time {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), 59, 59, 999999999, c.loc)
	return c
}

// StartOfMinute returns a Time instance for start of the minute.
// 分钟开始时间
func (c Time) StartOfMinute() Time {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), 0, 0, c.loc)
	return c
}

// EndOfMinute returns a Time instance for end of the minute.
// 分钟结束时间
func (c Time) EndOfMinute() Time {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), 59, 999999999, c.loc)
	return c
}

// StartOfSecond returns a Time instance for start of the second.
// 秒开始时间
func (c Time) StartOfSecond() Time {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), c.Second(), 0, c.loc)
	return c
}

// EndOfSecond returns a Time instance for end of the second.
// 秒结束时间
func (c Time) EndOfSecond() Time {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), c.Second(), 999999999, c.loc)
	return c
}
