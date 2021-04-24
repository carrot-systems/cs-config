package usecases

import "config/src/core/domain"

type ConfigRepo interface {
	GetConfig(appName string) (domain.ConfigList, error)
}

type interactor struct {
	configRepo ConfigRepo
}

func NewInteractor(cR ConfigRepo) interactor {
	return interactor{configRepo: cR}
}
