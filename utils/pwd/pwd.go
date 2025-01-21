package pwd

import (
	"goblog_server/global"
	"golang.org/x/crypto/bcrypt"
)

// HashPwd 加密密码
func HashPwd(passwd string) string {
	password, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.MinCost)
	if err != nil {
		global.Log.Error(err.Error())
	}
	return string(password)
}

// CheckPwd 验证hash之后的密码
func CheckPwd(hashpasswd string, passwd string) bool {
	bytehash := []byte(hashpasswd)

	err := bcrypt.CompareHashAndPassword(bytehash, []byte(passwd))
	if err != nil {
		global.Log.Error(err.Error())
		return false
	}
	//fmt.Printf("\n\n")
	//fmt.Println("passwd:", passwd)
	//fmt.Println("hashpasswd:", hashpasswd)
	return true
}
