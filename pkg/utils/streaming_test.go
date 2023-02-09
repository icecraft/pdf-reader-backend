package utils

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	mrand "math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStreamingRead(t *testing.T) {

	t.Run("equal | content ", func(t *testing.T) {
		// 对于 128 size 大小的 buff， 循环 10000 次随机读取 1-128 大小的数据。保证数据的完整性
		count := 0
		accSrcBuf := make([]byte, 0)
		destSrcBuf := new(bytes.Buffer)

		src := func() ([]byte, error) {

			if count == 1000000 {
				return nil, io.EOF
			}
			count += 1

			n := mrand.Intn(128) + 1
			buf := make([]byte, n)
			nr, err := rand.Read(buf)
			if err != nil {
				return nil, err
			}
			if nr != n {
				return nil, fmt.Errorf("tried to read %d bytes, but got %d bytes", n, nr)
			}
			accSrcBuf = append(accSrcBuf, buf...)
			return buf, nil
		}

		bufIO := NewStreamingRead(128, src)
		_, err := io.Copy(destSrcBuf, bufIO)
		assert.NoError(t, err)
		assert.Equal(t, accSrcBuf, destSrcBuf.Bytes())
	})
}

func TestChainRead(t *testing.T) {
	t.Run("chain | via io.Read ", func(t *testing.T) {
		// 对于 128 size 大小的 buff， 循环 10000 次随机读取 1-128 大小的数据。保证数据的完整性
		accSrcBuf := make([]byte, 0)
		destSrcBuf := new(bytes.Buffer)

		count1 := 0
		src1 := func() ([]byte, error) {

			if count1 == 10000 {
				return nil, io.EOF
			}
			count1 += 1

			n := mrand.Intn(128) + 1
			buf := make([]byte, n)
			nr, err := rand.Read(buf)
			if err != nil {
				return nil, err
			}
			if nr != n {
				return nil, fmt.Errorf("tried to read %d bytes, but got %d bytes", n, nr)
			}
			accSrcBuf = append(accSrcBuf, buf...)
			return buf, nil
		}
		bufIO1 := NewStreamingRead(128, src1)

		count2 := 0
		src2 := func() ([]byte, error) {

			if count2 == 10000 {
				return nil, io.EOF
			}
			count2 += 1

			n := mrand.Intn(128) + 1
			buf := make([]byte, n)
			nr, err := rand.Read(buf)
			if err != nil {
				return nil, err
			}
			if nr != n {
				return nil, fmt.Errorf("tried to read %d bytes, but got %d bytes", n, nr)
			}
			accSrcBuf = append(accSrcBuf, buf...)
			return buf, nil
		}
		bufIO2 := NewStreamingRead(128, src2)

		chainIO := &ChainIoReader{}
		chainIO.Add(bufIO1)
		chainIO.Add(bufIO2)
		_, err := io.Copy(destSrcBuf, chainIO)
		assert.NoError(t, err)
		assert.Equal(t, accSrcBuf, destSrcBuf.Bytes())
	})

	t.Run("chain | via Iterator ", func(t *testing.T) {
		// 对于 128 size 大小的 buff， 循环 10000 次随机读取 1-128 大小的数据。保证数据的完整性
		accSrcBuf := make([]byte, 0)
		destSrcBuf := new(bytes.Buffer)

		count1 := 0
		src1 := func() ([]byte, error) {

			if count1 == 10000 {
				return nil, io.EOF
			}
			count1 += 1

			n := mrand.Intn(128) + 1
			buf := make([]byte, n)
			nr, err := rand.Read(buf)
			if err != nil {
				return nil, err
			}
			if nr != n {
				return nil, fmt.Errorf("tried to read %d bytes, but got %d bytes", n, nr)
			}
			accSrcBuf = append(accSrcBuf, buf...)
			return buf, nil
		}
		bufIO1 := NewStreamingRead(128, src1)

		count2 := 0
		src2 := func() ([]byte, error) {

			if count2 == 10000 {
				return nil, io.EOF
			}
			count2 += 1

			n := mrand.Intn(128) + 1
			buf := make([]byte, n)
			nr, err := rand.Read(buf)
			if err != nil {
				return nil, err
			}
			if nr != n {
				return nil, fmt.Errorf("tried to read %d bytes, but got %d bytes", n, nr)
			}
			accSrcBuf = append(accSrcBuf, buf...)
			return buf, nil
		}
		bufIO2 := NewStreamingRead(128, src2)

		chainIO := &ChainIoReader{iteratorBufSize: 115}
		chainIO.Add(bufIO1)
		chainIO.Add(bufIO2)
		iterator := chainIO.Iterator()

		composedIt := NewStreamingRead(128, iterator)
		_, err := io.Copy(destSrcBuf, composedIt)
		assert.NoError(t, err)
		assert.Equal(t, accSrcBuf, destSrcBuf.Bytes())
	})

	t.Run("chain | empty ", func(t *testing.T) {
		destSrcBuf := new(bytes.Buffer)
		chainIO := &ChainIoReader{}

		_, err := io.Copy(destSrcBuf, chainIO)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(destSrcBuf.Bytes()))
	})
}
