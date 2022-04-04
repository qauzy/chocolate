package LocalDateTime

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/qauzy/chocolate/timex/LocalDate"
	"time"
)

type LocalDateTime time.Time

const (
	Nanosecond  time.Duration = 1
	Microsecond               = 1000 * Nanosecond
	Millisecond               = 1000 * Microsecond
	Second                    = 1000 * Millisecond
	Minute                    = 60 * Second
	Hour                      = 60 * Minute
)
const (
	dateTimeFormart = "2006-01-02 15:04:05"
	BeiJinAreaTime  = "Asia/Shanghai"
	TimeBarFormat   = "2006-01-02 15:04:05"
	TimeMMDD        = "01-02"
)

func GetFirstDateOfWeek(current *LocalDate.LocalDate) *LocalDate.LocalDate {
	now := time.Time(*current)
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekMonday := LocalDate.LocalDate(weekStartDate)
	return &weekMonday
}

//星期一为第一天 == 0
func GetDayIndexOfWeek(current *LocalDate.LocalDate) int {
	now := time.Time(*current)
	w := now.Weekday() - 1
	if w < 0 {
		w = w + 7
	}

	return int(w)
}

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

func (t *LocalDateTime) UnmarshalJSON(data []byte) (err error) {
	if data == nil || len(data) == 2 {
		return
	}
	now, err := time.ParseInLocation(`"`+dateTimeFormart+`"`, string(data), time.Local)
	*t = LocalDateTime(now)
	return
}

func (t LocalDateTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(dateTimeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, dateTimeFormart)
	b = append(b, '"')
	return b, nil
}

func (t LocalDateTime) Value() (driver.Value, error) {
	// MyTime 转换成 timex.Time 类型
	tTime := time.Time(t)
	return tTime.Format(dateTimeFormart), nil
}

func (t *LocalDateTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case string:
		// 字符串转成 timex.Time 类型
		tTime, _ := time.ParseInLocation(dateTimeFormart, vt, GetBjTimeLoc())
		*t = LocalDateTime(tTime)
	case []byte:
		tTime, _ := time.ParseInLocation(dateTimeFormart, string(vt), GetBjTimeLoc())
		*t = LocalDateTime(tTime)
	case time.Time:
		*t = LocalDateTime(vt)
	default:
		return errors.New(fmt.Sprintf("类型处理错误,%v", v))
	}
	return nil
}
func (t LocalDateTime) After(t2 LocalDateTime) bool {
	// MyTime 转换成 timex.Time 类型
	time1 := time.Time(t)
	time2 := time.Time(t2)
	return time1.After(time2)
}

func (t LocalDateTime) Before(t2 LocalDateTime) bool {
	// MyTime 转换成 timex.Time 类型
	time1 := time.Time(t)
	time2 := time.Time(t2)
	return time1.Before(time2)
}

func (t LocalDateTime) Equal(t2 LocalDateTime) bool {
	// MyTime 转换成 timex.Time 类型
	time1 := time.Time(t)
	time2 := time.Time(t2)
	return time1.Equal(time2)
}
func (t LocalDateTime) ToLocalDate() *LocalDate.LocalDate {
	// MyTime 转换成 timex.Time 类型
	tTime := LocalDate.LocalDate(t)
	return &tTime
}

func (t *LocalDateTime) PlusDays(d int) (t2 *LocalDateTime) {
	// MyTime 转换成 timex.Time 类型
	time1 := time.Time(*t).Add(time.Duration(d) * time.Hour * 24)
	time2 := LocalDateTime(time1)
	return &time2
}
func (t *LocalDateTime) MinusDays(d int) (t2 *LocalDateTime) {
	// MyTime 转换成 timex.Time 类型
	time1 := time.Time(*t).Add(-time.Duration(d) * time.Hour * 24)
	time2 := LocalDateTime(time1)
	return &time2
}

func (t LocalDateTime) Format(fmt string) string {
	// MyTime 转换成 timex.Time 类型
	tTime := time.Time(t)
	return tTime.Format(fmt)
}

func Now() LocalDateTime {
	// MyTime 转换成 timex.Time 类型
	t := time.Now()
	return LocalDateTime(t)
}

func Of(ld *LocalDate.LocalDate, lt time.Time) *LocalDateTime {
	// MyTime 转换成 timex.Time 类型
	t := time.Date(time.Time(*ld).Year(), time.Time(*ld).Month(), time.Time(*ld).Day(), time.Time(lt).Hour(), time.Time(lt).Minute(), time.Time(lt).Second(), time.Time(lt).Nanosecond(), GetBjTimeLoc())
	ldnew := LocalDateTime(t)
	return &ldnew
}
