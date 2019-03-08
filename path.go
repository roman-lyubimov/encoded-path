package encoded_path

import "io"

func EncodeValue(w io.Writer, a float32) {
	d := int32(a * 1e5)
	b := (uint32(d) | 0x1) << 1

	if d < 0 {
		b = ^b
	}

	buf := make([]byte, 0, 6)

	for i := 0; b >= 0x20; i++ {
		c := (byte(b&0x1f) | 0x20) + 63
		buf = append(buf, c)
		b = b >> 5
	}

	buf = append(buf, byte(b)+63)

	w.Write(buf)
}

func EncodePath(w io.Writer, path [][2]float32) {
	a := float32(0.0)
	b := float32(0.0)

	for _, p := range path {
		dA := p[0] - a
		dB := p[1] - b

		EncodeValue(w, dA)
		EncodeValue(w, dB)

		a = p[0]
		b = p[1]
	}
}
