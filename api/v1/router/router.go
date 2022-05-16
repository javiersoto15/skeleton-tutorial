package router

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/javiersoto15/skeleton-tutorial/api/v1/handlers"
)

func Initialize() *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		render.SetContentType(render.ContentTypeJSON), //forces Content-type
		middleware.RedirectSlashes,
		middleware.Recoverer,            //middleware to recover from panics
		middleware.Heartbeat("/health"), //for heartbeat process such as Kubernetes liveprobeness
		cors.Handler(cors.Options{
			// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
			AllowedOrigins: []string{"https://*", "http://*"},
			// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}),
	)

	//Sets context for all requests
	router.Use(middleware.Timeout(30 * time.Second))

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/", handlers.Routes()) //Implementation of routes from handlers.go
		r.Mount("/metrics", nil)        //for monitoring agents such as prometheus
	})

	return router
}

func ServeRouter() {
	r := Initialize()

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Error serving router")
	}
}
