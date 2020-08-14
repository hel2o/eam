package main

import (
	"github.com/hel2o/eam/db"
	"github.com/hel2o/eam/g"
	"github.com/hel2o/eam/router"
)

func main() {
	g.Init()
	g.AppLog()
	db.InitDB()
	router.Start()
}
