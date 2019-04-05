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
| `%r` | Yes | Time in 12 hour AM or PM format (hh:mm:ss AM/PM) | 11:59:58 PM | `03:04:05 AM` |
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

