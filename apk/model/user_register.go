package model

type UserRegister struct {
	Id       int
	UserName string `grom:"default:'galeone'"`
	NickName string
	UserFace string
	Password int
	Enable   int
	Email    string
	RegTime  string
}
