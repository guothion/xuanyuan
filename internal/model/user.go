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
	Username string  `json:"username,omitempty"  gorm:"size:50;not null;unique;comment:用户名"`
	Email    *string `json:"email,omitempty"  gorm:"size:100;not null;unique;comment:邮箱"`
	Password string  `json:"password,omitempty"  gorm:"size:255;not null;comment:密码哈希"`
	Role     int8    `json:"role,omitempty"  gorm:"type:tinyint;not null;default:1;comment:权限"`
	Timestamp
	SoftDeletes
}

func (User) TableName() string {
	return "users"
}

func (User) InsertableFields() FieldSet {
	return FieldSet{
		"username": true,
		"email":    true,
		"password": true,
		"role":     true,
	}
}

func (User) UpdatableFields() FieldSet {
	return FieldSet{
		"username":   true,
		"email":      true,
		"password":   true,
		"role":       true,
		"is_deleted": true,
	}
}

type ListUserResponse struct {
	Data  []*User `json:"data"`
	Total int64   `json:"total"`
}

func (u User) GetUid() string {
	return strconv.Itoa(int(u.ID.ID))
}
