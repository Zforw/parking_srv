package utils

import (
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"parking/global"
	"strings"
)

func EncryptPass(pass string) string {
	salt, encodedPwd := password.Encode(pass, global.OP)
	return fmt.Sprintf("%s$%s", salt, encodedPwd)
}

func VerifyPass(pass, encrypted string) bool {
	passwordInfo := strings.Split(encrypted, "$")
	return password.Verify(pass, passwordInfo[0], passwordInfo[1], global.OP)
}
