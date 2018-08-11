// link table functions
package links

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	auth "gorestapi/middleware"
)

// add links functions to the mux router
func AddToRouter(r *mux.Router) {
	r.HandleFunc("/api/links", auth.AuthMiddleware(GetLinks)).Methods("GET")
	r.HandleFunc("/api/link/{id}", auth.AuthMiddleware(GetLink)).Methods("GET")
	r.HandleFunc("/api/link", auth.AuthMiddleware(CreateLink)).Methods("POST")
	r.HandleFunc("/api/link/{id}", auth.AuthMiddleware(UpdateLink)).Methods("PUT")
	r.HandleFunc("/api/link/{id}", auth.AuthMiddleware(DeleteLink)).Methods("DELETE")
}

func GetLinks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetLinks called")
}

func GetLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetLink called")
}

func CreateLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CreateLink called")
}

func UpdateLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UpdateLink called")
}

func DeleteLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DeleteMededeling called")
}
