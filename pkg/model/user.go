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
	Username     string  `json:"username,omitempty" bson:"username"`
	Email        *string `json:"email,omitempty" bson:"email"`
	PasswordHash string  `json:"password_hash,omitempty" bson:"password_hash"`
	Role         int     `json:"role,omitempty" bson:"role"`
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
