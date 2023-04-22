package basic

import (
	"testing"
	"time"
)

const (
	dayLayout = "2006-01-02"
)

// GetDayOfWeek3 获取本周对应星期的日期
func GetDayOfWeek3(t time.Time, day time.Weekday, loc *time.Location) string {
	t = t.In(loc)
	// 刚好是今天
	if t.Weekday() == day {
		return t.Format(dayLayout)
	}
	d := int(day)
	if d == 0 {
		d = 7
	}
	today := int(t.Weekday())
	offset := d - today
	return t.AddDate(0, 0, offset).Format(dayLayout)
}

// GetEndOfWeekTime 获取一周的结束时间（周日 23:59:59）
func GetEndOfWeekTime(t time.Time, loc *time.Location) time.Time {
	t = t.In(loc)
	offset := 0
	// 刚好是今天
	if t.Weekday() != time.Sunday {
		today := int(t.Weekday())
		offset = 7 - today
	}
	t = t.AddDate(0, 0, offset)
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 1e9-1, loc)
}

func TestTime(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Riyadh")
	locCN, _ := time.LoadLocation("Asia/Shanghai")
	date := time.Date(2023, 3, 6, 5, 59, 59, 0, locCN)

	t.Log(GetDayOfWeek3(date, time.Sunday, loc))

	weekTime := GetEndOfWeekTime(date, loc)
	t.Log(weekTime)

}
