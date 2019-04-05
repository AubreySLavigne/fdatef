package fdatef

import (
	"strings"
	"time"
)

// Format returns a string representing the time, in the provided format.
//
// This uses the Data Formatting method found in MySQL:
// https://dev.mysql.com/doc/refman/8.0/en/date-and-time-functions.html#function_date-format
func Format(t time.Time, format string) string {

	rules := map[string]string{
		"%a": "Mon",
		"%b": "Jan",
		"%c": "1",
		// Broken
		"%D": "2st",
		"%d": "02",
		"%e": "2",
		// Broken
		"%f": "000000",
		"%H": "15",
		"%h": "03",
		"%I": "03",
		"%i": "04",
		// Broken
		"%j": "003",
		// Broken
		"%k": "5",
		"%l": "3",
		"%M": "January",
		"%m": "01",
		"%p": "PM",
		"%r": "03:04:05 PM",
		"%S": "05",
		"%s": "05",
		"%T": "15:04:05",
		// Broken
		"%U": "",
		// Broken
		"%u": "",
		// Broken
		"%V": "",
		// Broken
		"%v": "",
		"%W": "Monday",
		// Broken
		"%w": "",
		"%X": "2006",
		// Broken
		"%x": "2006",
		"%Y": "2006",
		"%y": "06",
		"%%": "%",
	}

	for k, v := range rules {
		format = strings.Replace(format, k, v, -1)
	}

	return t.Format(format)
}
