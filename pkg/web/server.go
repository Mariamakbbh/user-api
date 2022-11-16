package web

import (
	"user-api/pkg/user"

	config "github.com/shipperizer/miniature-monkey/config"
	core "github.com/shipperizer/miniature-monkey/core"
	"github.com/shipperizer/miniature-monkey/webiface"

	//monitoring "github.com/shipperizer/miniature-monkey/monitoring"

	"go.uber.org/zap"
)

// func NewAPI(userService user.Service, monitor monitoring.MonitorInterface, logger *zap.SugaredLogger) APIInterface {
func NewApi(userService user.Service, logger *zap.SugaredLogger) webiface.APIInterface {

	api := core.NewAPI(
		config.NewAPIConfig(
			"user-api",
			config.NewCORSConfig("google.com"),
			//monitor,
			nil,
			logger,
		),
	)

	// register blueprints
	api.RegisterBlueprints(
		api.Router(),
		user.NewBlueprint(userService),
		//user.NewBlueprint(userService, monitor),
	)

	return api
}
