package timeutil

import "time"

func FirstTime(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

func LastTime(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 23, 59, 59, 999999999, time.Local)
}

func FirstTimeToday() time.Time {
	return FirstTime(time.Now().Date())
}

func LastTimeToday() time.Time {
	return LastTime(time.Now().Date())
}

func Today(hour, min, sec, nsec int) time.Time {
	year, month, day := time.Now().Date()
	return time.Date(year, month, day, hour, min, sec, nsec, time.Local)
}

func FirstTimeOffsetDays(days int) time.Time {
	return FirstTime(time.Now().AddDate(0, 0, days).Date())
}

func LastTimeOffsetDays(days int) time.Time {
	return LastTime(time.Now().AddDate(0, 0, days).Date())
}

func FirstTimeTomorrow() time.Time {
	return FirstTimeOffsetDays(1)
}

func LastTimeTomorrow() time.Time {
	return LastTimeOffsetDays(1)
}

func FirstTimeYesterday() time.Time {
	return FirstTimeOffsetDays(-1)
}

func LastTimeYesterday() time.Time {
	return LastTimeOffsetDays(-1)
}
