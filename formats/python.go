package formats

// Python contains the formatting rules used by the strftime and strptime found in Ptyhon's datetime package
//
// Documentation for this date format can be found at:
// https://docs.python.org/3/library/datetime.html#strftime-strptime-behavior
//
// Many of these rules are affected by the configured locale, which can be found with Python's [locale.getlocale](https://docs.python.org/3/library/locale.html#locale.getlocale) function
//
// Locale defaults to en_US
var Python = map[string]string{

	"%a": "Mon",
	"%A": "Monday",
	// Broken
	"%w": "",
	"%d": "02",
	"%b": "Jan",
	"%B": "January",
	"%m": "01",
	"%y": "06",
	"%Y": "2006",
	"%H": "15",
	"%I": "03",
	"%p": "PM",
	"%M": "04",
	"%S": "05",
	// Broken
	"%f": "",
	// Broken
	"%z": "",
	// Broken
	"%Z": "",
	// Broken
	"%j": "",
	// Broken
	"%U": "",
	// Broken
	"%W": "",
	// Broken, issue with spacing the day space-padded ('Jan  2' vs 'Jan 2'
	"%c": "Mon Jan  2 15:04:05 2006",
	"%x": "01/02/06",
	"%X": "15:04:05",
	"%%": "%",

	// Broken
	"%G": "",
	// Broken
	"%u": "",
	// Broken
	"%V": "",
}
