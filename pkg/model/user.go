package model

import "github.com/kamva/mgm/v3"

type User struct {
	mgm.IDField `json:",inline" bson:",inline"`
	Name        string `json:"name,omitempty" bson:"name"`
	Position    string `json:"position,omitempty" bson:"position"`
	Country     string `json:"country,omitempty" bson:"country"`
	Email       string `json:"email,omitempty" bson:"email"`
	Phone       string `json:"phone,omitempty" bson:"phone"`
	Gender      int8   `json:"gender,omitempty" bson:"gender"`
}

type ListUserResponse struct {
	Data  []*User `json:"data"`
	Total int64   `json:"total"`
}
