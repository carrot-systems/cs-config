package config

import env "github.com/carrot-systems/csl-env"

type RepoConfiguration struct {
	RepoType string
	Path     string
}

func LoadRepoConfiguration() RepoConfiguration {
	return RepoConfiguration{
		RepoType: env.RequireEnvString("REPO_TYPE"),
		Path:     env.RequireEnvString("REPO_PATH"),
	}
}
