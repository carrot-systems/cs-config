package rest

import "config/src/core/usecases"

type RoutesHandler struct {
	Usecases usecases.Usecases
}

func NewRouter(ucHandler usecases.Usecases) RoutesHandler {
	return RoutesHandler{Usecases: ucHandler}
}
