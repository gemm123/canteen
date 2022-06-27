package helper

import (
	"fmt"
	"time"
)

func FormatDate(date time.Time) string {
	year, month, day := date.Date()
	return fmt.Sprintf("%d/%d/%d", year, month, day)
}
