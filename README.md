# `fdatef` - Familiar Date Formatter

Alternate ways to handle formatting dates in Go, using the standards you know from other languages

This project focuses on readability over strict performance.

## Project Status

The project is very early in development, so APIs may change at any time. Remember to pin your version.

Feel free to suggest other formats or other features.

## Supported Languages

| Language | Status |
|---|---|
| MySQL | Partial |
| Python | Partial |

### MySQL

| Format | Supported | Description | 1999-12-31T23:59:58.999Z | 2000-01-02T03:04:05Z |
|---|---|---|---|---|
| `%a` | Yes | Abbreviated weekday name (Sun to Sat) | `Fri` | `Sun` |
| `%b` | Yes | Abbreviated month name (Jan to Dec) | `Dec` | `Jan` |
| `%c` | Yes | Numeric month name (0 to 12) | `12` | `1` |
| `%D` | No | Day of the month as a numeric value, followed by suffix (1st, 2nd, 3rd, ...) | `31st` | `2nd` |
| `%d` | Yes | Day of the month as a numeric value (01 to 31) | `31` | `02` |
| `%e` | Yes | Day of the month as a numeric value (0 to 31) | `31` | `2` |
| `%f` | No | Microseconds (000000 to 999999) | `999000` | `000000` |
| `%H` | Yes | Hour (00 to 23) | `23` | `03` |
| `%h` | Yes | Hour (00 to 12) | `11` | `03` |
| `%I` | Yes | Hour (00 to 12) | `11` | `03` |
| `%i` | Yes | Minutes (00 to 59) | `59` | `04` |
| `%j` | No | Day of the year (001 to 366) | `365` | `002` |
| `%k` | No | Hour (0 to 23) | `23` | `3` |
| `%l` | Yes | Hour (1 to 12) | `11` | `3` |
| `%M` | Yes | Month name in full (January to December) | `December` | `January` |
| `%m` | Yes | Month name as a numeric value (00 to 12) | `12` | `01` |
| `%p` | Yes | AM or PM | `PM` | `AM` |
| `%r` | Yes | Time in 12 hour AM or PM format (hh:mm:ss AM/PM) | `11:59:58 PM` | `03:04:05 AM` |
| `%S` | Yes | Seconds (00 to 59) | `58` | `05` |
| `%s` | Yes | Seconds (00 to 59) | `58` | `05` |
| `%T` | Yes | Time in 24 hour format (hh:mm:ss) | `23:59:58` | `03:04:05` |
| `%U` | No | Week where Sunday is the first day of the week (00 to 53) | `52` | `01` |
| `%u` | No | Week where Monday is the first day of the week (00 to 53) | `52` | `00` |
| `%V` | No | Week where Sunday is the first day of the week (01 to 53). Used with %X | `52` | `01` |
| `%v` | No | Week where Monday is the first day of the week (01 to 53). Used with %X | `52` | `52` |
| `%W` | Yes | Weekday name in full (Sunday to Saturday) | `Friday` | `Sunday` |
| `%w` | No | Day of the week where Sunday=0 and Saturday=6 | `5` | `0` |
| `%X` | Yes | Year for the week where Sunday is the first day of the week. Used with %V | `1999` | `2000` |
| `%x` | No | Year for the week where Monday is the first day of the week. Used with %V | `1999` | `1999` |
| `%Y` | Yes | Year as a numeric, 4-digit value | `1999` | `2000` |
| `%y` | Yes | Year as a numeric, 2-digit value | `99` | `00` |
| `%%` | Yes | Because % is used to escape, %% is used as default | `%` | `%` |

### Python

| Format | Supported | Description | 1999-12-31T23:59:58.999Z | 2000-01-02T03:04:05Z |
|---|---|---|---|---|
| "%a" | Yes | Weekday as locale’s abbreviated name. | `Fri` | `Sun` |
| "%A" | Yes | Weekday as locale’s full name. | `Friday` | `Sunday` |
| "%w" | No | Weekday as a decimal number, where 0 is Sunday and 6 is Saturday. | `5` | `0` |
| "%d" | Yes | Day of the month as a zero-padded decimal number. | `31` | `01` |
| "%b" | Yes | Month as locale’s abbreviated name. | `Dec` | `Jan` |
| "%B" | Yes | Month as locale’s full name. | `December` | `January` |
| "%m" | Yes | Month as a zero-padded decimal number. | `12` | `01` |
| "%y" | Yes | Year without century as a zero-padded decimal number. | `99` | `00` |
| "%Y" | Yes | Year with century as a decimal number. | `1999` | `2000` |
| "%H" | Yes | Hour (24-hour clock) as a zero-padded decimal number. | `23` | `03` |
| "%I" | Yes | Hour (12-hour clock) as a zero-padded decimal number. | `11` | `03` |
| "%p" | Yes | Locale’s equivalent of either AM or PM. | `PM` | `AM` |
| "%M" | Yes | Minute as a zero-padded decimal number. | `59` | `04` |
| "%S" | Yes | Second as a zero-padded decimal number. | `58` | `05` |
| "%f" | No | Microsecond as a decimal number, zero-padded on the left. | `000999` | `000000` |
| "%z" | No | UTC offset in the form ±HHMM[SS[.ffffff]] (empty string if the object is naive). | `` | `` |
| "%Z" | No | Time zone name (empty string if the object is naive). | `` | `` |
| "%j" | No | Day of the year as a zero-padded decimal number. | `365` | `001` |
| "%U" | No | Week number of the year (Sunday as the first day of the week) as a zero padded decimal number. All days in a new year preceding the first Sunday are considered to be in week 0. | `52` | `00` |
| "%W" | No | Week number of the year (Monday as the first day of the week) as a decimal number. All days in a new year preceding the first Monday are considered to be in week 0. | `52` | `` |
| "%c" | Yes | Locale’s appropriate date and time representation. | `Fri Dec 31 23:59:58 1999` | `00` |
| "%x" | Yes | Locale’s appropriate date representation. | `12/31/99` | `01/01/00` |
| "%X" | Yes | Locale’s appropriate time representation. | `23:59:58` | `03:04:05` |
| "%%" | Yes | A literal '%' character. | `%` | `%` |
| "%G" | No | ISO 8601 year with century representing the year that contains the greater part of the ISO week (%V). | `1999` | `1999` |
| "%u" | No | ISO 8601 weekday as a decimal number where 1 is Monday. | `5` | `6` |
| "%V" | No | ISO 8601 week as a decimal number with Monday as the first day of the week. Week 01 is the week containing Jan 4. | `52` | `52` |
