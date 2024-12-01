package add

import (
	"WST_lab4_server/internal/models"
	"net/http"
)

type Request struct {
	models.Person `json:"person"`
}

type Response struct {
	Status string `json:"status"`
	Id     uint   `json:"id,omitempty"`
	Error  string `json:"error,omitempty"`
}

type PersonAdder interface {
	AddPersonHandler(request interface{}, w http.ResponseWriter, r *http.Request) (interface{}, error)
}
