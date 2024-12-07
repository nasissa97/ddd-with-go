package fundamentals

import (
	"encoding/json"
	"net/http"

	"github.com/nasissa97/ddd-with-go/fundamentals/oapi"
)

var _ oapi.ServerInterface = (*Server)(nil)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	resp := []User{
		{
			"abc",
			PaymentDetails{"1234"},
		},
		{
			"def",
			PaymentDetails{"5678"},
		},
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
