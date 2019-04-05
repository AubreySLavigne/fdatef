package fdatef

import (
	"fmt"
	"testing"
	"time"
)

func TestMySQLFormat(t *testing.T) {
	tests := []struct {
		date     string
		format   string
		expected string
	}{

		// Abbreviated weekday name (Sun to Sat)
		{"1999-12-31T23:59:58.999Z", "%a", "Fri"},
		{"2000-01-02T03:04:05Z", "%a", "Sun"},

		// Abbreviated month name (Jan to Dec)
		{"1999-12-31T23:59:58.999Z", "%b", "Dec"},
		{"2000-01-02T03:04:05Z", "%b", "Jan"},

		// Numeric month name (0 to 12)
		{"1999-12-31T23:59:58.999Z", "%c", "12"},
		{"2000-01-02T03:04:05Z", "%c", "1"},

		// // Day of the month as a numeric value, followed by suffix (1st, 2nd, 3rd, ...)
		// {"1999-12-31T23:59:58.999Z", "%D", "31st"},
		// {"2000-01-02T03:04:05Z", "%D", "2nd"},

		// Day of the month as a numeric value (01 to 31)
		{"1999-12-31T23:59:58.999Z", "%d", "31"},
		{"2000-01-02T03:04:05Z", "%d", "02"},

		// Day of the month as a numeric value (0 to 31)
		{"1999-12-31T23:59:58.999Z", "%e", "31"},
		{"2000-01-02T03:04:05Z", "%e", "2"},

		// // Microseconds (000000 to 999999)
		// {"1999-12-31T23:59:58.999Z", "%f", "999000"},
		// {"2000-01-02T03:04:05Z", "%f", "000000"},

		// Hour (00 to 23)
		{"1999-12-31T23:59:58.999Z", "%H", "23"},
		{"2000-01-02T03:04:05Z", "%H", "03"},

		// Hour (00 to 12)
		{"1999-12-31T23:59:58.999Z", "%h", "11"},
		{"2000-01-02T03:04:05Z", "%h", "03"},

		// Hour (00 to 12)
		{"1999-12-31T23:59:58.999Z", "%I", "11"},
		{"2000-01-02T03:04:05Z", "%I", "03"},

		// Minutes (00 to 59)
		{"1999-12-31T23:59:58.999Z", "%i", "59"},
		{"2000-01-02T03:04:05Z", "%i", "04"},

		// // Day of the year (001 to 366)
		// {"1999-12-31T23:59:58.999Z", "%j", "365"},
		// {"2000-01-02T03:04:05Z", "%j", "002"},

		// // Hour (0 to 23)
		// {"1999-12-31T23:59:58.999Z", "%k", "23"},
		// {"2000-01-02T03:04:05Z", "%k", "3"},

		// Hour (1 to 12)
		{"1999-12-31T23:59:58.999Z", "%l", "11"},
		{"2000-01-02T03:04:05Z", "%l", "3"},

		// Month name in full (January to December)
		{"1999-12-31T23:59:58.999Z", "%M", "December"},
		{"2000-01-02T03:04:05Z", "%M", "January"},

		// Month name as a numeric value (00 to 12)
		{"1999-12-31T23:59:58.999Z", "%m", "12"},
		{"2000-01-02T03:04:05Z", "%m", "01"},

		// AM or PM
		{"1999-12-31T23:59:58.999Z", "%p", "PM"},
		{"2000-01-02T03:04:05Z", "%p", "AM"},

		// Time in 12 hour AM or PM format (hh:mm:ss AM/PM)
		{"1999-12-31T23:59:58.999Z", "%r", "11:59:58 PM"},
		{"2000-01-02T03:04:05Z", "%r", "03:04:05 AM"},

		// Seconds (00 to 59)
		{"1999-12-31T23:59:58.999Z", "%S", "58"},
		{"2000-01-02T03:04:05Z", "%S", "05"},

		// Seconds (00 to 59)
		{"1999-12-31T23:59:58.999Z", "%s", "58"},
		{"2000-01-02T03:04:05Z", "%s", "05"},

		// Time in 24 hour format (hh:mm:ss)
		{"1999-12-31T23:59:58.999Z", "%T", "23:59:58"},
		{"2000-01-02T03:04:05Z", "%T", "03:04:05"},

		// // Week where Sunday is the first day of the week (00 to 53)
		// {"1999-12-31T23:59:58.999Z", "%U", "52"},
		// {"2000-01-02T03:04:05Z", "%U", "01"},

		// // Week where Monday is the first day of the week (00 to 53)
		// {"1999-12-31T23:59:58.999Z", "%u", "52"},
		// {"2000-01-02T03:04:05Z", "%u", "00"},

		// // Week where Sunday is the first day of the week (01 to 53). Used with %X
		// {"1999-12-31T23:59:58.999Z", "%V", "52"},
		// {"2000-01-02T03:04:05Z", "%V", "01"},

		// // Week where Monday is the first day of the week (01 to 53). Used with %X
		// // TODO: Both of these values will be the same, so we need to create a third to test this out
		// {"1999-12-31T23:59:58.999Z", "%v", "52"},
		// {"2000-01-02T03:04:05Z", "%v", "52"},

		// Weekday name in full (Sunday to Saturday)
		{"1999-12-31T23:59:58.999Z", "%W", "Friday"},
		{"2000-01-02T03:04:05Z", "%W", "Sunday"},

		// // Day of the week where Sunday=0 and Saturday=6
		// {"1999-12-31T23:59:58.999Z", "%w", "5"},
		// {"2000-01-02T03:04:05Z", "%w", "0"},

		// Year for the week where Sunday is the first day of the week. Used with %V
		{"1999-12-31T23:59:58.999Z", "%X", "1999"},
		{"2000-01-02T03:04:05Z", "%X", "2000"},

		// // Year for the week where Monday is the first day of the week. Used with %V
		// // TODO: Both of these values will be the same, so we need to create a third to test this out
		// {"1999-12-31T23:59:58.999Z", "%x", "1999"},
		// {"2000-01-02T03:04:05Z", "%x", "1999"},

		// Year as a numeric, 4-digit value
		{"1999-12-31T23:59:58.999Z", "%Y", "1999"},
		{"2000-01-02T03:04:05Z", "%Y", "2000"},

		// Year as a numeric, 2-digit value
		{"1999-12-31T23:59:58.999Z", "%y", "99"},
		{"2000-01-02T03:04:05Z", "%y", "00"},

		// Because % is used to escape, %% is used as default
		{"1999-12-31T23:59:58.999Z", "%%", "%"},
		{"2000-01-02T03:04:05Z", "%%", "%"},
	}

	for _, test := range tests {
		date, err := time.Parse(time.RFC3339Nano, test.date)
		if err != nil {
			t.Fatalf("%s", err)
		}
		if res := Format(date, test.format); res != test.expected {
			t.Errorf("Formatting date %-26v with format %v returned incorrect result. Got %8v, Expected %8v",
				fmt.Sprintf("'%s'", test.date),
				fmt.Sprintf("'%s'", test.format),
				fmt.Sprintf("'%s'", res),
				fmt.Sprintf("'%s'", test.expected))
		}
	}
}
