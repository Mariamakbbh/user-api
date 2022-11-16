package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type BlueprintInterface interface {
	Routes(router *mux.Router)
	PostHandler(w http.ResponseWriter, r *http.Request)
	PutHandler(w http.ResponseWriter, r *http.Request)
	DeleteHandler(w http.ResponseWriter, r *http.Request)
	ListHandler(w http.ResponseWriter, r *http.Request)
}

type Service interface {
	Create(user *User) error
	Update(user *User) error
	Delete(id string) error
	List() ([]User, error)
}

type Repo interface {
	Create(user *User) error
	Update(user *User) error
	Delete(id string) error
	List() ([]User, error)
}
