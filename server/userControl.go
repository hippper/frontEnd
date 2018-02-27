package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/hippper/frontEnd/utils"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserId      int64
	UserName    string
	Password    string
	PasswrodMd5 string
	Expired     int64
}

const (
	Cookie_auth_key       = "auth_token"
	Cookie_expired_second = 20 * 60
)

var userList = []User{
	User{UserId: 1, UserName: "admin", Password: "qazwsx"},
	User{UserId: 2, UserName: "user1", Password: "123456"},
}

var authMap = make(map[string]*User)

func generateSessionToken() string {

	timeStr := strconv.FormatInt(time.Now().UnixNano(), 10)
	return utils.Md5(timeStr)
}

func isUserValid(username, userpasswd string) bool {

	for _, u := range userList {
		if u.UserName == username && u.Password == userpasswd {
			return true
		}
	}
	return false
}

func getUserInfo(token string) *User {
	user := &User{}

	userInfo, ok := authMap[token]
	if ok {
		user = userInfo
	}

	return user
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {

		if auth_cookie, err := c.Request.Cookie(Cookie_auth_key); err == nil {
			value := auth_cookie.Value
			if userInfo, ok := authMap[value]; ok {
				if userInfo.Expired > utils.NowInS() {
					authMap[value].Expired = utils.NowInS() + Cookie_expired_second
					c.Next()
					return
				}
				delete(authMap, value)
			}
		}
		c.Redirect(http.StatusFound, "/login.html")
		c.Abort()
		return
	}
}

func (s *WebServer) UserLogin(c *gin.Context) {

	username := c.PostForm("username")
	userpasswd := c.PostForm("userpasswd")

	if isUserValid(username, userpasswd) {
		token := generateSessionToken()

		tmpS := &User{
			UserName:    username,
			Password:    userpasswd,
			PasswrodMd5: utils.Md5(userpasswd),
			Expired:     utils.NowInS() + Cookie_expired_second,
		}
		authMap[token] = tmpS
		c.SetCookie(Cookie_auth_key, token, 0, "", "", false, true)

		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    99,
		"message": "login failed",
	})

}

func (s *WebServer) UserLogout(c *gin.Context) {

	auth_cookie, err := c.Request.Cookie(Cookie_auth_key)
	if err == nil {
		delete(authMap, auth_cookie.Value)
	}

	c.SetCookie(Cookie_auth_key, "", -1, "", "", false, true)

	c.Redirect(http.StatusFound, "/login.html")
	return
}
