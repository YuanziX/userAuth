package utils

import (
	"fmt"
	"strings"
	"time"
)

// Takes date of format DD-MM-YYYY or DD/MM/YYYY and converts it to time.Time
func StringDateToTimeObject(date string) (time.Time, error) {
	var layout string
	if strings.Contains(date, "-") {
		layout = "02-01-2006"
	} else if strings.Contains(date, "/") {
		layout = "02/01/2006"
	} else {
		return time.Time{}, fmt.Errorf("invalid date format")
	}

	t, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
