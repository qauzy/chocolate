package xtime

import "time"

// DaysInYear gets total days in year.
// 获取本年的总天数
func (c Time) DaysInYear() int {
	if c.IsInvalid() {
		return 0
	}
	if c.IsLeapYear() {
		return DaysPerLeapYear
	}
	return DaysPerNormalYear
}

// DaysInMonth gets total days in month.
// 获取本月的总天数
func (c Time) DaysInMonth() int {
	if c.IsInvalid() {
		return 0
	}
	return c.EndOfMonth().time.In(c.loc).Day()
}

// MonthOfYear gets month of year.
// 获取本年的第几月
func (c Time) MonthOfYear() int {
	if c.IsInvalid() {
		return 0
	}
	return int(c.time.In(c.loc).Month())
}

// DayOfYear gets day of year.
// 获取本年的第几天
func (c Time) DayOfYear() int {
	if c.IsInvalid() {
		return 0
	}
	return c.time.In(c.loc).YearDay()
}

// DayOfMonth gets day of month.
// 获取本月的第几天
func (c Time) DayOfMonth() int {
	if c.IsInvalid() {
		return 0
	}
	return c.time.In(c.loc).Day()
}

// DayOfWeek gets day of week.
// 获取本周的第几天
func (c Time) DayOfWeek() int {
	if c.IsInvalid() {
		return 0
	}
	day := int(c.time.In(c.loc).Weekday())
	if day == 0 {
		return DaysPerWeek
	}
	return day
}

// WeekOfYear gets week of year, see https://en.wikipedia.org/wiki/ISO_8601#Week_dates.
// 获取本年的第几周
func (c Time) WeekOfYear() int {
	if c.IsInvalid() {
		return 0
	}
	_, week := c.time.In(c.loc).ISOWeek()
	return week
}

// WeekOfMonth gets week of month.
// 获取本月的第几周
func (c Time) WeekOfMonth() int {
	if c.IsInvalid() {
		return 0
	}
	days := c.Day() + c.StartOfMonth().DayOfWeek() - 1
	if days%DaysPerWeek == 0 {
		return days / DaysPerWeek
	}
	return days/DaysPerWeek + 1
}

// Century gets current century.
// 获取当前世纪
func (c Time) Century() int {
	if c.IsInvalid() {
		return 0
	}
	return c.Year()/YearsPerCentury + 1
}

// Decade gets current decade.
// 获取当前年代
func (c Time) Decade() int {
	if c.IsInvalid() {
		return 0
	}
	return c.Year() % YearsPerCentury / YearsPerDecade * YearsPerDecade
}

// Year gets current year.
// 获取当前年
func (c Time) Year() int {
	if c.IsInvalid() {
		return 0
	}
	return c.time.In(c.loc).Year()
}

// Quarter gets current quarter.
// 获取当前季度
func (c Time) Quarter() (quarter int) {
	if c.IsInvalid() {
		return 0
	}
	switch {
	case c.Month() >= 10:
		quarter = 4
	case c.Month() >= 7:
		quarter = 3
	case c.Month() >= 4:
		quarter = 2
	case c.Month() >= 1:
		quarter = 1
	}
	return
}

// Month gets current month.
// 获取当前月
func (c Time) Month() int {
	if c.IsInvalid() {
		return 0
	}
	return c.MonthOfYear()
}

// Week gets current week, start from 0.
// 获取当前周(从0开始)
func (c Time) Week() int {
	if c.IsInvalid() {
		return -1
	}
	return (c.DayOfWeek() + DaysPerWeek - int(c.weekStartsAt)) % DaysPerWeek
}

// Day gets current day.
// 获取当前日
func (c Time) Day() int {
	if c.IsInvalid() {
		return 0
	}
	return c.DayOfMonth()
}

// Hour gets current hour.
// 获取当前小时
func (c Time) Hour() int {
	if c.IsInvalid() {
		return 0
	}
	return c.time.In(c.loc).Hour()
}

// Minute gets current minute.
// 获取当前分钟数
func (c Time) Minute() int {
	if c.IsInvalid() {
		return 0
	}
	return c.time.In(c.loc).Minute()
}

// Second gets current second.
// 获取当前秒数
func (c Time) Second() int {
	if c.IsInvalid() {
		return 0
	}
	return c.time.In(c.loc).Second()
}

// Millisecond gets current millisecond.
// 获取当前毫秒数，3位数字
func (c Time) Millisecond() int {
	if c.IsInvalid() {
		return 0
	}
	return c.time.In(c.loc).Nanosecond() / 1e6
}

// Microsecond gets current microsecond.
// 获取当前微秒数，6位数字
func (c Time) Microsecond() int {
	if c.IsInvalid() {
		return 0
	}
	return c.time.In(c.loc).Nanosecond() / 1e3
}

// Nanosecond gets current nanosecond.
// 获取当前纳秒数，9位数字
func (c Time) Nanosecond() int {
	if c.IsInvalid() {
		return 0
	}
	return c.time.In(c.loc).Nanosecond()
}

// Timestamp gets timestamp with second, it is shorthand for TimestampWithSecond.
// 获取秒级时间戳, 是 TimestampWithSecond 的简写
func (c Time) Timestamp() int64 {
	if c.IsInvalid() {
		return 0
	}
	return c.TimestampWithSecond()
}

// TimestampWithSecond gets timestamp with second.
// 输出秒级时间戳
func (c Time) TimestampWithSecond() int64 {
	if c.IsInvalid() {
		return 0
	}
	return c.time.In(c.loc).Unix()
}

// TimestampWithMillisecond gets timestamp with millisecond.
// 获取毫秒级时间戳
func (c Time) TimestampWithMillisecond() int64 {
	if c.IsInvalid() {
		return 0
	}
	return c.time.UnixNano() / int64(time.Millisecond)
}

// TimestampWithMicrosecond gets timestamp with microsecond.
// 获取微秒级时间戳
func (c Time) TimestampWithMicrosecond() int64 {
	if c.IsInvalid() {
		return 0
	}
	return c.time.UnixNano() / int64(time.Microsecond)
}

// TimestampWithNanosecond gets timestamp with nanosecond.
// 获取纳秒级时间戳
func (c Time) TimestampWithNanosecond() int64 {
	if c.IsInvalid() {
		return 0
	}
	return c.time.UnixNano()
}

// Location gets location name.
// 获取位置
func (c Time) Location() string {
	return c.loc.String()
}

// Timezone gets timezone name.
// 获取时区
func (c Time) Timezone() string {
	name, _ := c.time.Zone()
	return name
}

// Offset gets offset seconds from the UTC timezone.
// 获取距离UTC时区的偏移量，单位秒
func (c Time) Offset() int {
	_, offset := c.time.Zone()
	return offset
}

// Locale gets locale name.
// 获取语言区域
func (c Time) Locale() string {
	return c.lang.locale
}

// Age gets age.
// 获取年龄
func (c Time) Age() int {
	if c.IsInvalid() {
		return 0
	}
	now := Now()
	if c.Timestamp() > now.Timestamp() {
		return 0
	}
	return int(c.DiffInYears(now))
}
