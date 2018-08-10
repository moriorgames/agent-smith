package smith

import (
	"fmt"
	"bytes"
	"net/http"
	"github.com/gorilla/mux"
	"time"
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
	renderTemplate(buf, "container.html", nil)

	return buf.String()
}

func persistContainer(r *http.Request) string {
	redisClient := ConnectToRedis()
	container := new(Container)
	createdAt := time.Now()

	r.ParseForm()
	container.ID = r.FormValue("id")
	if container.ID == "" {
		id, _ := uuid.NewV4()
		container.ID = id.String()
	}
	container.Name = r.FormValue("name")
	container.Ip = r.FormValue("ip")
	container.CreatedAt = createdAt
	container.Ports = r.FormValue("ports")
	container.Status = true

	containerRepository := NewContainerRepository(redisClient)
	containerRepository.Persist(container)

	return container.ID
}
