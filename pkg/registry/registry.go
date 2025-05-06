package registry

import "github.com/riwn/home_discord_buzybot/pkg/setting"

type Registry struct {
	// TODO: Add Service
	repo Repository
}

func New(config setting.Config) Registry {
	return Registry{
		repo: NewRepository(config),
	}
}

func (r Registry) Repository() Repository {
	return r.repo
}
