package app

import (
	config "devops-testing/cmd/config"
	"devops-testing/logger"
)

// App := app struct containing resources to implement business login
type App struct {
	Logger *logger.Logger
	Config *config.Config
}

// NewApp := returns new app object
func NewApp(c *config.Config) *App {
	l := logger.NewLogger(true, true, "catalogService")
	app := App{
		Logger: l,
		Config: c,
	}
	return &app
}
