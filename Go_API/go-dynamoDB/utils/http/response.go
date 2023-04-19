package http

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status int         `json:"status"`
	Result interface{} `json:"result"`
}

func NewResponse(data interface{}, status int) *Response {
	return &Response{
		status,
		data,
	}
}
func (res *Response) Bytes() []byte {
	data, _ := json.Marshal(res)
	return data
}
func (res *Response) String() string {
	return string(res.Bytes())
}

func (res *Response) SendResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(res.Status)
	w.Write(res.Bytes())
	log.Println(res.String())
	// json.NewEncoder(w).Encode(res)

}

// 200
func StatusOk(w http.ResponseWriter, r *http.Request, data interface{}) {
	NewResponse(data, http.StatusOK).SendResponse(w, r)
}

// 204
func StatusNoContent(w http.ResponseWriter, r *http.Request) {
	NewResponse(nil, http.StatusNoContent).SendResponse(w, r)
}

// 400
func StatusBadRequest(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]any{"error": err.Error}
	NewResponse(data, http.StatusBadRequest).SendResponse(w, r)
}

// 404
func StatusNotFound(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{"error": err.Error}
	NewResponse(data, http.StatusNotFound).SendResponse(w, r)
}

// 405
func StatusMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	NewResponse(nil, http.StatusMethodNotAllowed).SendResponse(w, r)
}

//409

func StatusConflict(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{"error": err.Error}
	NewResponse(data, http.StatusConflict).SendResponse(w, r)

}

//500

func StatusInternamServerError(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{"error": err.Error}
	NewResponse(data, http.StatusInternalServerError).SendResponse(w, r)

}
