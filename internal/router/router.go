package router

import (
	"golang-project/internal/controller"

	"github.com/gorilla/mux"
)

func NewRouter(r *mux.Router, c *controller.Controller) *mux.Router {

	exampleRouter := NewExampleRouter(r, c.ExampleController)

	// V1
	{
		path := "/v1"
		exampleRouter.V1(path)
	}

	return r
}
