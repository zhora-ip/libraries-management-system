package tasks

import "github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/db"

type TasksRepo struct {
	db db.DB
}

func NewTasks(database db.DB) *TasksRepo {
	return &TasksRepo{db: database}
}
