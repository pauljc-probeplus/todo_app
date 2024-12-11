package models

type User struct {
	Username string `bson:"user_name" json:"user_name"`
	Password string `bson:"password" json:"password"`
}
