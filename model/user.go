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

func (me *User) BeforeInsert() {
	me.CreatedUnix = time.Now().Unix()
	me.UpdatedUnix = me.CreatedUnix
}

func (me *User) BeforeUpdate() {
	me.UpdatedUnix = time.Now().Unix()
}

func (me *User) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "created_unix":
		me.Created = time.Unix(me.CreatedUnix, 0).Local()
	case "updated_unix":
		me.Updated = time.Unix(me.UpdatedUnix, 0)
	}
}