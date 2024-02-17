package domain

import "time"

// User 领域对象 是DDD重点entity
type User struct {
	Id       int64
	Email    string
	Password string
	Ctime    time.Time
}

//func (u User) Encrypt()string  {
//
//}
//func (u User) ComparePassword()  {
//
//}
