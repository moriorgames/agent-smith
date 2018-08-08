package smith

import "github.com/gorilla/mux"

func CreateRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", ViewHome).
		Methods("GET").Name("app_home")
	router.HandleFunc("/container", ViewContainer).
		Methods("GET").Name("app_new_container")
	router.HandleFunc("/container/{id}", ViewContainer).
		Methods("GET").Name("app_edit_container")
	router.HandleFunc("/container/{id}", UpdateContainer).
		Methods("POST").Name("app_update_container")
	router.HandleFunc("/static/css/{file}", StaticCss).
		Methods("GET")
	router.HandleFunc("/static/js/{file}", StaticJs).
		Methods("GET")
	router.HandleFunc("/static/img/{file}", StaticImg).
		Methods("GET")

	return router
}
