package database

import (
	"log"

	"github.com/go-xorm/xorm"
)

var Engine *xorm.Engine

func Init() {
	var err error
	Engine, err = xorm.NewEngine("mysql", "root:@/test?charset=utf8")
	if err != nil {
		log.Fatalf("init error", err)
	}

}
