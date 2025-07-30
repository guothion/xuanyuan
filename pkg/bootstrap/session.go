package bootstrap

import "fmt"

type sessionService struct{}

type LoginIdentity struct {
	ID   interface{} // unit64 for MySQL or string for mongo
	Name string
}

type LoginRequest struct {
	Username string `form: "username" json:"username" binding:"omitempty"`
	Password string `form: "password" json:"password" binding:"omitempty"`
}

// 这里我们定义了一个 String 方法，之后我们打印这个 LoginRequest 的时候直接就是返回这个
func (r *LoginRequest) String() string {
	return fmt.Sprintf("username: %s, password: %s", r.Username, r.Password)
}

func (s *sessionService) Login(req *LoginRequest) (result *LoginIdentity, err error) {
	err = fmt.Errorf("sessionService.Login not implementated yet")
	return
}

func (s *sessionService) ValidateAccess(accessToken, account string) (err error) { return }
