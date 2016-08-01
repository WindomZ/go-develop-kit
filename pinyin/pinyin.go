package pinyin

import (
	"strings"
)

var (
	pyTableLen = len(pyValue)
)

type Pinyin struct {
	Split string
	Upper bool
}

func New() Pinyin {
	return Pinyin{
		Split: " ",
		Upper: true,
	}
}

func (c *Pinyin) Convert(s string) (string, error) {
	var pyStr string

	gbk, err := Utf8ToGbk(s)
	if err != nil {
		return pyStr, err
	}

	pyArr := c.GbkToPinyin(gbk)
	pyStr = strings.Join(pyArr, c.Split)

	return pyStr, nil
}

func (c *Pinyin) GbkToPinyin(gbk string) []string {
	var pyStr []string

	for i := 0; i < len(gbk); i++ {
		p := int(gbk[i])
		if p > 0 && p < 160 {
			pyStr = append(pyStr, string(gbk[i]))
		} else {
			i++
			if i >= len(gbk) {
				break
			}
			q := int(gbk[i])
			p = p*256 + q - 65536

			py := pinyinSearch(p)
			if py != "" {
				if c.Upper == false {
					py = strings.ToLower(py)
				}
				pyStr = append(pyStr, py)
			}
		}
	}
	return pyStr
}

func Utf8ToGbk(s string) (string, error) {
	cd, err := Open("gbk", "utf-8")
	if err != nil {
		return "", err
	}
	defer cd.Close()
	gbk := cd.ConvString(s)
	return gbk, nil
}

func pinyinSearch(p int) string {
	if v, ok := tableMap[p]; ok {
		return v
	} else {
		for i := pyTableLen - 1; i >= 0; i-- {
			if pyValue[i] <= p {
				return pyName[i]
			}
		}
	}
	return ""
}
