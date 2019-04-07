package formats

// MySQL contains the formatting rules used by the MySQL DATE_FORMAT function
//
// Documentation for this date format can be found at:
// https://dev.mysql.com/doc/refman/8.0/en/date-and-time-functions.html#function_date-format
var MySQL = map[string]string{

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
