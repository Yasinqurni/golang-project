package di

import (
	controller "golang-project/internal/controller"
	repository "golang-project/internal/repository"
	"golang-project/internal/router"
	service "golang-project/internal/service"
	"golang-project/pkg/config"
	recovery "golang-project/pkg/recover"
	"net/http"

	"github.com/gorilla/mux"
)

func NewInternal(cfg *config.Config) *mux.Router {

	var (
		r = mux.NewRouter()

		repo = repository.NewRepository(cfg)
		s    = service.NewService(repo)
		c    = controller.NewController(s)
	)

	r.Use(recovery.Recovery)

	api := r.PathPrefix("/api").Subrouter()

	// Check health endpoint
	api.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	return router.NewRouter(api, c)
}
