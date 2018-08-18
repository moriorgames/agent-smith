package smith

import (
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
)

const CssPath string = "../static/css/"
const JsPath string = "../static/js/"
const ImgPath string = "../static/img/"

func StaticCss(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["file"]
	data, err := ioutil.ReadFile(string(CssPath + filename))

	if err == nil {
		w.Header().Set("Content-Type", "text/css")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 CSS File " + http.StatusText(http.StatusNotFound)))
	}
}

func StaticJs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["file"]
	data, err := ioutil.ReadFile(string(JsPath + filename))

	if err == nil {
		w.Header().Set("Content-Type", "application/javascript")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 JS File " + http.StatusText(http.StatusNotFound)))
	}
}

func StaticImg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["file"]
	data, err := ioutil.ReadFile(string(ImgPath + filename))

	if err == nil {
		w.Header().Set("Content-Type", "image/jpg")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 JPG File " + http.StatusText(http.StatusNotFound)))
	}
}
