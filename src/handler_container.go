package smith

import (
	"fmt"
	"bytes"
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"log"
)

func renderContainerContent(id string) string {
	loadTemplateConfig()
	loadTemplates()
	buf := new(bytes.Buffer)
	renderTemplate(buf, "container.html", nil)

	return buf.String()
}

func ViewContainer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Fprintf(w, renderContainerContent(id))
}

func UpdateContainer(w http.ResponseWriter, r *http.Request) {
	router := CreateRouter()
	newUrl, _ := router.Get("app_edit_container").URL("id", mux.Vars(r)["id"])

	saveContainer(r)

	http.Redirect(w, r, newUrl.String(), http.StatusSeeOther)
}

func saveContainer(r *http.Request) {
	redisClient := ConnectToRedis()
	container := new(Container)
	createdAt := time.Now()

	r.ParseForm()
	container.ID = r.FormValue("id")
	container.Name = r.FormValue("name")
	container.Ip = r.FormValue("ip")
	container.CreatedAt = createdAt
	container.Ports = r.FormValue("ports")
	container.Status = true

	containerRepository := NewContainerRepository(redisClient)
	err := containerRepository.Persist(container)
	if err != nil {
		log.Fatal(err)
	}
}
