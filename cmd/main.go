package main

import (
	"config/src/adapters/persistence/fs"
	"config/src/adapters/rest"
	"config/src/config"
	"config/src/core/usecases"
	env "github.com/carrot-systems/csl-env"
	"log"
)

func main() {
	env.LoadEnv()

	//Oh boy this is gonna get confusing real quick...
	configurationRepoConfiguration := config.LoadRepoConfiguration()
	ginConfiguration := config.LoadGinConfiguration()

	var configurationRepo usecases.ConfigRepo

	switch configurationRepoConfiguration.RepoType {
	case "LOCAL":
		configurationRepo = fs.NewFsRepository(configurationRepoConfiguration)
	default:
		log.Fatalln("Error: Only LOCAL is supported for REPO_TYPE at the moment")
	}

	usecasesHandler := usecases.NewInteractor(configurationRepo)

	restServer := rest.NewServer(ginConfiguration)
	routesHandler := rest.NewRouter(usecasesHandler)

	rest.SetRoutes(restServer.Router, routesHandler)
	restServer.Start()
}
