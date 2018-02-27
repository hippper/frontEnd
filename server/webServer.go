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

	router.GET("/login.html", s.HtmlFiles)
	router.POST("/userlogin", s.UserLogin)
	router.GET("/userlogout", s.UserLogout)

	router.Use(AuthMiddleWare())

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
		UserName := "None"
		if auth_cookie, err := c.Request.Cookie(Cookie_auth_key); err == nil {
			if userInfo, ok := authMap[auth_cookie.Value]; ok {
				UserName = userInfo.UserName
			}
		}
		c.HTML(http.StatusOK, file, gin.H{
			"UserName": UserName,
		})
	default:
		c.HTML(http.StatusOK, file, nil)
	}
}

func NewServer() *WebServer {
	s := new(WebServer)
	return s
}
