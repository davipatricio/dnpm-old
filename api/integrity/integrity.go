package integrity

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
)

// returns true if the data matches the shasum
func CheckIntegrity(data []byte, shasum string) bool {
	h := sha1.New()
	io.Copy(h, bytes.NewReader(data))
	return fmt.Sprintf("%x", h.Sum(nil)) == shasum
}