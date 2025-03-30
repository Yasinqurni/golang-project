package controller

import (
	"encoding/json"
	"golang-project/internal/entity/request"
	"golang-project/internal/service"
	"golang-project/pkg/response"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ExampleController interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type exampleController struct {
	exampleService service.ExampleService
}

func NewExampleController(exampleService service.ExampleService) ExampleController {
	return &exampleController{
		exampleService: exampleService,
	}
}

func (e *exampleController) Create(w http.ResponseWriter, r *http.Request) {
	var (
		req           request.ExampleBodyRequest
		errorMessages = make(map[string]string)
	)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		response.ErrorResponse(w, "Invalid JSON body", err, http.StatusBadRequest)
		return
	}

	validate := validator.New()

	if err := validate.Struct(&req); err != nil {
		validationErr, ok := err.(validator.ValidationErrors)
		if !ok {
			response.ErrorResponse(w, "Invalid JSON body", err, http.StatusBadRequest)
			return
		}

		for _, e := range validationErr {
			fieldJSONName := req.GetJsonFieldName(e.Field())
			errorMessages[fieldJSONName] = req.ErrMessages()[fieldJSONName][e.ActualTag()]
		}
	}
	if len(errorMessages) > 0 {
		response.ErrorResponse(w, "Invalid JSON body", errorMessages, http.StatusBadRequest)
		return
	}

	ids, err := e.exampleService.Create(req)
	if err != nil {
		if errors, ok := err.(*response.Err); ok {
			response.ErrorResponse(w, "Invalid JSON body", errors, http.StatusBadRequest)
			return
		}
		response.ErrorResponse(w, "Invalid JSON body", err, http.StatusInternalServerError)
		return
	}

	response.SuccessResponse(w, "Success Create Examples Data", ids, nil, http.StatusCreated)
}
