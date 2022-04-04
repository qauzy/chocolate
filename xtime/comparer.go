package xtime

import (
	"time"
)

// IsZero reports whether is zero time.
// 是否是零值时间
func (c Xtime) IsZero() bool {
	return c.time.IsZero()
}

// IsInvalid reports whether is invalid time.
// 是否是无效时间
func (c Xtime) IsInvalid() bool {
	if c.Error != nil || c.IsZero() {
		return true
	}
	return false
}

// IsNow reports whether is now time.
// 是否是当前时间
func (c Xtime) IsNow() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Timestamp() == c.Now().Timestamp()
}

// IsFuture reports whether is future time.
// 是否是未来时间
func (c Xtime) IsFuture() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Timestamp() > c.Now().Timestamp()
}

// IsPast reports whether is past time.
// 是否是过去时间
func (c Xtime) IsPast() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Timestamp() < c.Now().Timestamp()
}

// IsLeapYear reports whether is a leap year.
// 是否是闰年
func (c Xtime) IsLeapYear() bool {
	if c.IsInvalid() {
		return false
	}
	year := c.time.In(c.loc).Year()
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	}
	return false
}

// IsLongYear reports whether is a long year, see https://en.wikipedia.org/wiki/ISO_8601#Week_dates.
// 是否是长年
func (c Xtime) IsLongYear() bool {
	if c.IsInvalid() {
		return false
	}
	_, w := time.Date(c.Year(), time.December, 31, 0, 0, 0, 0, c.loc).ISOWeek()
	return w == weeksPerLongYear
}

// IsJanuary reports whether is January.
// 是否是一月
func (c Xtime) IsJanuary() bool {
	if c.IsInvalid() {
		return false
	}
	return c.time.In(c.loc).Month() == time.January
}

// IsFebruary reports whether is February.
// 是否是二月
func (c Xtime) IsFebruary() bool {
	if c.IsInvalid() {
		return false
	}
	return c.time.In(c.loc).Month() == time.February
}

// IsMarch reports whether is March.
// 是否是三月
func (c Xtime) IsMarch() bool {
	if c.IsInvalid() {
		return false
	}
	return c.time.In(c.loc).Month() == time.March
}

// IsApril reports whether is April.
// 是否是四月
func (c Xtime) IsApril() bool {
	if c.IsInvalid() {
		return false
	}
	return c.time.In(c.loc).Month() == time.April
}

// IsMay reports whether is May.
// 是否是五月
func (c Xtime) IsMay() bool {
	if c.IsInvalid() {
		return false
	}
	return c.time.In(c.loc).Month() == time.May
}

// IsJune reports whether is June.
// 是否是六月
func (c Xtime) IsJune() bool {
	if c.IsInvalid() {
		return false
	}
	return c.time.In(c.loc).Month() == time.June
}

// IsJuly reports whether is July.
// 是否是七月
func (c Xtime) IsJuly() bool {
	if c.IsInvalid() {
		return false
	}
	return c.time.In(c.loc).Month() == time.July
}

// IsAugust reports whether is August.
// 是否是八月
func (c Xtime) IsAugust() bool {
	if c.IsInvalid() {
		return false
	}
	return c.time.In(c.loc).Month() == time.August
}

// IsSeptember reports whether is September.
// 是否是九月
func (c Xtime) IsSeptember() bool {
	if c.IsInvalid() {
		return false
	}
	return c.time.In(c.loc).Month() == time.September
}

// IsOctober reports whether is October.
// 是否是十月
func (c Xtime) IsOctober() bool {
	if c.IsInvalid() {
		return false
	}
	return c.time.In(c.loc).Month() == time.October
}

// IsNovember reports whether is November.
// 是否是十一月
func (c Xtime) IsNovember() bool {
	if c.IsInvalid() {
		return false
	}
	return c.time.In(c.loc).Month() == time.November
}

// IsDecember reports whether is December.
// 是否是十二月
func (c Xtime) IsDecember() bool {
	if c.IsInvalid() {
		return false
	}
	return c.time.In(c.loc).Month() == time.December
}

// IsMonday reports whether is Monday.
// 是否是周一
func (c Xtime) IsMonday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.time.In(c.loc).Weekday() == time.Monday
}

// IsTuesday reports whether is Tuesday.
// 是否是周二
func (c Xtime) IsTuesday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.time.In(c.loc).Weekday() == time.Tuesday
}

