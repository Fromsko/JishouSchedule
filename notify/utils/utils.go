package utils

import (
	"strconv"
	"time"
)

// GetWeekly 获取当前星期
func GetWeekly() (weekly string) {
	now := time.Now()

	// 获取当前是星期几
	weekday := now.Weekday()

	// 定义一个星期几到中文的映射
	weekdayMap := map[time.Weekday]string{
		time.Monday:    "星期一",
		time.Tuesday:   "星期二",
		time.Wednesday: "星期三",
		time.Thursday:  "星期四",
		time.Friday:    "星期五",
		time.Saturday:  "星期六",
		time.Sunday:    "星期日",
	}

	// 根据映射获取中文表示
	return weekdayMap[weekday]
}

// GetWeek 获取当前第几周
func GetWeek(startMon int) string {
	_, weeks := time.Now().ISOWeek()
	return strconv.Itoa(weeks - startMon)
}
