package httpserver

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func NewServer() {

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This Hello World is under / "))
	})

	r.Route("/api/v1/health", func(r chi.Router) {
		r.Get("/", healthcheck)
	})

	r.Route("/api/v1/users", func(r chi.Router) {
		r.Get("/", getUser)
	})

	fmt.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")

}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")

}
