package data

import (
	"gorm.io/gorm"
	"time"
)

type UserRsp struct {
	Id         int            `json:"id" form:"id"`
	Username   string         `json:"username" form:"username"`
	Mobile     string         `json:"mobile" form:"mobile"`
	Department string         `json:"department" form:"department"`
	CreatedAt  time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" form:"deleted_at"`
}

type UserData struct {
}

func (ud *UserData) List(param interface{}) {

}
func (ud *UserData) Info(param interface{}) {

}

func (ud *UserData) Add(param interface{}) {

}

func (ud *UserData) Save(param interface{}) {

}

func (ud *UserData) Del(param interface{}) {

}
