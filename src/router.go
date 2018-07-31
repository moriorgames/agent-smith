package smith

import "github.com/gorilla/mux"

func createRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)

	return router
}
