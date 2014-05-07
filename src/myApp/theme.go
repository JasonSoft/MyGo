package main

import (
	"github.com/martini-contrib/binding"
	//"log"
	"net/http"
	"time"
)

type Theme struct {
	Id        int64  `xorm:"index"`
	StoreId   int64  `xorm:"not null unique(theme)" form:"-" json:"-"`
	Name      string `xorm:"not null unique(theme)" binding:"required"`
	IsDefault bool
	TimeStamp string `form:"-" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (theme Theme) Validate(errors *binding.Errors, req *http.Request) {

	/*
		if len(theme.Name) < 4 {
			errors.Fields["Name"] = "Too short; minimum 4 characters"
		} else if len(theme.Name) > 120 {
			errors.Fields["Name"] = "Too long; maximum 120 characters"
		}*/
}

func (theme *Theme) Create() error {
	//insert to database
	_, err := _engine.Insert(theme)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func getThemes(storeId int64) *[]Theme {
	themes := make([]Theme, 0)
	err := _engine.Where("StoreId = ?", storeId).Find(&themes)

	if err != nil {
		panic(err)
	}

	return &themes
}
