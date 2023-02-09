package utils

import (
	"errors"
	"io"

	"github.com/smallnest/ringbuffer"
)

type IterFunc func() ([]byte, error)

type StreamingRead struct {
	iter       IterFunc
	ringbuf    *ringbuffer.RingBuffer
	bufferSize int
}

func NewStreamingRead(buffSize int, iterSrc IterFunc) *StreamingRead {
	return &StreamingRead{iter: iterSrc, ringbuf: ringbuffer.New(buffSize), bufferSize: buffSize}
}

func (o *StreamingRead) Read(p []byte) (int, error) {
	if o.ringbuf.Length() == 0 {
		data, err := o.iter()
		if err != nil {
			return 0, err
		}

		if len(p) >= len(data) {
			return copy(p, data), nil
		} else {
			_, err = o.ringbuf.Write(data)
			if err != nil {
				return 0, err
			}
		}
	}
	return o.ringbuf.Read(p)

}

func NewChainIoReader(iteratorBufSize int) *ChainIoReader {
	return &ChainIoReader{iteratorBufSize: iteratorBufSize}
}

type ChainIoReader struct {
	r               []io.Reader
	idx             int
	iteratorBufSize int
}

func (o *ChainIoReader) Add(src io.Reader) {
	if o.r == nil {
		o.r = make([]io.Reader, 0)
	}
	o.r = append(o.r, src)
}

func (o *ChainIoReader) Iterator() func() ([]byte, error) {
	bufSize := o.iteratorBufSize
	if o.iteratorBufSize == 0 {
		bufSize = 4096
	}
	buf := make([]byte, bufSize)
	return func() ([]byte, error) {
		count, err := o.Read(buf)
		if err != nil {
			return nil, err
		}
		return buf[:count], nil
	}
}

func (o *ChainIoReader) Read(p []byte) (int, error) {
	count, err := o.read(p)
	if err != nil && errors.Is(err, io.EOF) && len(o.r) > o.idx {
		return o.read(p)
	}
	return count, err
}

func (o *ChainIoReader) read(p []byte) (int, error) {
	if o.idx >= len(o.r) {
		return 0, io.EOF
	}
	count, err := o.r[o.idx].Read(p)
	if err != nil {
		if errors.Is(err, io.EOF) {
			o.idx += 1
		}
	}
	return count, err
}
