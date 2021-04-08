package crypto

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

//------------------------------------------------------------------------------------------------------------[crypto.String]

// String ...
type String string

// ToEncodeBase64 ...
func (r String) ToEncodeBase64() string {
	return base64.StdEncoding.EncodeToString([]byte(r))
}

// ToDecodeBase64 ...
func (r String) ToDecodeBase64() ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(r))
}

// ToMD5 ...
func (r String) ToMD5() string {
	h := md5.New()
	io.WriteString(h, string(r))

	return fmt.Sprintf("%x", h.Sum(nil))
}

// ToSHA256 ...
func (r String) ToSHA256() string {
	h := sha256.New()
	h.Write([]byte(r))
	// hex.EncodeToString(...)
	return fmt.Sprintf("%x", h.Sum(nil))
}
