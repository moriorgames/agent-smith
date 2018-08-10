package smith

import "github.com/gorilla/mux"

func CreateRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", ViewHome).Methods("GET").Name("home")
	router.HandleFunc("/container", ViewContainer).Methods("GET").Name("new_container")
	router.HandleFunc("/container/{id}", ViewContainer).Methods("GET").Name("edit_container")
	router.HandleFunc("/container", SaveContainer).Methods("POST").Name("save_new_container")
	router.HandleFunc("/container/{id}", SaveContainer).Methods("POST").Name("save_container")
	router.HandleFunc("/container/{id}", DeleteContainer).Methods("DELETE").Name("del_container")
	router.HandleFunc("/static/css/{file}", StaticCss).Methods("GET")
	router.HandleFunc("/static/js/{file}", StaticJs).Methods("GET")
	router.HandleFunc("/static/img/{file}", StaticImg).Methods("GET")

	return router
}
