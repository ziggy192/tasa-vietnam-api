package controller

import (
	"github.com/jinzhu/gorm"
	"github.com/ziggy192/tasa-vietnam-api/pkg/repo"
	"github.com/ziggy192/tasa-vietnam-api/pkg/util"
)

func check(err error) {
	util.Check(err)
}

//package internal varable
var db *gorm.DB = repo.DB
