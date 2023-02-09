package utils

func MinInt(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func RemovePrefix(s, prefix string) string {
	return s[len(prefix):]
}

func InStrArray(str string, arr []string) bool {
	for _, v := range arr {
		if str == v {
			return true
		}
	}
	return false
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
