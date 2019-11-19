package app

import (
	"fmt"
)

func NewApp(config Config) (App, error) {
	var app App
	switch config.Router {
	case ROUTER_GIN:
		app = NewAppGin()
	case ROUTER_DEFAULT_TESTER:
		app = NewDefaultAppTester()
	default:
		return nil, fmt.Errorf("Router not implemented: %s", config.Router)
	}

	app.InitMiddlewares(config)

	err := app.InitContext(config)
	if err != nil {
		return nil, err
	}

	return app, nil
}
