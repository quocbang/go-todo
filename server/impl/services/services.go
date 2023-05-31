package services

import (
	"quocbang/golang-to-do-list/swagger/restapi/operations/to_do_list"

	"github.com/go-openapi/runtime/middleware"
)

type ToDoList interface {
	GetAllToDoList(params to_do_list.GetAllToDoListParams) middleware.Responder
	CreateToDoList(params to_do_list.CreateToDoListParams) middleware.Responder
	UpdateStaus(params to_do_list.UpdateStatusParams) middleware.Responder
	DeleteToDoLists(params to_do_list.DeleteToDoListsParams) middleware.Responder
}

type Service struct {
	toDoList ToDoList
}

func NewService(toDoList ToDoList) *Service {
	return &Service{toDoList: toDoList}
}

func (s *Service) ToDoList() ToDoList {
	return s.toDoList
}
