package todo

import (
	"net/http"
	ormModels "quocbang/golang-to-do-list/impl/orm/models"
	"quocbang/golang-to-do-list/impl/services"
	"quocbang/golang-to-do-list/swagger/models"
	"quocbang/golang-to-do-list/swagger/restapi/operations/to_do_list"

	"github.com/go-openapi/runtime/middleware"
	"gorm.io/gorm"
)

var (
	Doing    = "Doing"
	Finished = "Finished"
)

type DB struct {
	db *gorm.DB
}

func NewToDoList(db *gorm.DB) services.ToDoList {
	return DB{db: db}
}

func (d DB) GetAllToDoList(params to_do_list.GetAllToDoListParams) middleware.Responder {
	todo := []ormModels.ToDoList{}
	if err := d.db.Find(&todo).Error; err != nil {
		return to_do_list.NewGetAllToDoListDefault(http.StatusInternalServerError).WithPayload(&models.Error{Details: err.Error()})
	}

	modelsToDo := make(models.GetToDoListRespone, len(todo))
	for idx, t := range todo {
		modelsToDo[idx] = &models.GetToDoListResponeItems0{
			ID:     t.ID,
			Status: t.Status,
			Title:  t.Title,
		}
	}
	return to_do_list.NewGetAllToDoListOK().WithPayload(&to_do_list.GetAllToDoListOKBody{
		Data: modelsToDo,
	})
}

func (d DB) CreateToDoList(params to_do_list.CreateToDoListParams) middleware.Responder {
	if err := d.db.Create(&ormModels.ToDoList{
		Title:  params.Body.Title,
		Status: Doing,
	}).Error; err != nil {
		return to_do_list.NewCreateToDoListDefault(http.StatusInternalServerError).WithPayload(&models.Error{Details: err.Error()})
	}

	return to_do_list.NewCreateToDoListOK()
}

func (d DB) UpdateStaus(params to_do_list.UpdateStatusParams) middleware.Responder {
	if err := d.db.Model(&ormModels.ToDoList{}).Where("id = ?", params.Body.ID).UpdateColumn("status", Finished).Error; err != nil {
		return to_do_list.NewUpdateStatusDefault(http.StatusInternalServerError).WithPayload(&models.Error{Details: err.Error()})
	}

	return to_do_list.NewUpdateStatusOK()
}

func (d DB) DeleteToDoLists(params to_do_list.DeleteToDoListsParams) middleware.Responder {
	delIDS := make([]ormModels.ToDoList, len(params.Body))
	for idx, p := range params.Body {
		delIDS[idx] = ormModels.ToDoList{ID: p}
	}

	if err := d.db.Delete(&delIDS).Error; err != nil {
		return to_do_list.NewDeleteToDoListsDefault(http.StatusInternalServerError).WithPayload(&models.Error{Details: err.Error()})
	}

	return to_do_list.NewDeleteToDoListsOK()
}
