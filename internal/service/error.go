package service

import "fmt"

// 定义一个自定义错误类型
type ex string

// 实现 error 接口
func (e ex) Error() error {
	return fmt.Errorf(string(e))
}

// 定义错误枚举
const (
	UnknownError               ex = "UnknownError"
	UserNotFoundError          ex = "UserNotFoundError"
	UserAlreadyExistsError     ex = "UserAlreadyExistsError"
	CreateUserFailedError      ex = "CreateUserFailedError"
	LinkNotFoundError          ex = "LinkNotFoundError"
	NoPermissionToOperateError ex = "NoPermissionToOperateError"
)
