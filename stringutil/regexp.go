package stringutil

import (
	"regexp"
	"strings"
)

func regexpSubstrings(expr, s, exp_head, exp_end string, n int) (r []string) {
	strs := regexp.MustCompile(expr).FindAllString(s, -1)
	for i, str := range strs {
		if n > 0 && i >= n {
			break
		}
		var start, end int = 0, 0
		if len(exp_head) != 0 {
			start = strings.Index(str, exp_head) + len(exp_head)
		}
		if len(exp_end) != 0 {
			end = strings.LastIndex(str, exp_end)
		}
		if str = Substring(str, start, end); len(str) != 0 {
			r = append(r, str)
		}
	}
	return
}

func RegexpSubstrings(expr, s, exp_head, exp_end string) []string {
	return regexpSubstrings(expr, s, exp_head, exp_end, -1)
}

func RegexpSubstring(expr, s, exp_head, exp_end string) string {
	if strs := regexpSubstrings(expr, s, exp_head, exp_end, 1); len(strs) != 0 {
		return strs[0]
	}
	return ""
}
