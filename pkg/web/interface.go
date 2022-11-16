package web

import (
	"net/http"

	bp "user-api/pkg/user"

	"github.com/gorilla/mux"
)

// APIInterface
type APIInterface interface {
	Router() *mux.Router
	Handler() http.Handler
	RegisterBlueprints(r *mux.Router, blueprints ...bp.BlueprintInterface)
	UseMiddlewares(r *mux.Router, mwf ...mux.MiddlewareFunc)
}
