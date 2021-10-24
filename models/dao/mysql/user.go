package mysql

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id         int            `gorm:"column:id" db:"id" json:"id" form:"id"`
	Username   string         `gorm:"column:username" db:"username" json:"username" form:"username"`
	Mobile     string         `gorm:"column:mobile" db:"mobile" json:"mobile" form:"mobile"`
	Password   string         `gorm:"column:password" db:"username" json:"-" form:"password"`
	Department string         `gorm:"column:department" db:"department" json:"department" form:"department"`
	CreatedAt  time.Time      `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" db:"deleted_at" json:"deleted_at" form:"deleted_at"`
}
