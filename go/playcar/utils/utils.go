package utils

import "time"

func GetCurrentDay() string {
	return time.Now().Format(time.DateOnly)
}

func GetTotalWeeksOfYear() int {
	currentTime := time.Now()
	startOfYear := time.Date(currentTime.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
	weeks := int(currentTime.Sub(startOfYear).Hours() / (24 * 7))
	return weeks + 1 // Adding 1 to include the current week
}

func GetCurrentMonth() int {
	return int(time.Now().Month())
}

func GetCurrentYear() int {
	return time.Now().Year()
}
