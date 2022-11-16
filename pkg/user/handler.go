package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	errors "user-api/internal/errors"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

//monitoring "github.com/shipperizer/miniature-monkey/monitoring"

type Blueprint struct {
	service Service
	//monitor monitoring.MonitorInterface
}

// Routes exposes the handler on a route attached to the router
func (bp *Blueprint) Routes(router *mux.Router) {
	routes := router.PathPrefix("/user_API").Subrouter()

	routes.HandleFunc("/newUser", bp.PostHandler).Methods("POST")
}

// PostHandler handler for POST request.
func (bp Blueprint) PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Here 5")
	user := &User{}

	err := bp.service.Create(user)

	if err != nil {
		errors.ProcessError(w, err, http.StatusInternalServerError)
		fmt.Println("Here 6")
		return
	}

	log.Info("User Created")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}

// PutHandler handler for PUT request.
func (bp Blueprint) PutHandler(w http.ResponseWriter, r *http.Request) {

}

// DeleteHandler handler for DELETE request.
func (bp Blueprint) DeleteHandler(w http.ResponseWriter, r *http.Request) {
}

// ListHandler handler for GET request /user?<filters>
func (bp Blueprint) ListHandler(w http.ResponseWriter, r *http.Request) {
}

// func NewBlueprint(service Service, monitor monitoring.MonitorInterface) *Blueprint {
func NewBlueprint(service Service) *Blueprint {
	return &Blueprint{
		service: service,
		//	monitor: monitor,
	}
}
