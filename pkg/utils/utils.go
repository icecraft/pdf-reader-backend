package utils

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	AK_LENGTH = 20
	SK_LENGTH = 40
)


func MayCreateDir(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		if err = os.MkdirAll(path, 0777); err != nil {
			return err
		}
		return nil
	}
	return err
}

func IsDirExisted(path string) (bool, error) {
	stats, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	if stats.IsDir() {
		return true, nil
	}
	return false, nil
}

func SampleByRatio(ids []int, maxCount int) []int {
	rand.Seed(time.Now().UTC().UnixNano())

	maxCount = MinInt(len(ids), maxCount)
	ret := make([]int, 0)
	l := len(ids)
	var tmp int

	for i := 0; maxCount > i; i++ {
		j := rand.Intn(l)
		ret = append(ret, ids[j])
		//swap data to tail
		tmp = ids[l-1]
		ids[l-1] = ids[j]
		ids[j] = tmp
		l = l - 1
	}
	return ret
}

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

func IterStreamigIO(r io.Reader, proc func(int, string) error) error {
	lineNo := 1
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if err := proc(lineNo, scanner.Text()); err != nil {
			return err
		}
		lineNo += 1
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func ParseConversationID(idStr string) (int64, error) {
	if idStr == "" {
		return 0, nil
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid conversation id: %s", idStr)
	}
	return id, nil
}

func SnakeCase2CamelCase(v string) string {
	strs := strings.Split(v, "_")
	for i := range strs {
		strs[i] = strings.Title(strs[i])
	}
	return strings.Join(strs, "")
}

func StrMillSeconds(millSeconds int64) string {
	t := time.Unix(millSeconds/1000, 0)
	return t.Format(time.RFC3339)
}

func IterSha256(r io.Reader) (string, error) {
	h := sha256.New()
	buf := make([]byte, 1024*1024)
	for {
		count, err := r.Read(buf)
		if err != nil {
			if err != io.EOF {
				return "", err
			}
			break
		}
		h.Write(buf[:count])
	}

	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs), nil
}

func Sha256(r io.Reader) (string, error) {
	h := sha256.New()
	io.Copy(h, r)
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs), nil
}

func InStrArray(str string, arr []string) bool {
	for _, v := range arr {
		if str == v {
			return true
		}
	}
	return false
}

func BriefString(s string) string {
	if 15 >= len(s) {
		return s
	}
	return s[:6] + "..." + s[len(s)-6:]
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
