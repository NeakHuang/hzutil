// Date: 2023/6/3
// Author:
// Description：

package hztime

import "time"

const (
	GoTimeFmtBaseSimple = "20060102"
	GoTimeFmtBase       = "2006-01-02 15:04:05"
	GoTimeFmtBase1      = "2006/01/02 15:04:05"
)

// 解析日期文本("20060102")
func DateFmt(dateStr string) time.Time {
	parseTime, err := time.ParseInLocation(GoTimeFmtBaseSimple, dateStr, time.Now().Location())
	if err != nil {
		return time.Now()
	}
	return parseTime
}

// 解析时间文本("2006-01-02 15:04:05")
func DatetimeFmt(dateStr string) time.Time {
	parseTime, err := time.ParseInLocation(GoTimeFmtBase, dateStr, time.Now().Location())
	if err != nil {
		return time.Now()
	}
	return parseTime
}

// 解析时间文本("2006/01/02 15:04:05")
func DatetimeFmt1(dateStr string) time.Time {
	parseTime, err := time.ParseInLocation(GoTimeFmtBase1, dateStr, time.Now().Location())
	if err != nil {
		return time.Now()
	}
	return parseTime
}
