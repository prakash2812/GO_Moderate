package product

import (
	"errors"
	"net/http"
	"time"

	HttpStatus "github.com/arjun/modules/go-dynamoDB/utils/http"
	"github.com/go-chi/chi"
	"github.com/google/uuid"

	"github.com/arjun/modules/go-dynamoDB/internal/controllers/product"
	EntityProduct "github.com/arjun/modules/go-dynamoDB/internal/entities/product"
	"github.com/arjun/modules/go-dynamoDB/internal/handlers"
	"github.com/arjun/modules/go-dynamoDB/internal/repository/adapter"
	Rules "github.com/arjun/modules/go-dynamoDB/internal/rules"
	RulesProduct "github.com/arjun/modules/go-dynamoDB/internal/rules/product"
)

type Handler struct {
	handlers.Interface
	Controller product.Interface
	Rules      Rules.Interface
}

func NewHandler(repository adapter.Interface) handlers.Interface {
	return &Handler{
		Controller: product.NewController(repository),
		Rules:      RulesProduct.NewRules(),
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	if chi.URLParam(r, "ID") != "" {
		h.getOne(w, r)
	} else {
		h.getAll(w, r)
	}
}

func (h *Handler) getOne(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, errors.New("ID is not valid UUID"))
		return
	}
	response, err := h.Controller.ListOne(ID)
	if err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}
	HttpStatus.StatusOk(w, r, response)

}
func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	response, err := h.Controller.ListAll()
	if err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}
	HttpStatus.StatusOk(w, r, response)
}
func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	productBody, err := getBodyAndValidate(r, uuid.Nil)
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, err)
		return
	}
	ID, err := h.Controller.Create(productBody)
	if err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}
	HttpStatus.StatusOk(w, r, map[string]interface{}{"id": ID.String()})
}
func (h *Handler) getBodyAndValidate(r *http.Request, id uuid.UUID) (*EntityProduct.Product, error) {
	productBody := &EntityProduct.product{}
	body, err := h.Rules.ConvertIoReaderToStruct(r.Body, productBody)
	if err != nil {
		return &EntityProduct.product{}, errors.New("body is required")
	}
	productParsed, err := EntityProduct.InterfaceToModel(body)
	if err != nil {
		return &EntityProduct.product{}, errors.New("error on converting body to model")
	}
	setDefaultValues(productParsed, id)
	return productParsed, h.Rule.validate(productParsed)
	// return
}
func setDefaultValues(product *EntityProduct.Product, ID uuid.UUID) {
	product.UpdateAt = time.Now()
	if ID == uuid.Nil {
		product.ID = uuid.New()
		product.cretedAt = time.Now()
	} else {
		product.ID = ID
	}
}
func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, errors.New("Invalid ID"))
		return
	}
	productBody, err := h.getBodyAndValidate(r, ID)
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, err)
		return
	}
	if err := h.Controller.Update(ID, productBody); err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}
	HttpStatus.StatusNoContent(w, r)
}
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, errors.New("Invalid ID"))
		return
	}
	if err := h.Controller.Remove(ID); err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}
	HttpStatus.StatusNoContent(w, r)

}
func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusNoContent(w, r)
}
