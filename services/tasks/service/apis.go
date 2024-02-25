package service

import (
	"context"

	"github.com/tonychill/ifitu/apis/pb/go/tasks"
)

func (s *ServiceImpl) CreateTask(ctx context.Context, req *tasks.CreateTaskRequest) (*tasks.CreateTaskResponse, error) {
	panic("implement me please")
}
func (s *ServiceImpl) GetTask(ctx context.Context, req *tasks.GetTaskRequest) (*tasks.GetTaskResponse, error) {
	panic("implement me please")
}
func (s *ServiceImpl) ListTasks(ctx context.Context, req *tasks.ListTasksRequest) (*tasks.ListTasksResponse, error) {
	panic("implement me please")
}
func (s *ServiceImpl) UpdateTask(ctx context.Context, req *tasks.UpdateTaskRequest) (*tasks.UpdateTaskResponse, error) {
	panic("implement me please")
}
func (s *ServiceImpl) DeleteTask(ctx context.Context, req *tasks.DeleteTaskRequest) (*tasks.DeleteTaskResponse, error) {
	panic("implement me please")
}

func (s *ServiceImpl) Shutdown(ctx context.Context) error {
	s.shutdown = true
	s.repo.Shutdown(ctx)
	return nil
}
