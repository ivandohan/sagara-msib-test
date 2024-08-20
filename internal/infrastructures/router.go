package infrastructures

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func (s *Server) InitiateRouter() (r *mux.Router) {
	r = mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	r.HandleFunc("", defaultHandler).Methods("GET")
	r.HandleFunc("/", defaultHandler).Methods("GET")

	sub := r.PathPrefix("/inventory-api").Subrouter()

	sub.HandleFunc("/baju", s.BajuDelivery.HandleClient).Methods("GET", "POST", "PUT", "DELETE")

	return r
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Service API Inventory Baju"))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(r)
}
