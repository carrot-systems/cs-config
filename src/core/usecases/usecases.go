package usecases

import "config/src/core/domain"

type Usecases interface {
	GetConfig(appName string) (domain.ConfigList, error)
}
