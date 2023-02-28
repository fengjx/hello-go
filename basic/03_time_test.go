package basic

import (
	"testing"
	"time"
)

func StringToTime(str string) (time.Time, error) {
	layout := "2006-01-02"
	return time.Parse(layout, str)
}

func GetWeek(datetime string) (y, w int) {
	timeLayout := "20060102"
	loc, _ := time.LoadLocation("Local")
	tmp, _ := time.ParseInLocation(timeLayout, datetime, loc)
	return tmp.ISOWeek()
}

func TestGetWeek(t *testing.T) {
	t.Log(GetWeek("20230217"))
	d, err := StringToTime("2023-02-20")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(d.ISOWeek())
}

//获取某一天的0点时间
func GetZeroTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

//获取本周周一的日期
func GetDayOfWeek(t time.Time, weekday time.Weekday) time.Time {
	t0 := GetZeroTime(t)
	if t.Weekday() == weekday {
		offset := int(t.Weekday() - weekday)
		t0 = t0.AddDate(0, 0, -offset)
	} else {
		offset := int(weekday - t.Weekday())
		if offset > 0 {
			offset = -6
		}
		t0 = t0.AddDate(0, 0, offset)
	}
	return t0
}

func TestWeek(t *testing.T) {
	now := time.Now()
	t.Log(now.Weekday())
	t.Log(GetDayOfWeek(now, time.Tuesday))
}
