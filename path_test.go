package encoded_path

import (
	"bytes"
	"testing"
)

func TestEncodeValue(t *testing.T) {
	s := ""
	b := bytes.NewBufferString(s)
	EncodeValue(b, -126.453)
}

func TestEncodePath(t *testing.T) {
	s := ""
	b := bytes.NewBufferString(s)

	path := [][2]float32{
		{38.5, -120.2},
		{40.7, -120.95},
		{43.252, -126.453}}

	EncodePath(b, path)
}

func BenchmarkEncodeValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf []byte
		b := bytes.NewBuffer(buf)
		EncodeValue(b, -179.9832104)
	}
}

func BenchmarkEncodePath(b *testing.B) {
	path := [][2]float32{
		{38.5, -120.2},
		{40.7, -120.95},
		{43.252, -126.453}}

	for i := 0; i < b.N; i++ {
		var buf []byte
		b := bytes.NewBuffer(buf)
		EncodePath(b, path)
	}
}
