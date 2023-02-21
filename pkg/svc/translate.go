package svc

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/icecraft/pdf-reader-backend/pkg/types"
)

/*

	Inspired by https://blog.csdn.net/KYuruyan/article/details/115272212
*/

var (
	cibaApiEndpoint  = "http://dict.iciba.com/dictionary/word/query/web"
	cibaParams       = []string{"client", "key", "timestamp", "word"}
	cibaSignMagicNum = "7ece94d9f9c202b0d2ec557dg4r9bc"
	headers          = map[string]string{
		"Origin":     "http://www.iciba.com",
		"Referer":    "http://www.iciba.com/",
		"User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
	}
)

type Ciba struct {
	paramValues []string
	params      map[string]string
}

func (o *Ciba) GetSignature(params map[string]string) string {
	if len(o.paramValues) == 0 {
		o.paramValues = make([]string, len(cibaParams))
	}
	for i, key := range cibaParams {
		o.paramValues[i] = params[key]
	}

	str := "/dictionary/word/query/web" + strings.Join(o.paramValues, "") + cibaSignMagicNum
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func (o *Ciba) Fetch(word string) (*types.CibaResp, error) {
	if len(o.params) == 0 {
		o.params = make(map[string]string)
	}
	delete(o.params, "signature")

	o.params["client"] = "6"
	o.params["key"] = "1000006"
	o.params["timestamp"] = fmt.Sprintf("%d", time.Now().UnixMilli())
	o.params["word"] = word
	o.params["signature"] = o.GetSignature(o.params)

	url := cibaApiEndpoint + "?" + ConvertQueryMapToQuerystring(o.params)

	var ret types.CibaResp

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	json.Unmarshal(data, &ret)
	return &ret, nil
}

func ConvertQueryMapToQuerystring(m map[string]string) string {
	q := url.Values{}
	for key := range m {
		q.Add(key, m[key])
	}
	return q.Encode()
}
