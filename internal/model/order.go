package model

type StatusEnum string

type Orders struct {
	ID
	Timestamp
}

func (Orders) TableName() string {
	return "orders"
}

func (Orders) InsertableFields() FieldSet {
	return FieldSet{}
}

func (Orders) UpdatableFields() FieldSet {
	return FieldSet{}
}
