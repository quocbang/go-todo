package models

type ToDoList struct {
	ID         int64  `gorm:"autoIncrement"`
	Title      string `gorm:"type:text"`
	Status     string `gorm:"type:text"`
	Created_at int64  `gorm:"autoCreateTime:nano;not null"`
}

func (ToDoList) TableName() string {
	return "to_do_list"
}
