package smith

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
)

func ViewContainer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Fprintf(w, renderContainer(id))
}

func SaveContainer(w http.ResponseWriter, r *http.Request) {
	id := persistContainer(r)
	router := CreateRouter()
	newUrl, _ := router.Get("edit_container").URL("id", id)

	http.Redirect(w, r, newUrl.String(), http.StatusSeeOther)
}

func DeleteContainer(w http.ResponseWriter, r *http.Request) {
	router := CreateRouter()
	newUrl, _ := router.Get("edit_container").URL("id", mux.Vars(r)["id"])

	http.Redirect(w, r, newUrl.String(), http.StatusSeeOther)
}

func renderContainer(id string) string {
	loadTemplateConfig()
	loadTemplates()
	buf := new(bytes.Buffer)
	redisClient := ConnectToRedis()
	containerRepository := NewContainerRepository(redisClient)
	container, _ := containerRepository.FindByID(id)
	// Anonymous struct
	data := struct {
		Container *Container
	}{
		container,
	}
	renderTemplate(buf, "container.html", data)

	return buf.String()
}

func persistContainer(r *http.Request) string {
	redisClient := ConnectToRedis()
	container := new(Container)

	r.ParseForm()
	container.ID = r.FormValue("id")
	if container.ID == "" {
		id, _ := uuid.NewV4()
		container.ID = id.String()
	}
	container.Name = r.FormValue("name")
	container.Image = r.FormValue("image")
	container.CreatedAt = time.Now()
	container.Ports = r.FormValue("ports")

	containerRepository := NewContainerRepository(redisClient)
	containerRepository.Persist(container)

	return container.ID
}
