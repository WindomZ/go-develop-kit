package stringutil

import (
	"regexp"
	"strings"
)

func RegexpSubstring(expr, s, exp_head, exp_end string) (r []string) {
	strs := regexp.MustCompile(expr).FindAllString(s, -1)
	for _, str := range strs {
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
