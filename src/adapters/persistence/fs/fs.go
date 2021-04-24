package fs

import (
	"config/src/config"
	"config/src/core/domain"
	"config/src/core/usecases"
	"fmt"
	"io/ioutil"
	"strings"
)

type FSRepository struct {
	Path string
}

func (F FSRepository) GetConfig(appName string) (domain.ConfigList, error) {
	dat, err := ioutil.ReadFile(fmt.Sprintf("%s%s", F.Path, appName))

	if err != nil {
		return nil, err
	}

	config := domain.ConfigList{}
	str := string(dat)

	stringList := strings.Split(str, "\n")

	for _, str := range stringList {
		i := strings.Index(str, "=")
		if i != -1 {
			key := str[:i]
			value := str[i+1:]

			entry := domain.ConfigEntry{Key: key, Value: value}
			config = append(config, entry)
		}
	}

	return config, nil
}

func NewFsRepository(configuration config.RepoConfiguration) usecases.ConfigRepo {
	return FSRepository{
		Path: configuration.Path,
	}
}
