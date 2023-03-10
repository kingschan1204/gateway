package util

import (
	"crypto/md5"
	"encoding/hex"
)

/*
func GetUUID() string {
	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return ""
	}
	return strings.Replace(u2.String(), "-", "", -1)
}*/

func GetMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
