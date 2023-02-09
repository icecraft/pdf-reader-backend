package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestSubS struct {
	ID   int64  `json:"id,omitempty" bson:"id,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}

type TestS struct {
	ID   int64     `json:"id,omitempty" bson:"id,omitempty"`
	SI   int64     `json:"si,omitempty" bson:"si,omitempty"`
	SII  int       `json:"sii,omitempty" bson:"sii,omitempty"`
	SI32 int32     `json:"si_32,omitempty" bson:"si_32,omitempty"`
	SS   string    `json:"ss,omitempty" bson:"ss,omitempty"`
	SF   float64   `json:"sf,omitempty" bson:"sf,omitempty"`
	VI   []int64   `json:"vi,omitempty" bson:"vi,omitempty"`
	VII  []int     `json:"vii,omitempty" bson:"vii,omitempty"`
	VI32 []int32   `json:"vi_32,omitempty" bson:"vi_32,omitempty"`
	VF   []float64 `json:"vf,omitempty" bson:"vf,omitempty"`
	VS   []string  `json:"vs,omitempty" bson:"vs,omitempty"`

	SubStruct TestSubS `json:"sub_struct,omitempty" bson:"sub_struct,omitempty"`

	M  map[string]interface{} `json:"m,omitempty" bson:"m,omitempty"`
	NM map[string]interface{} `json:"mm,omitempty" bson:"nm,omitempty"`
}

func TestReflectSetField(t *testing.T) {
	t1 := TestS{ID: 1, SII: 2, SI32: 3, SS: "hello", SF: 10.6,
		VI: []int64{1, 2}, VII: []int{3, 4}, VI32: []int32{5, 6}, VF: []float64{2.5, 1.25}, VS: []string{"tom", "hanks"}}

	t.Run("set int | set int32 | set int64 | set float64 ｜ set string", func(t *testing.T) {
		t2 := t1

		ReflectSetField(&t2, "SI", 101)
		assert.Equal(t, t2.SI, int64(101))

		ReflectSetField(&t2, "SII", 102)
		assert.Equal(t, t2.SII, 102)

		ReflectSetField(&t2, "SI32", 103)
		assert.Equal(t, t2.SI32, int32(103))

		ReflectSetField(&t2, "SS", "gogo")
		assert.Equal(t, t2.SS, "gogo")

		ReflectSetField(&t2, "SF", 0.125)
		assert.Equal(t, t2.SF, float64(0.125))
	})

	t.Run("set non-existed field", func(t *testing.T) {
		t3 := t1
		err := ReflectSetField(&t3, "XXXIII", 101)
		assert.True(t, errors.Is(err, ErrNoFieldFound))
	})

}
