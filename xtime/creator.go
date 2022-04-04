package xtime

import (
	"strconv"
	"time"
)

// CreateFromTimestamp creates a Time instance from a given timestamp, second, millisecond, microsecond and nanosecond are supported.
// 从给定的时间戳创建 Time 实例，支持秒、毫秒、微秒和纳秒
func (c Time) CreateFromTimestamp(timestamp int64, timezone ...string) Time {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	ts, count := timestamp, len(strconv.FormatInt(timestamp, 10))
	if timestamp < 0 {
		count -= 1
	}
	switch count {
	case 10:
		ts = timestamp
	case 13:
		ts = timestamp / 1e3
	case 16:
		ts = timestamp / 1e6
	case 19:
		ts = timestamp / 1e9
	}
	c.time = time.Unix(ts, 0)
	return c
}

// CreateFromTimestamp creates a Time instance from a given timestamp.
// 从给定的时间戳创建 Time 实例
func CreateFromTimestamp(timestamp int64, timezone ...string) Time {
	return NewTime().CreateFromTimestamp(timestamp, timezone...)
}

// CreateFromDateTime creates a Time instance from a given date and time.
// 从给定的年月日时分秒创建 Time 实例
func (c Time) CreateFromDateTime(year int, month int, day int, hour int, minute int, second int, timezone ...string) Time {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	c.time = time.Date(year, time.Month(month), day, hour, minute, second, time.Now().Nanosecond(), c.loc)
	return c
}

// CreateFromDateTime creates a Time instance from a given date and time.
// 从给定的年月日时分秒创建 Time 实例
func CreateFromDateTime(year int, month int, day int, hour int, minute int, second int, timezone ...string) Time {
	return NewTime().CreateFromDateTime(year, month, day, hour, minute, second, timezone...)
}

// CreateFromDate creates a Time instance from a given date.
// 从给定的年月日创建 Time 实例
func (c Time) CreateFromDate(year int, month int, day int, timezone ...string) Time {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	hour, minute, second := time.Now().In(c.loc).Clock()
	c.time = time.Date(year, time.Month(month), day, hour, minute, second, time.Now().Nanosecond(), c.loc)
	return c
}

// CreateFromDate creates a Time instance from a given date.
// 从给定的年月日创建 Time 实例
func CreateFromDate(year int, month int, day int, timezone ...string) Time {
	return NewTime().CreateFromDate(year, month, day, timezone...)
}

// CreateFromTime creates a Time instance from a given time.
// 从给定的时分秒创建 Time 实例
func (c Time) CreateFromTime(hour int, minute int, second int, timezone ...string) Time {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	year, month, day := time.Now().In(c.loc).Date()
	c.time = time.Date(year, month, day, hour, minute, second, time.Now().Nanosecond(), c.loc)
	return c
}

// CreateFromTime creates a Time instance from a given time.
// 从给定的时分秒创建 Time 实例
func CreateFromTime(hour int, minute int, second int, timezone ...string) Time {
	return NewTime().CreateFromTime(hour, minute, second, timezone...)
}
