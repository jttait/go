package dateparser

import (
	"regexp"
	"time"
)

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
	match, err = regexp.MatchString("[0-9]{4} [JAN|FEB|MAR|APR|MAY|JUN|JUL|AUG|SEP|OCT|NOV|DEC]", date)
	if err != nil {
		return time.Time{}, err
	}
	if match {
		monthString := date[len(date)-3:]
		monthToInteger := map[string]string{"JAN": "01", "FEB": "02", "MAR": "03", "APR": "04", "MAY": "05", "JUN": "06", "JUL": "07", "AUG": "08", "SEP": "09", "OCT": "10", "NOV": "11", "DEC": "12"}
		date = date[:len(date)-4] + "-" + monthToInteger[monthString] + "-01"
		start, err := time.Parse("2006-01-02", date)
		if err != nil {
			return time.Time{}, nil
		}
		return start, nil

	}
	return time.Time{}, nil
}
