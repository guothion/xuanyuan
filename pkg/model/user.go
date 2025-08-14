package model

import "strconv"

type Permission int

const (
	UserRole Permission = 1 << iota
	EditorRole
	AdminRole
)

type User struct {
	ID
	Username     string  `json:"username,omitempty" bson:"username" gorm:"size:50;not null;unique;comment:用户名"`
	Email        *string `json:"email,omitempty" bson:"email" gorm:"size:100;not null;unique;comment:邮箱"`
	PasswordHash string  `json:"password_hash,omitempty" bson:"password_hash" gorm:"size:255;not null;comment:密码哈希"`
	Role         int8    `json:"role,omitempty" bson:"role" gorm:"type:tinyint;not null;default:1;comment:权限"`
	Timestamp
	SoftDeletes
}

type ListUserResponse struct {
	Data  []*User `json:"data"`
	Total int64   `json:"total"`
}

func (u User) GetUid() string {
	return strconv.Itoa(int(u.ID.ID))
}
