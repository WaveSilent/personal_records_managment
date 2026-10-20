package http

import (
	"encoding/json"
	"fmt"
	model "managment/models"
	"managment/personal_records"
	"net/http"
	"uuid"

	"github.com/gorilla/mux"
)

type HTTPHandler struct {
	managment *personal_records.Managment
}

func NewHTTPHandler(managment *personal_records.Managment) *HTTPHandler {
	return &HTTPHandler{
		managment: managment,
	}
}

/*
pattern: /employees
method: POST
info: JSON in HTTP requst body

succeed:
  - status code: 201 Created
  - response_body: JSON represent created worker

failed:
  - status code: 500, 404, 400, ...
  - response body: JSON with error + time
*/
func (h *HTTPHandler) HandlerCreateWorker(w http.ResponseWriter, r *http.Request) {
	var workerDTO WorkerDTO
	if err := json.NewDecoder(r.Body).Decode(&workerDTO); err != nil {
		errDTO := NewErrorDTO(err)

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)

		return
	}

	worker := model.NewWorker(workerDTO.FullName, workerDTO.Position)

	if err := h.managment.AddWorker(worker); err != nil {
		errDTO := NewErrorDTO(err)

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)

		return
	}

	b, err := json.MarshalIndent(worker, "", "	")

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to response body:", err.Error())
		return
	}
}

/*
pattern: /employees
method: GET
info: -

succeed:
  - status code: 200 OK
  - response_body: JSON represent all workers

failed:
  - status code: 500, 404, 400, ...
  - response body: JSON with error + time
*/
func (h *HTTPHandler) HandlerGetAllWorkers(w http.ResponseWriter, r *http.Request) {
	err, workers := h.managment.GetAllWorkers()

	if err != nil {
		errDTO := NewErrorDTO(err)

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)

		return
	}

	b, err := json.MarshalIndent(workers, "", "	")

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to response body:", err.Error())
		return
	}
}

/*
pattern: /employees/{id}
method: DELETE
info: pattern

succeed:
  - status code: 204 No Content
  - response_body: -

failed:
  - status code: 500, 404, 400, ...
  - response body: JSON with error + time
*/
func (h *HTTPHandler) HandlerDeleteWorker(w http.ResponseWriter, r *http.Request) {
	strID := mux.Vars(r)["id"]

	id, err := uuid.Parse(strID)
	if err != nil {
		errDTO := NewErrorDTO(err)

		http.Error(w, errDTO.ToString(), http.StatusConflict)

		return
	}

	if err := h.managment.DeleteWorker(id); err != nil {
		errDTO := NewErrorDTO(err)

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
