package LocalDate

import (
	"database/sql/driver"
	"errors"
	"github.com/jinzhu/now"
	"time"
)

type LocalDate time.Time

const (
	dateFormart      = "2006-01-02"
	dateFormartShort = "20060102"
	DTF_YMD          = "2006-01-02"
	DTF_YM           = "2006_01"
	BeiJinAreaTime   = "Asia/Shanghai"
)

func GetBjTimeLoc() *time.Location {
	// 获取北京时间, 在 windows系统上 timex.LoadLocation 会加载失败, 最好的办法是用 timex.FixedZone
	var bjLoc *time.Location
	var err error
	bjLoc, err = time.LoadLocation(BeiJinAreaTime)
	if err != nil {
		bjLoc = time.FixedZone("CST", 8*3600)
	}

	return bjLoc
}

func Now() *LocalDate {
	l := LocalDate(time.Now())
	return &l
}
func (t *LocalDate) UnmarshalJSON(data []byte) (err error) {
	if data == nil || len(data) == 2 {
		return
	}

	now, err := time.ParseInLocation(`"`+dateFormart+`"`, string(data), time.Local)
	if err != nil {
		now, err = time.ParseInLocation(`"`+dateFormartShort+`"`, string(data), time.Local)
	}
	*t = LocalDate(now)
	return
}

func (t LocalDate) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(dateFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, dateFormart)
	b = append(b, '"')
	return b, nil
}

func (t LocalDate) Value() (driver.Value, error) {
	// MyTime 转换成 timex.Time 类型
	tTime := time.Time(t)
	return tTime.Format(dateFormart), nil
}

func (t *LocalDate) Scan(v interface{}) error {
	switch vt := v.(type) {
	case string:
		// 字符串转成 timex.Time 类型
		tTime, err := time.Parse(dateFormart, vt)
		if err != nil {
			tTime, _ = time.Parse(dateFormartShort, vt)
		}
		*t = LocalDate(tTime)
	case []byte:
		tTime, err := time.Parse(dateFormart, string(vt))
		if err != nil {
			tTime, _ = time.Parse(dateFormartShort, string(vt))
		}
		*t = LocalDate(tTime)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *LocalDate) IsBefore(t2 *LocalDate) bool {
	// MyTime 转换成 timex.Time 类型
	tt0 := time.Time(*t)
	tTime := time.Date(tt0.Year(), tt0.Month(), tt0.Day(), 0, 0, 0, 0, time.Local)
	tt2 := time.Time(*t2)
	tTime2 := time.Date(tt2.Year(), tt2.Month(), tt2.Day(), 0, 0, 0, 0, time.Local)
	return tTime.Before(tTime2)
}

func (t *LocalDate) IsAfter(t2 *LocalDate) bool {
	// MyTime 转换成 timex.Time 类型
	tt0 := time.Time(*t)
	tTime := time.Date(tt0.Year(), tt0.Month(), tt0.Day(), 0, 0, 0, 0, time.Local)
	tt2 := time.Time(*t2)
	tTime2 := time.Date(tt2.Year(), tt2.Month(), tt2.Day(), 0, 0, 0, 0, time.Local)
	return tTime.After(tTime2)
}

func (t *LocalDate) IsEqual(t2 *LocalDate) bool {
	// MyTime 转换成 timex.Time 类型
	tt0 := time.Time(*t)
	tTime := time.Date(tt0.Year(), tt0.Month(), tt0.Day(), 0, 0, 0, 0, time.Local)
	tt2 := time.Time(*t2)
	tTime2 := time.Date(tt2.Year(), tt2.Month(), tt2.Day(), 0, 0, 0, 0, time.Local)
	return tTime.Equal(tTime2)
}

func (t *LocalDate) MinusDays(d int) (t2 *LocalDate) {
	// MyTime 转换成 timex.Time 类型
	time1 := time.Time(*t).Add(-time.Duration(d) * time.Hour * 24)
	time2 := LocalDate(time1)
	return &time2
}
func (t *LocalDate) MinusMonths(m int) (t2 *LocalDate) {
	// MyTime 转换成 timex.Time 类型
	t0 := time.Time(*t)
	time1 := AddDate(t0, 0, -1)
	time2 := LocalDate(time1)
	return &time2
}

func (t *LocalDate) PlusDays(d int) (t2 *LocalDate) {
	// MyTime 转换成 timex.Time 类型
	time1 := time.Time(*t).Add(time.Duration(d) * time.Hour * 24)
	time2 := LocalDate(time1)
	return &time2
}
func (t *LocalDate) ToEpochDay() (days int) {
	// MyTime 转换成 timex.Time 类型
	d := time.Date(time.Time(*t).Year(), time.Time(*t).Month(), time.Time(*t).Day(), 0, 0, 0, 0, time.Local).Unix() / 86400

	return int(d)
}

func (t LocalDate) Format(fmt string) string {
	// MyTime 转换成 timex.Time 类型
	tTime := time.Time(t)
	return tTime.Format(fmt)
}
func (t LocalDate) String() string {
	// MyTime 转换成 timex.Time 类型
	return t.Format(dateFormart)
}
func Parse(str string, f string) *LocalDate {
	t, _ := time.ParseInLocation(f, str, GetBjTimeLoc())
	lt := LocalDate(t)
	return &lt
}

func ConvertStringToLocalDate(src string, f string) *LocalDate {
	t, _ := time.ParseInLocation(f, src, GetBjTimeLoc())
	lt := LocalDate(t)
	return &lt
}

// AddDate 时间增减
// 类似于ruby中的时间增减，和 time.AddDate 不同
// 如：
// loc, _ := time.LoadLocation("Asia/Shanghai")
// t := time.Date(2010, 3, 31, 12, 0, 0, 0, loc)
// utils.AddDate(t, 0, 1)
// => 2010-04-30 12:00:00 +0800 CST
// 不会因为4月没有31号，而变成5月1号
func AddDate(t time.Time, years int, months int) time.Time {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()

	// firstDayOfMonthAfterAddDate: years 年，months 月后的 那个月份的1号
	firstDayOfMonthAfterAddDate := time.Date(year+years, month+time.Month(months), 1,
		hour, min, sec, t.Nanosecond(), t.Location())
	// firstDayOfMonthAfterAddDate 月份的最后一天
	lastDay := now.New(firstDayOfMonthAfterAddDate).EndOfMonth().Day()

	// 如果 t 的天 > lastDay，则设置为lastDay
	// 如：t 为 2020-03-31 12:00:00 +0800，增加1个月，为4月31号
	// 但是4月没有31号，则设置为4月最后一天lastDay（30号）
	if day > lastDay {
		day = lastDay
	}

	return time.Date(year+years, month+time.Month(months), day,
		hour, min, sec, t.Nanosecond(), t.Location())
}
