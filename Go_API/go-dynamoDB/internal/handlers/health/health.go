package health

import (
	"net/http"

	"github.com/arjun/modules/go-dynamoDB/internal/handlers"
	"github.com/arjun/modules/go-dynamoDB/internal/repository/adapter"
	HttpStatus "github.com/arjun/modules/go-dynamoDB/utils/http"
)

type Handler struct {
	handlers.Interface
	Repository adapter.Interface
}

func NewHandler(respository adapter.Interface) handlers.Interface {
	return &Handler{
		Repository: respository,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	if !h.Repository.Health() {
		HttpStatus.StatusInternalServerError(w,r errors.New("Relation database is not started"))
	}
	HttpStatus.StatusOk(w, r, "service ok")
}
func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {

}
