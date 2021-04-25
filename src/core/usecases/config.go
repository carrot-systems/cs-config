package usecases

import "config/src/core/domain"

func (i interactor) GetConfig(appName string) (domain.ConfigList, error) {
	return i.configRepo.GetConfig(appName)
}
