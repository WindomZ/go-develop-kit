package timeutil

import "time"

// ParseTodayHHMMSS parses a formatted string and returns the time value it represents.
// The layout  defines the format by showing how the reference time,
// defined to be
//	15:04:05
// would be interpreted if it were the value; it serves as an example of
// the input format. The same interpretation will then be made to the
// input string.
// ParseTodayHHMMSS returns a time in Local`s UTC, and force define year, month and day
// same as today.
func ParseTodayHHMMSS(layout, value string) (time.Time, error) {
	t, err := time.ParseInLocation(layout, value, time.Local)
	if err != nil {
		return t, err
	}
	year, month, day := time.Now().Date()
	t = time.Date(year, month, day,
		t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	return t, nil
}

// DefaultParseTodayHHMMSS is like ParseTodayHHMMSS but differs in one way,
// define the format as
//	15:04:05
func DefaultParseTodayHHMMSS(value string) (time.Time, error) {
	return ParseTodayHHMMSS("15:04:05", value)
}

// ParseHHMMSS is like ParseTodayHHMMSS but differs in one important way,
// returns the time default today if before now, else define yesterday.
func ParseHHMMSS(layout, value string) (time.Time, error) {
	t, err := ParseTodayHHMMSS(layout, value)
	if err != nil {
		return t, err
	} else if t.After(time.Now()) {
		t = t.AddDate(0, 0, -1)
	}
	return t, nil
}

// DefaultParseHHMMSS is like ParseHHMMSS but differs in one way,
// define the format as
//	15:04:05
func DefaultParseHHMMSS(value string) (time.Time, error) {
	return ParseHHMMSS("15:04:05", value)
}
