package request

type ExampleRequest struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required"`
	Phone   string `json:"phone"`
	Age     int    `json:"age" validate:"required"`
	Address string `json:"address"`
}

type ExampleBodyRequest struct {
	Examples []ExampleRequest `json:"examples" validate:"required,min=1,dive"`
}

func (req *ExampleBodyRequest) GetJsonFieldName(field string) string {
	return map[string]string{
		"Examples": "examples",
		"Name":     "name",
		"Email":    "email",
		"Age":      "age",
	}[field]
}

func (req *ExampleBodyRequest) ErrMessages() map[string]map[string]string {
	return map[string]map[string]string{
		"examples": {
			"required": "examples is required",
			"min":      "examples must contain at least one item",
		},
		"name":  {"required": "name is required"},
		"email": {"required": "email is required"},
		"age":   {"required": "age is required"},
	}
}
