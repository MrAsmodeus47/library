package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hour, minute, second int
}

type TimeParseError struct {
	msg   string
	input string
}

func (t TimeParseError) Error() string {
	return fmt.Sprintf("%v : %v", t.msg, t.input)
}
func PareTime(input string) (Time, error) {
	components := strings.Split(input, ":")
	if len(components) != 3 {
		return Time{}, &TimeParseError{"invalide number of time component", input}
	} else {
		hour, err := strconv.Atoi(components[0])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing hour : %v", err), input}
		}
		minute, err := strconv.Atoi(components[1])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing minute : %v", err), input}
		}
		second, err := strconv.Atoi(components[2])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing second : %v", err), input}
		}
		if hour > 23 || hour < 0 {
			return Time{}, &TimeParseError{"hour out of range :", fmt.Sprintf("%v", hour)}
		}
		if minute > 60 || minute < 0 {
			return Time{}, &TimeParseError{"hour out of range :", fmt.Sprintf("%v", minute)}
		}
		if second > 60 || second < 0 {
			return Time{}, &TimeParseError{"hour out of range :", fmt.Sprintf("%v", second)}
		}
		return Time{hour, minute, second}, nil
	}
}
