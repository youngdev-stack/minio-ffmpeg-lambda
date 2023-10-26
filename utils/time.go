package utils

import (
	"fmt"
	"time"
)

// FormatTime 格式化时间
// 如果时间超过 12 个月,显示年月日
// 如果时间超过 24 小时，显示天
// 如果时间超过 60 分钟，显示小时
// 如果时间超过 60 秒，显示分钟
func FormatTime(t time.Time) string {
	now := time.Now()
	duration := now.Sub(t)

	if duration >= 365*24*time.Hour {
		return fmt.Sprintf("%d年%d天", int(duration.Hours()/24/365), int(duration.Hours()/24)%365)
	} else if duration >= 24*time.Hour {
		return fmt.Sprintf("%d天", int(duration.Hours()/24))
	} else if duration >= time.Hour {
		return fmt.Sprintf("%d小时", int(duration.Minutes()/60))
	} else {
		return fmt.Sprintf("%d分钟", int(duration.Seconds()/60))
	}
}
