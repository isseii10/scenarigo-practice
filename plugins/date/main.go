package main

import "time"

var Layout = "2006-01-02"

func TodayIn(s string) (string, error) {
	loc, err := time.LoadLocation(s)
	if err != nil {
		return "", err
	}
	return time.Now().In(loc).Format(Layout), nil
}