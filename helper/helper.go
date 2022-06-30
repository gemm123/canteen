package helper

import (
	"fmt"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func FormatDate(date time.Time) string {
	year, month, day := date.Date()
	return fmt.Sprintf("%d/%d/%d", year, month, day)
}

func Sum(sumSliceInt []int) int {
	var result int
	for _, i := range sumSliceInt {
		result += i
	}
	return result
}

func ConvertSliceStringToInteger(sumSliceString []string) []int {
	var sumSliceInt []int
	for _, i := range sumSliceString {
		is, _ := strconv.Atoi(i)
		sumSliceInt = append(sumSliceInt, int(is))
	}
	return sumSliceInt
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
