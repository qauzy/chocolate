package xtime

import (
	"strconv"
	"strings"
	"time"
)

// Parse parses a standard string as a Time instance.
// 将标准格式时间字符串解析成 Time 实例
func (c Time) Parse(value string, timezone ...string) Time {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	layout := DateTimeFormat
	if value == "" || value == "0" || value == "0000-00-00 00:00:00" || value == "0000-00-00" || value == "00:00:00" {
		return c
	}
	if len(value) == 10 && strings.Count(value, "-") == 2 {
		layout = DateFormat
	}
	if strings.Index(value, "T") == 10 {
		layout = RFC3339Format
	}
	if _, err := strconv.ParseInt(value, 10, 64); err == nil {
		switch len(value) {
		case 8:
			layout = ShortDateFormat
		case 14:
			layout = ShortDateTimeFormat
		}
	}
	c = c.ParseByLayout(value, layout)
	if c.Error != nil {
		c.Error = invalidValueError(value)
		return c
	}
	return c
}

// Parse parses a standard string as a Time instance.
// 将标准时间字符串解析成 Time 实例
func Parse(value string, timezone ...string) Time {
	return NewTime().Parse(value, timezone...)
}

// ParseByFormat parses a string as a Time instance by format.
// 通过格式化字符将字符串解析成 xtime 实例
func (c Time) ParseByFormat(value string, format string, timezone ...string) Time {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	if value == "" || value == "0" || value == "0000-00-00 00:00:00" || value == "0000-00-00" || value == "00:00:00" {
		return c
	}
	c = c.ParseByLayout(value, format2layout(format))
	if c.Error != nil {
		c.Error = invalidFormatError(value, format)
		return c
	}
	return c
}

// ParseByFormat parses a string as a Time instance by format.
// 通过布局字符将字符串解析成 xtime 实例
func ParseByFormat(value string, format string, timezone ...string) Time {
	return NewTime().ParseByFormat(value, format, timezone...)
}

// ParseByLayout parses a string as a Time instance by layout.
// 通过布局字符将字符串解析成 xtime 实例
func (c Time) ParseByLayout(value string, layout string, timezone ...string) Time {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	if value == "" || value == "0" || value == "0000-00-00 00:00:00" || value == "0000-00-00" || value == "00:00:00" {
		return c
	}
	tt, err := time.ParseInLocation(layout, value, c.loc)
	if err != nil {
		c.Error = invalidLayoutError(value, layout)
		return c
	}
	c.time = tt
	return c
}

// ParseByLayout parses a string as a Time instance by layout.
// 将布局时间字符串解析成 Time 实例
func ParseByLayout(value string, layout string, timezone ...string) Time {
	return NewTime().ParseByLayout(value, layout, timezone...)
}
