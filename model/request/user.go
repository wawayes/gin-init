package request

// Register 用户注册
type Register struct {
	UserAccount   string `json:"userAccount" example:"账户名"`
	UserPassword  string `json:"userPassword" example:"密码"`
	CheckPassword string `json:"checkPassword" example:"二次输入密码"`
}

// Login 用户登录
type Login struct {
	UserAccount  string `json:"userAccount"`  // 账户名
	UserPassword string `json:"userPassword"` // 密码
}

// GetCurrentUser 获取当前用户
type GetCurrentUser struct {
	ID uint `json:"id"`
}

// GetUserById 根据id查询用户
type GetUserById struct {
	ID uint `json:"id"`
}
