package fdatef

import (
	"strings"
	"time"

	"github.com/AubreySLavigne/fdatef/formats"
)

// Format returns a string representing the time t, in the provided format.
//
// Format defaults to MySQL type formatting
func Format(t time.Time, format string) string {
	defaultRules := formats.MySQL
	return FormatWithRules(t, format, defaultRules)
}

// FormatWithRules is a generic formatting function.
//
// It accepts a map[string]string defining search and replace rules that
// define the expected formatting.
func FormatWithRules(t time.Time, format string, rules map[string]string) string {
	for k, v := range rules {
		format = strings.Replace(format, k, v, -1)
	}

	return t.Format(format)
}
