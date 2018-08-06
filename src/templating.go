package smith

import (
	"github.com/oxtoacart/bpool"
	"html/template"
	"path/filepath"
	"log"
	"fmt"
	"io"
)

type TemplateConfig struct {
	LayoutPath  string
	IncludePath string
}

var templates map[string]*template.Template
var bufferPool *bpool.BufferPool
var mainTemplate = `{{define "main"}}{{template "base" .}}{{end}}`
var templateConfig TemplateConfig

func loadTemplateConfig() {
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

	currentTemplate := template.New("main")

	currentTemplate, err = currentTemplate.Parse(mainTemplate)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range includeFiles {
		fileName := filepath.Base(file)
		files := append(layoutFiles, file)
		templates[fileName], err = currentTemplate.Clone()
		if err != nil {
			log.Fatal(err)
		}
		templates[fileName] = template.Must(templates[fileName].ParseFiles(files...))
	}

	log.Println("templates loading successful")

	bufferPool = bpool.NewBufferPool(64)
	log.Println("buffer allocation successful")
}

func renderTemplate(w io.Writer, name string, data interface{}) {
	tmpl, ok := templates[name]
	if !ok {
		fmt.Printf("The template %s does not exist.", name)
	}

	buf := bufferPool.Get()
	defer bufferPool.Put(buf)

	err := tmpl.Execute(buf, data)
	if err != nil {
		fmt.Printf(err.Error())
	}

	buf.WriteTo(w)
}
