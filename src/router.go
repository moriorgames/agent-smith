package smith

import "github.com/gorilla/mux"

func CreateRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", ViewHome)
	router.HandleFunc("/static/css/{file}", StaticCss)
	router.HandleFunc("/static/js/{file}", StaticJs)
	router.HandleFunc("/static/img/{file}", StaticImg)

	return router
}
