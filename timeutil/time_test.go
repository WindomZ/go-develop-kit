package timeutil

import "testing"

func TestFirstTimeToday(t *testing.T) {
	r := FirstTimeToday()
	if r.Hour() != 0 {
		t.Fatal("Error hour", r.Hour())
	} else if r.Minute() != 0 {
		t.Fatal("Error minute", r.Minute())
	} else if r.Second() != 0 {
		t.Fatal("Error second", r.Second())
	} else if r.Nanosecond() != 0 {
		t.Fatal("Error nanosecond", r.Nanosecond())
	}
}

func TestLastTimeToday(t *testing.T) {
	r := LastTimeToday()
	if r.Hour() != 23 {
		t.Fatal("Error hour", r.Hour())
	} else if r.Minute() != 59 {
		t.Fatal("Error minute", r.Minute())
	} else if r.Second() != 59 {
		t.Fatal("Error second", r.Second())
	} else if r.Nanosecond() != 999999999 {
		t.Fatal("Error nanosecond", r.Nanosecond())
	}
}

func TestToday(t *testing.T) {
	r := Today(1, 2, 3, 4)
	if r.Hour() != 1 {
		t.Fatal("Error hour", r.Hour())
	} else if r.Minute() != 2 {
		t.Fatal("Error minute", r.Minute())
	} else if r.Second() != 3 {
		t.Fatal("Error second", r.Second())
	} else if r.Nanosecond() != 4 {
		t.Fatal("Error nanosecond", r.Nanosecond())
	}
}
