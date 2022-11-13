package application

import "ddd_demo/internal/infrastructure"

type Apps struct {
	UserApp *UserApp
}

func NewApps(apps *infrastructure.Repos) *Apps {
	return &Apps{
		UserApp: NewUserApp(apps.UserRepo, apps.AuthRepo),
	}
}
