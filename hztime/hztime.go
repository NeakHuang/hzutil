// Date: 2023/4/10
// Author:
// Description：

package hztime

import "time"

const (
	Second = 1
	Minute = 60 * Second // 一分钟的秒数
	Hour   = 60 * Minute // 一小时的秒数
	Day    = 24 * Hour   // 一天的秒数
	Week   = 7 * Day     // 一周的秒数
)

// HourBase
// The current hour starting from the specified date and time.
// 指定日期时间的当前小时开始
func HourBase(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, t.Hour(), 0, 0, 0, t.Location())
}

// HourBaseUnix
// The current hour starting from the specified timestamp.
// 指定时间戳的当前小时开始
func HourBaseUnix(ts int64) time.Time {
	return HourBase(time.Unix(ts, 0))
}

// Day0Clock
// 0 o'clock of the specified date and time
// 指定日期时间的0点
func Day0Clock(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// TsDay0Clock
// 0 o'clock of the specified timestamp
// 指定时间戳的0点
func TsDay0Clock(ts int64) time.Time {
	return Day0Clock(time.Unix(ts, 0))
}

// Today0Clock
// 0 o'clock of today
// 今日0点
func Today0Clock() time.Time {
	return Day0Clock(time.Now())
}

// TodayHourBase
// Specified hour of the today
// 今日指定小时时间
func TodayHourBase(h int) time.Time {
	t := time.Now()
	year, month, day := t.Date()
	return time.Date(year, month, day, h, 0, 0, 0, t.Location())
}

// Monday0Clock
// Monday at 0:00 specified by the timestamp.
// 指定时间戳所在的周一0点
func Monday0Clock(t time.Time) time.Time {
	weekDay := int64(t.Weekday())
	if weekDay == 0 {
		weekDay = 7
	}
	ts := Day0Clock(t).Unix() - (weekDay-1)*Day
	return time.Unix(ts, 0)
}

// TsMonday0Clock
// Monday at 0:00 of the specified date and time.
// 指定日期时间所在的周一0点
func TsMonday0Clock(ts int64) time.Time {
	return Monday0Clock(time.Unix(ts, 0))
}

// NowMonday0Clock
// Get the timestamp of 0:00 on Monday of the current week.
// 当前时间的周一0点
func NowMonday0Clock() time.Time {
	t := time.Now()
	weekDay := int64(t.Weekday())
	if weekDay == 0 {
		weekDay = 7
	}
	ts := Day0Clock(t).Unix() - (weekDay-1)*Day
	return time.Unix(ts, 0)
}

// Month1st0Clock
// 1st day of the corresponding month at 0:00 specified time.
// 指定日期时间对应月份的1号0点
func Month1st0Clock(t time.Time) time.Time {
	year, month, _ := t.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, t.Location())
}

// TsMonth1st0Clock
// 1st day of the corresponding month at 0:00 specified time.
// 指定时间戳对应月份的1号0点
func TsMonth1st0Clock(ts int64) time.Time {
	t := time.Unix(ts, 0)
	return Month1st0Clock(t)
}

// NowMonth1st0Clock
// The 1st day of the current month at 0:00.
// 当前时间的本月1号0点
func NowMonth1st0Clock() time.Time {
	t := time.Now()
	return Month1st0Clock(t)
}

// UponMonth1st0Clock
// 1st day of the corresponding month at 0:00 specified time.
// 指定日期时间对应月份的1号0点
func UponMonth1st0Clock(t time.Time) time.Time {
	year, month, _ := t.Date()
	return time.Date(year, month-1, 1, 0, 0, 0, 0, t.Location())
}

// TsUponMonth1st0Clock
// 1st day of the corresponding month at 0:00 specified time.
// 指定时间戳对应月份的1号0点
func TsUponMonth1st0Clock(ts int64) time.Time {
	t := time.Unix(ts, 0)
	return UponMonth1st0Clock(t)
}

// NowUponMonth1st0Clock
// The 1st day of the current month at 0:00.
// 当前时间的本月1号0点
func NowUponMonth1st0Clock() time.Time {
	t := time.Now()
	return UponMonth1st0Clock(t)
}

// 0:day, 1:week, 2:month
// 默认返回今日0点
func TsCycle0Clock(cycle int32, ts int64) uint64 {
	switch cycle {
	case 0: // 指定时间0点
		return uint64(TsDay0Clock(ts).Unix())
	case 1: // 指定时间所在周一0点
		return uint64(TsMonday0Clock(ts).Unix())
	case 2: // 指定时间所在月份1号0点
		return uint64(TsMonth1st0Clock(ts).Unix())
	}
	return uint64(TsDay0Clock(ts).Unix())
}

// 0:day, 1:week, 2:month
// 默认返回今日0点
func TsUponCycle0Clock(cycle int32, ts int64) uint64 {
	switch cycle {
	case 0: // 指定时间0点
		return TsCycle0Clock(cycle, ts) - Day
	case 1: // 指定时间所在周一0点
		return TsCycle0Clock(cycle, ts) - Week
	case 2: // 指定时间所在月份1号0点
		return uint64(TsUponMonth1st0Clock(ts).Unix())
	}
	return TsCycle0Clock(cycle, ts) - Day
}
