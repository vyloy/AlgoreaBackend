package items

import (
	"github.com/go-chi/chi"

	"github.com/France-ioi/AlgoreaBackend/app/auth"
	s "github.com/France-ioi/AlgoreaBackend/app/service"
)

// Service is the mount point for services related to `items`
type Service struct {
	s.Base
}

// SetRoutes defines the routes for this package in a route group
func (srv *Service) SetRoutes(router chi.Router) {
	router.Use(auth.UserIDMiddleware(&srv.Config.Auth))
	router.Post("/items/", s.AppHandler(srv.addItem).ServeHTTP)
	router.Get("/items/", s.AppHandler(srv.getList).ServeHTTP)
}
