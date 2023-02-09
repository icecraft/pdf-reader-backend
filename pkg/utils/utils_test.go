package utils

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	//go:embed testdata/4k.dat
	fourk []byte

	//go:embed testdata/1k.dat
	onek []byte

	//go:embed testdata/0.5k.dat
	halfk []byte

	//go:embed testdata/less.1.dat
	less1k []byte

	//go:embed testdata/less.2.dat
	less2k []byte
)

func TestSampleRatio(t *testing.T) {
	proc := func(arr []int) map[int]bool {
		ret := make(map[int]bool)
		for _, v := range arr {
			ret[v] = true
		}
		return ret
	}

	ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	ret := SampleByRatio(ids, 10)
	assert.Equal(t, len(ret), 10)
	retM := proc(ret)
	assert.Equal(t, len(retM), 10)

	ret = SampleByRatio(ids, 11)
	assert.Equal(t, len(ret), 11)

	retM = proc(ret)
	assert.Equal(t, len(retM), 11)
}

func TestSnakeCase2CamelCase(t *testing.T) {
	t1 := ""
	t2 := "hello"
	t3 := "hello_world"
	t4 := "_"
	assert.Equal(t, SnakeCase2CamelCase(t1), "")
	assert.Equal(t, SnakeCase2CamelCase(t2), "Hello")
	assert.Equal(t, SnakeCase2CamelCase(t3), "HelloWorld")
	assert.Equal(t, SnakeCase2CamelCase(t4), "")
}

func TestIterSha256(t *testing.T) {

	t.Run("4k", func(t *testing.T) {
		h, err := IterSha256(bytes.NewReader(fourk))
		assert.NoError(t, err)
		assert.Equal(t, h, "732f4322ef0710e1a8b1b498e6aaaafca7bf9b426ad9e3290c69c9d575261425")
	})

	t.Run("1k", func(t *testing.T) {
		h, err := IterSha256(bytes.NewReader(onek))
		assert.NoError(t, err)
		assert.Equal(t, h, "93c6cf65aec3474b79f5006970b28a584bbb2717f24bcdfeda1c7ea0342ea0a8")
	})

	t.Run("0.5k", func(t *testing.T) {
		h, err := IterSha256(bytes.NewReader(halfk))
		assert.NoError(t, err)
		assert.Equal(t, h, "a6aec426362ba28054d2c74a0cce98cb38c82ed4b03122a909e51914aa6d5802")
	})

	t.Run("less.1", func(t *testing.T) {
		h, err := IterSha256(bytes.NewReader(less1k))
		assert.NoError(t, err)
		assert.Equal(t, h, "aa4c0aaccfe692d64ece7b14086c9344f3254797d0d421348062fb536c9a7684")
	})

	t.Run("less.2", func(t *testing.T) {
		h, err := IterSha256(bytes.NewReader(less2k))
		assert.NoError(t, err)
		assert.Equal(t, h, "3f8d6c64cf591897f4377faa29f76047c74f8d5e7f66343bf6c5171d5a5e4d26")

		h2, err := Sha256(bytes.NewReader(less2k))
		assert.NoError(t, err)
		assert.Equal(t, h, h2)
	})

	t.Run("less.2 not iter", func(t *testing.T) {
		h, err := Sha256(bytes.NewReader(less2k))
		assert.NoError(t, err)
		assert.Equal(t, h, "3f8d6c64cf591897f4377faa29f76047c74f8d5e7f66343bf6c5171d5a5e4d26")
	})
}

func TestBriefString(t *testing.T) {
	t1 := "1c8d4e5c8c02cc8532501be928768e757dcda2beb23ccc82d15527ac5d8d9087"
	assert.Equal(t, BriefString(t1), "1c8d4e...8d9087")
}
