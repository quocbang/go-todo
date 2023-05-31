package impl

import (
	"quocbang/golang-to-do-list/impl/handlers/todo"
	"quocbang/golang-to-do-list/impl/services"
	"quocbang/golang-to-do-list/swagger/restapi/operations"
	"quocbang/golang-to-do-list/swagger/restapi/operations/to_do_list"

	"gorm.io/gorm"
)

func RegisterService(db *gorm.DB) *services.Service {
	return services.NewService(todo.NewToDoList(db))
}

// RegisterHandlers register real handlers
func RegisterHandlers(api *operations.ToDoListAPI, config *gorm.DB) error {
	s := RegisterService(config)

	api.ToDoListGetAllToDoListHandler = to_do_list.GetAllToDoListHandlerFunc(s.ToDoList().GetAllToDoList)
	api.ToDoListCreateToDoListHandler = to_do_list.CreateToDoListHandlerFunc(s.ToDoList().CreateToDoList)
	api.ToDoListUpdateStatusHandler = to_do_list.UpdateStatusHandlerFunc(s.ToDoList().UpdateStaus)
	api.ToDoListDeleteToDoListsHandler = to_do_list.DeleteToDoListsHandlerFunc(s.ToDoList().DeleteToDoLists)

	return nil
}
