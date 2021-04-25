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

//todo: this is way too big, should be optimized
func (F FSRepository) GetConfig(appName string) (domain.ConfigList, error) {
	var configList domain.ConfigList
	if appName != "global" {
		configList, _ = F.GetConfig("global")
	}

	if configList == nil {
		configList = domain.ConfigList{}
	}

	dat, err := ioutil.ReadFile(fmt.Sprintf("%s%s", F.Path, appName))

	if err != nil {
		return nil, err
	}

	str := string(dat)

	stringList := strings.Split(str, "\n")

	for _, str := range stringList {
		i := strings.Index(str, "=")
		if i != -1 {
			key := str[:i]
			value := str[i+1:]

			addValue := true
			for _, entry := range configList {
				if entry.Key == key {
					entry.Value = value
					addValue = false
				}
			}

			if addValue {
				entry := domain.ConfigEntry{Key: key, Value: value}
				configList = append(configList, &entry)
			}
		}
	}

	return configList, nil
}

func NewFsRepository(configuration config.RepoConfiguration) usecases.ConfigRepo {
	return FSRepository{
		Path: configuration.Path,
	}
}