// IsWednesday reports whether is Wednesday.
// 是否是周三
func (c Xtime) IsWednesday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.time.In(c.loc).Weekday() == time.Wednesday
}

// IsThursday reports whether is Thursday.
// 是否是周四
func (c Xtime) IsThursday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.time.In(c.loc).Weekday() == time.Thursday
}

// IsFriday reports whether is Friday.
// 是否是周五
func (c Xtime) IsFriday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.time.In(c.loc).Weekday() == time.Friday
}

// IsSaturday reports whether is Saturday.
// 是否是周六
func (c Xtime) IsSaturday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.time.In(c.loc).Weekday() == time.Saturday
}

// IsSunday reports whether is Sunday.
// 是否是周日
func (c Xtime) IsSunday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.time.In(c.loc).Weekday() == time.Sunday
}

// IsWeekday reports whether is weekday.
// 是否是工作日
func (c Xtime) IsWeekday() bool {
	if c.IsInvalid() {
		return false
	}
	return !c.IsSaturday() && !c.IsSunday()
}

// IsWeekend reports whether is weekend.
// 是否是周末
func (c Xtime) IsWeekend() bool {
	if c.IsInvalid() {
		return false
	}
	return c.IsSaturday() || c.IsSunday()
}

// IsYesterday reports whether is yesterday.
// 是否是昨天
func (c Xtime) IsYesterday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.ToDateString() == Now().SubDay().ToDateString()
}

// IsToday reports whether is today.
// 是否是今天
func (c Xtime) IsToday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.ToDateString() == c.Now().ToDateString()
}

// IsTomorrow reports whether is tomorrow.
// 是否是明天
func (c Xtime) IsTomorrow() bool {
	if c.IsInvalid() {
		return false
	}
	return c.ToDateString() == Now().AddDay().ToDateString()
}

// Compare compares by an operator.
// 时间比较
func (c Xtime) Compare(operator string, t Xtime) bool {
	switch operator {
	case "=":
		return c.Eq(t)
	case "<>", "!=":
		return !c.Eq(t)
	case ">":
		return c.Gt(t)
	case ">=":
		return c.Gte(t)
	case "<":
		return c.Lt(t)
	case "<=":
		return c.Lte(t)
	}
	return false
}

// Gt reports whether greater than.
// 是否大于
func (c Xtime) Gt(t Xtime) bool {
	return c.time.After(t.time)
}

// Lt reports whether less than.
// 是否小于
func (c Xtime) Lt(t Xtime) bool {
	return c.time.Before(t.time)
}

// Eq reports whether equal.
// 是否等于
func (c Xtime) Eq(t Xtime) bool {
	return c.time.Equal(t.time)
}

// Ne reports whether not equal.
// 是否不等于
func (c Xtime) Ne(t Xtime) bool {
	return !c.Eq(t)
}

// Gte reports whether greater than or equal.
// 是否大于等于
func (c Xtime) Gte(t Xtime) bool {
	return c.Gt(t) || c.Eq(t)
}

// Lte reports whether less than or equal.
// 是否小于等于
func (c Xtime) Lte(t Xtime) bool {
	return c.Lt(t) || c.Eq(t)
}

// Between reports whether between two times, excluded the start and end time.
// 是否在两个时间之间(不包括这两个时间)
func (c Xtime) Between(start Xtime, end Xtime) bool {
	if c.Gt(start) && c.Lt(end) {
		return true
	}
	return false
}

// BetweenIncludedStart reports whether between two times, included the start time.
// 是否在两个时间之间(包括开始时间)
func (c Xtime) BetweenIncludedStart(start Xtime, end Xtime) bool {
	if c.Gte(start) && c.Lt(end) {
		return true
	}
	return false
}

// BetweenIncludedEnd reports whether between two times, included the end time.
// 是否在两个时间之间(包括结束时间)
func (c Xtime) BetweenIncludedEnd(start Xtime, end Xtime) bool {
	if c.Gt(start) && c.Lte(end) {
		return true
	}
	return false
}

// BetweenIncludedBoth reports whether between two times, included the start and end time.
// 是否在两个时间之间(包括这两个时间)
func (c Xtime) BetweenIncludedBoth(start Xtime, end Xtime) bool {
	if c.Gte(start) && c.Lte(end) {
		return true
	}
	return false
}
