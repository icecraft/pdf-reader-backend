package svc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslateCiba(t *testing.T) {

	o := &Ciba{}

	t.Run("signature", func(t *testing.T) {
		params := make(map[string]string)

		params["client"] = "6"
		params["key"] = "1000006"
		params["timestamp"] = "1676809728411"
		params["word"] = "bee"

		signature := o.GetSignature(params)
		assert.Equal(t, signature, "aa6bdc77c5fc05ea47417f461d3be9ed")
	})

}

func TestConvertQueryMap(t *testing.T) {
	param := make(map[string]string)
	param["a"] = "1"
	param["b"] = "2"

	res := ConvertQueryMapToQuerystring(param)
	assert.Equal(t, res, "a=1&b=2")
}
