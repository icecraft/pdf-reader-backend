package svc

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.shlab.tech/xurui/pdf-reader-backend/pkg/types"
)

func TestEs(t *testing.T) {
	o := &CibaEsIndex{cli: testEsCli, indexName: "ciba_translate"}
	var docId string

	t.Run("put", func(t *testing.T) {
		t1 := types.RetrieveItem{Word: "Bee", EN: "insect", CN: "蜜蜂", Examples: []string{"a bee fly above the flower"}}
		err := o.Put(&t1)
		assert.NoError(t, err)
	})

	t.Run("search", func(t *testing.T) {
		resp, id, err := o.RetrieveWordTrans("Bee")
		assert.NoError(t, err)
		assert.Equal(t, resp.Word, "Bee")
		assert.Equal(t, resp.EN, "insect")
		assert.Equal(t, len(resp.Examples), 1)

		docId = id
	})

	t.Run("incrHit", func(t *testing.T) {
		err := o.IncrHit(docId)
		assert.NoError(t, err)
	})
}
