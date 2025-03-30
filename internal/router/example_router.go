package router

import (
	"fmt"
	"golang-project/internal/controller"

	"github.com/gorilla/mux"
)

type ExampleRouter struct {
	exampleController controller.ExampleController
	r                 *mux.Router
}

func NewExampleRouter(r *mux.Router, exampleController controller.ExampleController) *ExampleRouter {
	return &ExampleRouter{
		exampleController: exampleController,
		r:                 r,
	}
}

func (er *ExampleRouter) V1(path string) {
	er.r.HandleFunc(fmt.Sprintf("%s/example", path), er.exampleController.Create).Methods("POST")
}
