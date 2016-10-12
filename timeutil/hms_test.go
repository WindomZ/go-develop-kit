package timeutil

import (
	"testing"
	"time"
)

func TestParseTodayHHMMSS(t *testing.T) {
	_time, err := ParseTodayHHMMSS("150405", "000159")
	if err != nil {
		t.Fatal(err)
	}
	_now := time.Now()
	if _time.Year() != _now.Year() ||
		_time.Month() != _now.Month() ||
		_time.Day() != _now.Day() ||
		_time.Hour() != 00 ||
		_time.Minute() != 1 ||
		_time.Second() != 59 {
		t.Fatal(_time.String(), _now.String())
	}
}

func TestDefaultParseTodayHHMMSS(t *testing.T) {
	_time, err := DefaultParseTodayHHMMSS("00:01:59")
	if err != nil {
		t.Fatal(err)
	}
	_now := time.Now()
	if _time.Year() != _now.Year() ||
		_time.Month() != _now.Month() ||
		_time.Day() != _now.Day() ||
		_time.Hour() != 00 ||
		_time.Minute() != 1 ||
		_time.Second() != 59 {
		t.Fatal(_time.String(), _now.String())
	}
}

func TestParseHHMMSS(t *testing.T) {
	_time, err := ParseHHMMSS("150405", "235959")
	if err != nil {
		t.Fatal(err)
	}
	_now := time.Now().AddDate(0, 0, -1)
	if _time.Year() != _now.Year() ||
		_time.Month() != _now.Month() ||
		_time.Day() != _now.Day() ||
		_time.Hour() != 23 ||
		_time.Minute() != 59 ||
		_time.Second() != 59 {
		t.Fatal(_time.String(), _now.String())
	}
}

func TestDefaultParseHHMMSS(t *testing.T) {
	_time, err := DefaultParseHHMMSS("23:59:59")
	if err != nil {
		t.Fatal(err)
	}
	_now := time.Now().AddDate(0, 0, -1)
	if _time.Year() != _now.Year() ||
		_time.Month() != _now.Month() ||
		_time.Day() != _now.Day() ||
		_time.Hour() != 23 ||
		_time.Minute() != 59 ||
		_time.Second() != 59 {
		t.Fatal(_time.String(), _now.String())
	}
}
