package gacos

import "bytes"

func bytesCombine(pBytes ...[]byte) []byte {
	lenth := len(pBytes)
	s := make([][]byte, lenth)
	for i, v := range pBytes {
		s[i] = v
	}
	sep := []byte("")
	return bytes.Join(s, sep)
}
