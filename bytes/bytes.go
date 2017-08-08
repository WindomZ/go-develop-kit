// Copyright (c) 2017, WindomZ & Contributors
// All rights reserved.

package bytes

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
	EB
)

// Format formats bytes integer to human readable string.
// For example, 31323 bytes will return 30.59KB.
func Format(b uint64) string {
	var unit string
	value := float64(b)

	switch {
	case b < KB:
		return strconv.FormatUint(b, 10) + "B"
	case b < MB:
		value /= KB
		unit = "KB"
	case b < GB:
		value /= MB
		unit = "MB"
	case b < TB:
		value /= GB
		unit = "GB"
	case b < PB:
		value /= TB
		unit = "TB"
	case b < EB:
		value /= PB
		unit = "PB"
	default:
		return strconv.FormatUint(b, 10) + "B"
	}

	return fmt.Sprintf("%.02f%s", value, unit)
}

var pattern = regexp.MustCompile(`(?i)^(-?[0-9.]+)([KMGTP]B?|B)$`)

// Parse parses human readable bytes string to bytes integer.
// For example, 6GB (6G is also valid) will return 6442450944.
func Parse(s string) (uint64, error) {
	parts := pattern.FindStringSubmatch(s)
	if len(parts) <= 2 {
		return 0, fmt.Errorf("error parsing: %s", s)
	}

	f, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return 0, err
	}

	switch strings.ToUpper(parts[2]) {
	case "B":
		f *= B
	case "K", "KB":
		f *= KB
	case "M", "MB":
		f *= MB
	case "G", "GB":
		f *= GB
	case "T", "TB":
		f *= TB
	case "P", "PB":
		f *= PB
	}

	return uint64(f), nil
}

// MustParse same as Parse, if fail to parses then panic a error.
func MustParse(s string) uint64 {
	b, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return b
}
