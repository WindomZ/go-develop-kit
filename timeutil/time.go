package timeutil

import "time"

// FirstTime return first time(00:00:00) with year, month and day
func FirstTime(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

// LastTime return last time(23:59:59.999999999) with year, month and day
func LastTime(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 23, 59, 59, 999999999, time.Local)
}

// FirstTimeToday return first time(00:00:00) of today
func FirstTimeToday() time.Time {
	return FirstTime(time.Now().Date())
}

// LastTimeToday return last time(23:59:59.999999999) of today
func LastTimeToday() time.Time {
	return LastTime(time.Now().Date())
}

// Today return time with year, month and day of today, and with hour, min, sec and nsec
func Today(hour, min, sec, nsec int) time.Time {
	year, month, day := time.Now().Date()
	return time.Date(year, month, day, hour, min, sec, nsec, time.Local)
}

// FirstTimeOffsetDays like as FirstTimeToday but offset days
func FirstTimeOffsetDays(days int) time.Time {
	return FirstTime(time.Now().AddDate(0, 0, days).Date())
}

// LastTimeOffsetDays like as LastTimeToday but offset days
func LastTimeOffsetDays(days int) time.Time {
	return LastTime(time.Now().AddDate(0, 0, days).Date())
}

// FirstTimeToday return first time(00:00:00) of tomorrow
func FirstTimeTomorrow() time.Time {
	return FirstTimeOffsetDays(1)
}

// LastTimeToday return last time(23:59:59.999999999) of tomorrow
func LastTimeTomorrow() time.Time {
	return LastTimeOffsetDays(1)
}

// FirstTimeToday return first time(00:00:00) of yesterday
func FirstTimeYesterday() time.Time {
	return FirstTimeOffsetDays(-1)
}

// LastTimeToday return last time(23:59:59.999999999) of yesterday
func LastTimeYesterday() time.Time {
	return LastTimeOffsetDays(-1)
}
