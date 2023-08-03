package helpers

import "github.com/GehirnInc/crypt/apr1_crypt"

// HashApr1 模拟htpasswd生成密码
func HashApr1(password string) string {
	s, err := apr1_crypt.New().Generate([]byte(password), nil)
	if err != nil {
		panic(err)
	}
	return s
}
