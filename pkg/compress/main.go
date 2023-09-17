package compress

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"github.com/dsnet/compress/bzip2"
	"io"
)

const (
	Gzip    = "gzip"
	Brotli  = "br"
	Deflate = "deflate"
)

type (
	Compressor interface {
		Compress(encoding string, body string) []byte
		Decompress(encoding string, body io.ReadCloser) string
	}

	compressor struct{}
)

func compress(data []byte, encoder io.WriteCloser) {
	if _, err := encoder.Write(data); err != nil {
		_ = encoder.Close()
		panic(err)
	}
	err := encoder.Close()
	if err != nil {
		panic(err)
	}
}

func (c *compressor) Compress(encoding string, data string) []byte {
	dataBytes := []byte(data)
	var buffer bytes.Buffer

	switch encoding {
	case Gzip:
		compress(dataBytes, gzip.NewWriter(&buffer))
	case Brotli:
		encoder, err := bzip2.NewWriter(&buffer, nil)
		if err != nil {
			panic(err)
		}
		compress(dataBytes, encoder)
	case Deflate:
		encoder, err := flate.NewWriter(&buffer, flate.BestCompression)
		if err != nil {
			panic(err)
		}
		compress(dataBytes, encoder)
	default:
		buffer.Write(dataBytes)
	}

	return buffer.Bytes()
}

func (c *compressor) Decompress(encoding string, body io.ReadCloser) string {
	var reader io.Reader
	var err error

	switch encoding {
	case Gzip:
		reader, err = gzip.NewReader(body)
		if err != nil {
			panic(err)
		}
	case Brotli:
		reader, err = bzip2.NewReader(body, nil)
		if err != nil {
			panic(err)
		}
	case Deflate:
		reader = flate.NewReader(body)
	default:
		reader = body
	}
	decompressed, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	return string(decompressed)
}

func NewCompressor() Compressor {
	return &compressor{}
}
