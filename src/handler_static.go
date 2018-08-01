package smith

import (
	"github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
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
		w.WriteHeader(200)
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 CSS File " + http.StatusText(404)))
	}
}

func StaticJs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["file"]
	data, err := ioutil.ReadFile(string(JsPath + filename))

	if err == nil {
		w.Header().Set("Content-Type", "application/javascript")
		w.WriteHeader(200)
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 JS File " + http.StatusText(404)))
	}
}

func StaticImg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["file"]
	data, err := ioutil.ReadFile(string(ImgPath + filename))

	if err == nil {
		w.Header().Set("Content-Type", "image/jpg")
		w.WriteHeader(200)
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 JPG File " + http.StatusText(404)))
	}
}
