package main

import (
	"github.com/martini-contrib/binding"
	//"log"
	"net/http"
	"strings"
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

func (theme *Theme) create() error {
	//insert to database
	theme.CreatedAt = time.Now().UTC()
	theme.UpdatedAt = time.Now().UTC()
	_, err := _engine.Insert(theme)
	if err != nil {
		errMsg := err.Error()
		if strings.Contains(errMsg, "UNIQUE constraint failed:") {
			myErr := appError{Ex: err, Message: "theme name was already existing.", Code: 4001}
			return &myErr
		}
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

func getThemeByName(storeId int64, themeName string) *Theme {
	theme := Theme{}
	err := _engine.Where("storeId = ? and name = ?", storeId, themeName).Find(&theme)

	if err != nil {
		panic(err)
	}
	return &theme
}
