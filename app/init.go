package app

import (
	"github.com/revel/revel"
	"nttdata/app/models/mongodb"
)

func init() {
	// Filters is the default set of global filters.
	revel.OnAppStart(initApp)
}

func initApp() {
	mongodb.MaxPool = revel.Config.IntDefault("mongo.maxPool", 0)
	mongodb.PATH,_ = revel.Config.String("mongo.path")
	mongodb.DBNAME, _ = revel.Config.String("mongo.database")
	mongodb.CheckAndInitServiceConnection()
}
