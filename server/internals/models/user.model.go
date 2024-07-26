package models

type User struct {
	UserName string `bson:"username"`
	Password string `bson:"password"`
	Id       int    `bson:"_id, omitempty"`
}
