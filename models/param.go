package models

// 定义请求参数结构体

// 利用validator库参数校验
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type Login struct {
	Username string `db:"username" binding:"required"`
	Password string `db:"password" binding:"required"`
}
