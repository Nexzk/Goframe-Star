package utility

import "github.com/gogf/gf/v2/crypto/gmd5"

var JwtKey = []byte("db03d23b03ec405793b38f10592a2f34")

func EncryptPassword(password string) string {
	return gmd5.MustEncryptString(password)
}
