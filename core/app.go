package core

import (
	"github.com/kataras/iris/v12"
)

type Application struct {
	*iris.Application
}

func New() *Application {
	app := &Application{
		iris.New(),
	}
	return app
}
