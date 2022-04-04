package xtime

import (
	"time"
)

// AddDuration adds one duration.
// 按照持续时长字符串增加时间,支持整数/浮点数和符号ns(纳秒)、us(微妙)、ms(毫秒)、s(秒)、m(分钟)、h(小时)的组合
func (c Time) AddDuration(duration string) Time {
	if c.IsInvalid() {
		return c
	}
	td, err := parseByDuration(duration)
	c.time = c.time.In(c.loc).Add(td)
	c.Error = err
	return c
}

// SubDuration subtracts one duration.
// 按照持续时长字符串减少时间,支持整数/浮点数和符号ns(纳秒)、us(微妙)、ms(毫秒)、s(秒)、m(分钟)、h(小时)的组合
func (c Time) SubDuration(duration string) Time {
	return c.AddDuration("-" + duration)
}

// AddCenturies adds some centuries.
// N个世纪后
func (c Time) AddCenturies(centuries int) Time {
	return c.AddYears(centuries * YearsPerCentury)
}

// AddCenturiesNoOverflow adds some centuries without overflowing month.
// N个世纪后(月份不溢出)
func (c Time) AddCenturiesNoOverflow(centuries int) Time {
	return c.AddYearsNoOverflow(centuries * YearsPerCentury)
}

// AddCentury adds one century.
// 1个世纪后
func (c Time) AddCentury() Time {
	return c.AddCenturies(1)
}

// AddCenturyNoOverflow adds one century without overflowing month.
// 1个世纪后(月份不溢出)
func (c Time) AddCenturyNoOverflow() Time {
	return c.AddCenturiesNoOverflow(1)
}

// SubCenturies subtracts some centuries.
// N个世纪前
func (c Time) SubCenturies(centuries int) Time {
	return c.SubYears(centuries * YearsPerCentury)
}

// SubCenturiesNoOverflow subtracts some centuries without overflowing month.
// N个世纪前(月份不溢出)
func (c Time) SubCenturiesNoOverflow(centuries int) Time {
	return c.SubYearsNoOverflow(centuries * YearsPerCentury)
}

// SubCentury subtracts one century.
// 1个世纪前
func (c Time) SubCentury() Time {
	return c.SubCenturies(1)
}

// SubCenturyNoOverflow subtracts one century without overflowing month.
// 1个世纪前(月份不溢出)
func (c Time) SubCenturyNoOverflow() Time {
	return c.SubCenturiesNoOverflow(1)
}

// AddDecades adds some decades.
// N个年代后
func (c Time) AddDecades(decades int) Time {
	return c.AddYears(decades * YearsPerDecade)
}

// AddDecadesNoOverflow adds some decades without overflowing month.
// N个年代后(月份不溢出)
func (c Time) AddDecadesNoOverflow(decades int) Time {
	return c.AddYearsNoOverflow(decades * YearsPerDecade)
}

// AddDecade adds one decade.
// 1个年代后
func (c Time) AddDecade() Time {
	return c.AddDecades(1)
}

// AddDecadeNoOverflow adds one decade without overflowing month.
// 1个年代后(月份不溢出)
func (c Time) AddDecadeNoOverflow() Time {
	return c.AddDecadesNoOverflow(1)
}

// SubDecades subtracts some decades.
// N个年代后
func (c Time) SubDecades(decades int) Time {
	return c.SubYears(decades * YearsPerDecade)
}

// SubDecadesNoOverflow subtracts some decades without overflowing month.
// N个年代后(月份不溢出)
func (c Time) SubDecadesNoOverflow(decades int) Time {
	return c.SubYearsNoOverflow(decades * YearsPerDecade)
}

// SubDecade subtracts one decade.
// 1个年代后
func (c Time) SubDecade() Time {
	return c.SubDecades(1)
}

// SubDecadeNoOverflow subtracts one decade without overflowing month.
// 1个年代后(月份不溢出)
func (c Time) SubDecadeNoOverflow() Time {
	return c.SubDecadesNoOverflow(1)
}

// AddYears adds some years.
// N年后
func (c Time) AddYears(years int) Time {
	if c.IsInvalid() {
		return c
	}
	c.time = c.time.In(c.loc).AddDate(years, 0, 0)
	return c
}

