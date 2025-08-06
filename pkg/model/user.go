package model

import (
	"github.com/guothion/xuanyuan/pkg/dataSource/mysql"
	"time"
)

type User struct {
	Id           uint   `json:"id,omitempty" gorm:"primary_key"`
	Username     string `json:"username,omitempty" bson:"username"`
	Email        string `json:"email,omitempty" bson:"email"`
	PasswordHash string `json:"password_hash,omitempty" bson:"password_hash"`
	Role         string `json:"role,omitempty" bson:"role"`
	CreatedAt    string `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt    string `json:"updated_at,omitempty" bson:"updated_at"`
}

type ListUserResponse struct {
	Data  []*User `json:"data"`
	Total int64   `json:"total"`
}

type UserRepositoryImpl struct{}

func NewUser(username string, email string, password string, role string) *User {
	return &User{
		Username:     username,
		Email:        email,
		PasswordHash: password,
		Role:         role,
		CreatedAt:    time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:    time.Now().Format("2006-01-02 15:04:05"),
	}
}

func (db *UserRepositoryImpl) Create(username string, email string, password string, role string) error {
	user := NewUser(username, email, password, role)
	err := mysql.Creates(*user)
	return err
}

func (db *UserRepositoryImpl) Update() {

}

func (db *UserRepositoryImpl) List() {

}

func (db *UserRepositoryImpl) GetByID() {

}
