package models

type Task struct {
	Username string `bson:"user_name" json:"user_name" form:"user_name"`
	TaskName string `bson:"task_name" json:"task_name" form:"task_name"`
	Status   bool   `bson:"status" json:"status" form:"status"`
}
