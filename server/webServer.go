package server

import (
	"net/http"
	"regexp"

	"github.com/hippper/frontEnd/define"

	"github.com/gin-gonic/gin"
)

type WebServer struct {
}

func (s *WebServer) StartServer() error {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/index.html")
		return
	})
	router.Static("/js", "./static/js")
	router.Static("/img", "./static/img")
	router.Static("/css", "./static/css")
	router.Static("/fonts", "./static/fonts")

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	router.GET("/"+define.HTML_INDEX, s.HtmlFiles)

	return router.Run("0.0.0.0:8088")
}

func (s *WebServer) HtmlFiles(c *gin.Context) {
	pattern := `\w+.html`
	r, _ := regexp.Compile(pattern)

	path := c.Request.RequestURI
	file := r.FindString(path)

	switch file {
	case define.HTML_INDEX:
		c.HTML(http.StatusOK, file, nil)
	}
}

func NewServer() *WebServer {
	s := new(WebServer)
	return s
}
