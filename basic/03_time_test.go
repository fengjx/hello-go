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
	// loc, _ := time.LoadLocation("Asia/Riyadh")
	loc, _ := time.LoadLocation("Asia/Shanghai")

	date := time.Date(2023, 3, 5, 5, 0, 0, 0, loc)
	// date, _ = time.Parse("2006-01-02 15:04:05", "2023-03-05 06:00:00")
	y, w := date.ISOWeek()
	t.Log(y, w, date)

	y, w = date.In(loc).ISOWeek()
	t.Log(y, w, date)
	now := time.Now()
	t.Log(now.Weekday())
	t.Log(GetDayOfWeek(now, time.Tuesday))
}

func GetDayOfWeek2(t time.Time, day time.Weekday, fmtStr string) (dayStr string) {
	if t.Weekday() == day {
		//修改hour、min、sec = 0后格式化
		dayStr = t.Format(fmtStr)
	} else {
		d := int(day)
		if d == 0 {
			d = 7
		}
		today := int(t.Weekday())
		offset := d - today
		dayStr = t.AddDate(0, 0, offset).Format(fmtStr)
	}
	return
}

//func GetLastWeekSunday(t time.Time, fmtStr string) (day string, err error) {
//	monday := GetMondayOfWeek(t, fmtStr)
//	dayObj, err := time.Parse(fmtStr, monday)
//	if err != nil {
//		return
//	}
//	day = dayObj.AddDate(0, 0, -1).Format(fmtStr)
//	return
//}

func TestWeek2(t *testing.T) {
	//loc, _ := time.LoadLocation("Asia/Riyadh")
	locCN, _ := time.LoadLocation("Asia/Shanghai")
	//
	date := time.Date(2023, 3, 2, 4, 59, 59, 0, locCN)
	//
	//date2 := date.In(loc)

	t.Log(GetDayOfWeek2(date, time.Friday, "2006-01-02"))

}
