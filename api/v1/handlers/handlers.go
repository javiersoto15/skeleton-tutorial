package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Routes( /* any dependency injection comes here*/ ) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/:exampleId", ClosureHandlerExample())
	r.Post("/item", RegularHandlerExample)
	r.Delete("/item/:exampleId", nil)
	return r
}

//ClosureHandlerExample shows different ways you can pass dependencies into your handler
func ClosureHandlerExample() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		//handler logic
	}
}

//RegularHandlerExample implements a simple handler to use the RW and Request object
func RegularHandlerExample(rw http.ResponseWriter, r *http.Request) {
	//handler logic
}
