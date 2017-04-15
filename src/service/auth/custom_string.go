package auth

import
(
	"crypto/sha1"
	"encoding/hex"
)

// copied from
// https://github.com/ikeikeikeike/gopkg/blob/master/convert/convert.go

type CustomStr string

func (f *CustomStr) Clear() {
	*f = CustomStr(0x1E)
}

func (f CustomStr) Exist() bool {
	return string(f) != string(0x1E)
}

func (m CustomStr) String() string {
	if m.Exist() {
		return string(m)
	}
	return ""
}

func (m CustomStr) ToSHA1() string {
	h := sha1.New()
	h.Write([]byte(m.String()))
	return hex.EncodeToString(h.Sum(nil))
}