// AddYearsNoOverflow adds some years without overflowing month.
// N年后(月份不溢出)
func (c Time) AddYearsNoOverflow(years int) Time {
	if c.IsInvalid() {
		return c
	}
	// 获取N年后本月的最后一天
	last := time.Date(c.Year()+years, time.Month(c.Month()), 1, c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.loc).AddDate(0, 1, -1)
	day := c.Day()
	if c.Day() > last.Day() {
		day = last.Day()
	}
	c.time = time.Date(last.Year(), last.Month(), day, c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.loc)
	return c
}

// AddYear adds one year.
// 1年后
func (c Time) AddYear() Time {
	return c.AddYears(1)
}

// AddYearNoOverflow adds one year without overflowing month.
// 1年后(月份不溢出)
func (c Time) AddYearNoOverflow() Time {
	return c.AddYearsNoOverflow(1)
}

// SubYears subtracts some years.
// N年前
func (c Time) SubYears(years int) Time {
	if c.IsInvalid() {
		return c
	}
	return c.AddYears(-years)
}

// SubYearsNoOverflow subtracts some years without overflowing month.
// N年前(月份不溢出)
func (c Time) SubYearsNoOverflow(years int) Time {
	return c.AddYearsNoOverflow(-years)
}

// SubYear subtracts one year.
// 1年前
func (c Time) SubYear() Time {
	return c.SubYears(1)
}

// SubYearNoOverflow subtracts one year without overflowing month.
// 1年前(月份不溢出)
func (c Time) SubYearNoOverflow() Time {
	return c.SubYearsNoOverflow(1)
}

// AddQuarters adds some quarters
// N个季度后
func (c Time) AddQuarters(quarters int) Time {
	return c.AddMonths(quarters * MonthsPerQuarter)
}

// AddQuartersNoOverflow adds quarters without overflowing month.
// N个季度后(月份不溢出)
func (c Time) AddQuartersNoOverflow(quarters int) Time {
	return c.AddMonthsNoOverflow(quarters * MonthsPerQuarter)
}

// AddQuarter adds one quarter
// 1个季度后
func (c Time) AddQuarter() Time {
	return c.AddQuarters(1)
}

// AddQuarterNoOverflow adds one quarter without overflowing month.
// 1个季度后(月份不溢出)
func (c Time) AddQuarterNoOverflow() Time {
	return c.AddQuartersNoOverflow(1)
}

// SubQuarters subtracts some quarters.
// N个季度前
func (c Time) SubQuarters(quarters int) Time {
	return c.AddQuarters(-quarters)
}

// SubQuartersNoOverflow subtracts some quarters without overflowing month.
// N个季度前(月份不溢出)
func (c Time) SubQuartersNoOverflow(quarters int) Time {
	return c.AddMonthsNoOverflow(-quarters * MonthsPerQuarter)
}

// SubQuarter subtracts one quarter.
// 1个季度前
func (c Time) SubQuarter() Time {
	return c.SubQuarters(1)
}

// SubQuarterNoOverflow subtracts one quarter without overflowing month.
// 1个季度前(月份不溢出)
func (c Time) SubQuarterNoOverflow() Time {
	return c.SubQuartersNoOverflow(1)
}

// AddMonths adds some months.
// N个月后
func (c Time) AddMonths(months int) Time {
	if c.IsInvalid() {
		return c
	}
	c.time = c.time.In(c.loc).AddDate(0, months, 0)
	return c
}

// AddMonthsNoOverflow adds some months without overflowing month.
// N个月后(月份不溢出)
func (c Time) AddMonthsNoOverflow(months int) Time {
	if c.IsInvalid() {
		return c
	}
	month := c.Month() + months
	// 获取N月后的最后一天
	last := time.Date(c.Year(), time.Month(month), 1, c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.loc).AddDate(0, 1, -1)
	day := c.Day()
	if c.Day() > last.Day() {
		day = last.Day()
	}
	c.time = time.Date(last.Year(), last.Month(), day, c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.loc)
	return c
}

// AddMonth adds one month.
// 1个月后
func (c Time) AddMonth() Time {
	return c.AddMonths(1)
}

