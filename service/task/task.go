package task

import (
	"fmt"
	"todo-list/entity"
)

type ServiceRepository interface {
	//DoesThisUserHaveThisCategoryID(userID, categoryID int) bool
	CreateNewTask(t entity.Task) (entity.Task, error)
	ListUserTasks(userID int) ([]entity.Task, error)
}

type Service struct {
	repository ServiceRepository
}

func NewService(repo ServiceRepository) Service {
	return Service{
		repository: repo,
	}
}

type CreateRequest struct {
	Title               string
	DueDate             string
	CategoryID          int
	AuthenticatedUserID int
}

type CreateResponse struct {
	Task entity.Task
}

func (t Service) Create(req CreateRequest) (CreateResponse, error) {

	//if t.repository.DoesThisUserHaveThisCategoryID(req.AuthenticatedUserID, req.CategoryID) {
	//	return CreateResponse{}, fmt.Errorf("user does not have this category: %d", req.CategoryID)
	//}

	createdTask, cErr := t.repository.CreateNewTask(entity.Task{
		Title:      req.Title,
		DueDate:    req.DueDate,
		CategoryID: req.CategoryID,
		IsDone:     false,
		UserID:     req.AuthenticatedUserID,
	})
	if cErr != nil {
		return CreateResponse{}, fmt.Errorf("can't create new task: %v", cErr)
	}

	return CreateResponse{Task: createdTask}, nil
}

type ListRequest struct {
	UserID int
}

type ListResponse struct {
	Tasks []entity.Task
}

func (t Service) List(req ListRequest) (ListResponse, error) {
	tasks, err := t.repository.ListUserTasks(req.UserID)
	if err != nil {
		return ListResponse{}, fmt.Errorf("can't list user tasks: %v", err)
	}

	return ListResponse{Tasks: tasks}, nil
}
