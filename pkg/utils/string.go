package utils

import (
	"regexp"
)

func UniqueString(strList []string) []string {
	m := map[string]bool{}
	ret := []string{}
	for _, str := range strList {
		if _, ok := m[str]; !ok {
			m[str] = true
			ret = append(ret, str)
		}
	}
	return ret
}

func ExtractContent(content, key, leftSep, rightSep string) []string {
	if len(key) == 0 {
		return nil
	}
	pattern := regexp.MustCompile(key)

	ret := make([]string, 0)

	for _, indic := range pattern.FindAllStringIndex(content, -1) {
		index := indic[0]

		var leftPos, rightPos int
		count := 0
		// found leftSep
		for pos := index - 1; pos >= 0; pos-- {
			if string(content[pos]) == leftSep {
				count++
			}
			if string(content[pos]) == rightSep {
				count--
			}
			if count == 1 {
				leftPos = pos
				break
			}
		}
		if count != 1 {
			continue
		}

		count = 0
		// found rightSep
		for pos := index + len(key); len(content) > pos; pos++ {
			if string(content[pos]) == leftSep {
				count--
			}
			if string(content[pos]) == rightSep {
				count++
			}
			if count == 1 {
				rightPos = pos
				break
			}
		}
		if count != 1 {
			continue
		}
		ret = append(ret, content[leftPos:rightPos+1])
	}
	return ret
}
