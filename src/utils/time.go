package utils

import (
	"strings"
	"time"
)

func GetCurTimeStr() string {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	return formattedTime
}

func GetCurTime() time.Time {
	parse, _ := time.Parse("2006-01-02 15:04:05", GetCurTimeStr())
	return parse
}

func Time2Str(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func Str2Time(str string) time.Time {
	var layout string
	if strings.Index(str, "T") > -1 {
		layout = time.RFC3339
	} else {
		layout = "2006-01-02 15:04:05"
	}
	t, _ := time.Parse(layout, str)
	return t
}
