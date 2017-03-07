package auth

import
(
	"crypto/sha1"
	"encoding/hex"
)

// copied from
// https://github.com/ikeikeikeike/gopkg/blob/master/convert/convert.go

type customStr string

func (f *customStr) Clear() {
	*f = customStr(0x1E)
}

func (f customStr) Exist() bool {
	return string(f) != string(0x1E)
}

func (m customStr) String() string {
	if m.Exist() {
		return string(m)
	}
	return ""
}

func (m customStr) ToSHA1() string {
	h := sha1.New()
	h.Write([]byte(m.String()))
	return hex.EncodeToString(h.Sum(nil))
}