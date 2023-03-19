package util

import (
	"io"
	"log"
	"strings"
)

func ReaderToString(r io.Reader) string {
	buf := new(strings.Builder)

	_, err := io.Copy(buf, r)
	if err != nil {
		log.Fatalf("Failed copying buffers: %v", err)
		return ""
	}

	return buf.String()
}
