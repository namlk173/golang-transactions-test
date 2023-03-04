package boostrap

import (
	"banking/mongoImplement"
)

type Application struct {
	Env    *Env
	Client mongoImplement.Client
}

func NewApplication() Application {
	env := NewEnv()
	return Application{
		Env:    env,
		Client: NewMongoDatabase(env),
	}
}

func (app *Application) CloseDBConnection() {
	CloseMongoDatabase(app.Client)
}
