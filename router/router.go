package router

import (
	"html/template"
	"log"
	"net/http"
	"vgserver/view"

	"github.com/gin-gonic/gin"
)

// New creates the router
func New() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Load templates
	r.SetHTMLTemplate(template.Must(template.New("index.html").Parse(view.IndexTemplate)))
	log.Println("Loaded index.html template file")

	r.GET("/", index)

	return r
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