// AddMonthNoOverflow adds one month without overflowing month.
// 1个月后(月份不溢出)
func (c Time) AddMonthNoOverflow() Time {
	return c.AddMonthsNoOverflow(1)
}

// SubMonths subtracts some months.
// N个月前
func (c Time) SubMonths(months int) Time {
	return c.AddMonths(-months)
}

// SubMonthsNoOverflow subtracts some months without overflowing month.
// N个月前(月份不溢出)
func (c Time) SubMonthsNoOverflow(months int) Time {
	return c.AddMonthsNoOverflow(-months)
}

// SubMonth subtracts one month.
// 1个月前
func (c Time) SubMonth() Time {
	return c.SubMonths(1)
}

// SubMonthNoOverflow subtracts one month without overflowing month.
// 1个月前(月份不溢出)
func (c Time) SubMonthNoOverflow() Time {
	return c.SubMonthsNoOverflow(1)
}

// AddWeeks adds some weeks.
// N周后
func (c Time) AddWeeks(weeks int) Time {
	return c.AddDays(weeks * DaysPerWeek)
}

// AddWeek adds one week.
// 1周后
func (c Time) AddWeek() Time {
	return c.AddWeeks(1)
}

// SubWeeks subtracts some weeks.
// N周前
func (c Time) SubWeeks(weeks int) Time {
	return c.SubDays(weeks * DaysPerWeek)
}

// SubWeek subtracts one week.
// 1周前
func (c Time) SubWeek() Time {
	return c.SubWeeks(1)
}

// AddDays adds some days.
// N天后
func (c Time) AddDays(days int) Time {
	if c.IsInvalid() {
		return c
	}
	c.time = c.time.In(c.loc).AddDate(0, 0, days)
	return c
}

// AddDay adds one day.
// 1天后
func (c Time) AddDay() Time {
	return c.AddDays(1)
}

// SubDays subtracts some days.
// N天前
func (c Time) SubDays(days int) Time {
	return c.AddDays(-days)
}

// SubDay subtracts one day.
// 1天前
func (c Time) SubDay() Time {
	return c.SubDays(1)
}

// AddHours adds some hours.
// N小时后
func (c Time) AddHours(hours int) Time {
	if c.IsInvalid() {
		return c
	}
	td := time.Duration(hours) * time.Hour
	c.time = c.time.In(c.loc).Add(td)
	return c
}

// AddHour adds one hour.
// 1小时后
func (c Time) AddHour() Time {
	return c.AddHours(1)
}

// SubHours subtracts some hours.
// N小时前
func (c Time) SubHours(hours int) Time {
	return c.AddHours(-hours)
}

// SubHour subtracts one hour.
// 1小时前
func (c Time) SubHour() Time {
	return c.SubHours(1)
}

// AddMinutes adds some minutes.
// N分钟后
func (c Time) AddMinutes(minutes int) Time {
	if c.IsInvalid() {
		return c
	}
	td := time.Duration(minutes) * time.Minute
	c.time = c.time.Add(td)
	return c
}

// AddMinute adds one minute.
// 1分钟后
func (c Time) AddMinute() Time {
	return c.AddMinutes(1)
}

// SubMinutes subtracts some minutes.
// N分钟前
func (c Time) SubMinutes(minutes int) Time {
	return c.AddMinutes(-minutes)
}

// SubMinute subtracts one minute.
// 1分钟前
func (c Time) SubMinute() Time {
	return c.SubMinutes(1)
}

// AddSeconds adds some seconds.
// N秒钟后
func (c Time) AddSeconds(seconds int) Time {
	if c.IsInvalid() {
		return c
	}
	td := time.Duration(seconds) * time.Second
	c.time = c.time.Add(td)
	return c
}

// AddSecond adds one second.
// 1秒钟后
func (c Time) AddSecond() Time {
	return c.AddSeconds(1)
}

// SubSeconds subtracts some seconds.
// N秒钟前
func (c Time) SubSeconds(seconds int) Time {
	return c.AddSeconds(-seconds)
}

// SubSecond subtracts one second.
// 1秒钟前
func (c Time) SubSecond() Time {
	return c.SubSeconds(1)
}
