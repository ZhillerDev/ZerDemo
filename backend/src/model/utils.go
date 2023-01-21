package model

import "time"

func Unix2Time(timetamp int) string {
	t := time.Unix(int64(timetamp), 0)
	return t.Format("2006-01-02 15:04:05")
}
