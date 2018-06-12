package model

import "time"

type User struct {
	Id           int64     `json:"id" xorm:"pk"`
	Name         string    `json:"name" xorm:"varchar(250)"`
	Birthday     time.Time `json:"birthday" xorm:"-"`
	BirthdayUnix int64     `json:"-"`
	Gender       string    `json:"gender"`
	Age          int64     `json:"age"`
	TotalSpend   float64   `json:"totalSpend"`
	Address      string    `json:"address"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`

	Created     time.Time `xorm:"-"`
	CreatedUnix int64
	Updated     time.Time `xorm:"-"`
	UpdatedUnix int64
}
