package xtime

import (
	"time"
)

// SetTimezone sets timezone.
// 设置时区
func (c Time) SetTimezone(name string) Time {
	if c.Error != nil {
		return c
	}
	c.loc, c.Error = getLocationByTimezone(name)
	return c
}

// SetTimezone sets timezone.
// 设置时区
func SetTimezone(name string) Time {
	return NewTime().SetTimezone(name)
}

// SetLocale sets locale.
// 设置语言区域
func (c Time) SetLocale(locale string) Time {
	if c.Error != nil {
		return c
	}
	c.lang.SetLocale(locale)
	c.Error = c.lang.Error
	return c
}

// SetLocale sets locale.
// 设置语言区域
func SetLocale(locale string) Time {
	c := NewTime()
	c.lang.SetLocale(locale)
	c.Error = c.lang.Error
	return c
}

// SetLanguage sets language.
// 设置语言对象
func (c Time) SetLanguage(lang *Language) Time {
	if c.Error != nil {
		return c
	}
	c.lang, c.Error = lang, lang.Error
	return c
}

// SetLanguage sets language.
// 设置语言对象
func SetLanguage(lang *Language) Time {
	c := NewTime()
	lang.SetLocale(lang.locale)
	c.lang, c.Error = lang, lang.Error
	return c
}

// SetYear sets year.
// 设置年份
func (c Time) SetYear(year int) Time {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(year, time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.loc)
	return c
}

// SetYearNoOverflow sets year without overflowing month.
// 设置年份(月份不溢出)
func (c Time) SetYearNoOverflow(year int) Time {
	if c.IsInvalid() {
		return c
	}
	return c.AddYearsNoOverflow(year - c.Year())
}

// SetMonth sets month.
// 设置月份
func (c Time) SetMonth(month int) Time {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year(), time.Month(month), c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.loc)
	return c
}

// SetMonthNoOverflow sets month without overflowing month.
// 设置月份(月份不溢出)
func (c Time) SetMonthNoOverflow(month int) Time {
	if c.IsInvalid() {
		return c
	}
	return c.AddMonthsNoOverflow(month - c.Month())
}

// SetWeekStartsAt sets start day of the week.
// 设置一周的开始日期
func (c Time) SetWeekStartsAt(day string) Time {
	if c.IsInvalid() {
		return c
	}
	switch day {
	case Sunday:
		c.weekStartsAt = time.Sunday
	case Monday:
		c.weekStartsAt = time.Monday
	case Tuesday:
		c.weekStartsAt = time.Tuesday
	case Wednesday:
		c.weekStartsAt = time.Wednesday
	case Thursday:
		c.weekStartsAt = time.Thursday
	case Friday:
		c.weekStartsAt = time.Friday
	case Saturday:
		c.weekStartsAt = time.Saturday
	}
	return c
}

// SetDay sets day.
// 设置日期
func (c Time) SetDay(day int) Time {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year(), time.Month(c.Month()), day, c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.loc)
	return c
}

// SetHour sets hour.
// 设置小时
func (c Time) SetHour(hour int) Time {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), hour, c.Minute(), c.Second(), c.Nanosecond(), c.loc)
	return c
}

// SetMinute sets minute.
// 设置分钟
func (c Time) SetMinute(minute int) Time {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), minute, c.Second(), c.Nanosecond(), c.loc)
	return c
}

// SetSecond sets second.
// 设置秒数
func (c Time) SetSecond(second int) Time {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), second, c.Nanosecond(), c.loc)
	return c
}

// SetMillisecond sets millisecond.
// 设置毫秒
func (c Time) SetMillisecond(millisecond int) Time {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), c.Second(), millisecond*1e6, c.loc)
	return c
}

// SetMicrosecond sets microsecond.
// 设置微秒
func (c Time) SetMicrosecond(microsecond int) Time {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), c.Second(), microsecond*1e3, c.loc)
	return c
}

// SetNanosecond sets nanosecond.
// 设置纳秒
func (c Time) SetNanosecond(nanosecond int) Time {
	if c.IsInvalid() {
		return c
	}
	c.time = time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), c.Second(), nanosecond, c.loc)
	return c
}
