package models

import "gorm.io/gorm/schema"

type Model interface {
	schema.Tabler
}

func GetModelList() []Model {
	return []Model{
		&ToDoList{},
	}
}
