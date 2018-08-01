package smith

import (
	"github.com/oxtoacart/bpool"
	"html/template"
	"path/filepath"
	"log"
	"fmt"
	"bytes"
	"io"
	"net/http"
)

type TemplateConfig struct {
	LayoutPath  string
	IncludePath string
}

var templates map[string]*template.Template
var bufpool *bpool.BufferPool
var mainTmpl = `{{define "main"}}{{template "base" .}}{{end}}`
var templateConfig TemplateConfig

func loadConfiguration() {
	templateConfig.LayoutPath = "../templates/layout/"
	templateConfig.IncludePath = "../templates/"
}

func loadTemplates() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	layoutFiles, err := filepath.Glob(templateConfig.LayoutPath + "*.html")
	if err != nil {
		log.Fatal(err)
	}

	includeFiles, err := filepath.Glob(templateConfig.IncludePath + "*.html")
	if err != nil {
		log.Fatal(err)
	}

	mainTemplate := template.New("main")

	mainTemplate, err = mainTemplate.Parse(mainTmpl)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range includeFiles {
		fileName := filepath.Base(file)
		files := append(layoutFiles, file)
		templates[fileName], err = mainTemplate.Clone()
		if err != nil {
			log.Fatal(err)
		}
		templates[fileName] = template.Must(templates[fileName].ParseFiles(files...))
	}

	log.Println("templates loading successful")

	bufpool = bpool.NewBufferPool(64)
	log.Println("buffer allocation successful")
}

func renderTemplate(w io.Writer, name string, data interface{}) {
	tmpl, ok := templates[name]
	if !ok {
		fmt.Printf("The template %s does not exist.", name)
	}

	buf := bufpool.Get()
	defer bufpool.Put(buf)

	err := tmpl.Execute(buf, data)
	if err != nil {
		fmt.Printf(err.Error())
	}

	buf.WriteTo(w)
}

func renderHomeContent() string {

	loadConfiguration()
	loadTemplates()

	buf := new(bytes.Buffer)
	renderTemplate(buf, "home.html", nil)

	return buf.String()
}

func ViewHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, renderHomeContent())
}
