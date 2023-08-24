package dateparser

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

func monthStringToNumber(s string) string {
	s = strings.ToUpper(s)
	m := map[string]string{
		"JAN": "01",
		"FEB": "02",
		"MAR": "03",
		"APR": "04",
		"MAY": "05",
		"JUN": "06",
		"JUL": "07",
		"AUG": "08",
		"SEP": "09",
		"OCT": "10",
		"NOV": "11",
		"DEC": "12",
	}
	return m[s]
}

func ParseDate(date string) (time.Time, error) {

	match, err := regexp.MatchString("[0-9]{4}-[0-9]{2}", date)
	if err != nil {
		return time.Time{}, err
	}
	if match {
		start, err := time.Parse("2006-01-02", date+"-01")
		if err != nil {
			return time.Time{}, err
		}
		return start, nil
	}

	match, err = regexp.MatchString("[0-9]{4} (JAN|FEB|MAR|APR|MAY|JUN|JUL|AUG|SEP|OCT|NOV|DEC)", date)
	if err != nil {
		return time.Time{}, err
	}
	if match {
		monthString := date[len(date)-3:]
		date = date[:len(date)-4] + "-" + monthStringToNumber(monthString) + "-01"
		start, err := time.Parse("2006-01-02", date)
		if err != nil {
			return time.Time{}, err
		}
		return start, nil

	}

	match, err = regexp.MatchString("[0-9]{2} (Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) [0-9]{2}", date)
	if err != nil {
		return time.Time{}, err
	}
	if match {
		dayString := date[:2]
		monthString := date[3:6]
		yearString := date[7:]
		yearInteger, err := strconv.Atoi(yearString)
		if err != nil {
			return time.Time{}, err
		}
		if yearInteger > 50 {
			yearString = "19" + yearString
		} else {
			yearString = "20" + yearString
		}
		date = yearString + "-" + monthStringToNumber(monthString) + "-" + dayString
		start, err := time.Parse("2006-01-02", date)
		if err != nil {
			return time.Time{}, err
		}
		return start, nil
	}

	return time.Time{}, nil
}
