package service

import (
	//"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(username, password string, c *gin.Context) error {
	// 简单的硬编码验证（实际应查询数据库）
	if username != "admin" || password != "password" {
		return ErrInvalidCredentials
	}

	//// 设置 Session
	//session := sessions.Default(c)
	//session.Set("username", username)
	//if err := session.Save(); err != nil {
	//	return err
	//}

	return nil
}

var ErrInvalidCredentials = &Error{Message: "Invalid credentials"}

type Error struct {
	Message string
}

func (e *Error) Error() string {
	return e.Message
}
