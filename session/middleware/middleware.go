package middleware

import (
	mgr "github.com/JosephChan007/go-Gin/session/manager"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func SessionMiddleware(m *mgr.SessionManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		var sd *mgr.SessionData
		sid, err := c.Cookie(mgr.SessionCookieName)
		if err != nil {
			sd = m.CreateSession()
			sid = sd.Id
		} else {
			sd, err = m.GetSessionData(sid)
			if err != nil {
				sd = m.CreateSession()
				sid = sd.Id
			}
		}

		c.Set(mgr.SessionContextName, sd)
		log.Printf("[Session]session data is: %#v", sd)
		c.SetCookie(mgr.SessionCookieName, sid, 30, "/", "hdfs-host3", false, true)
		c.Next()
	}
}

func AuthMiddleware(c *gin.Context) {
	sdobj, ok := c.Get(mgr.SessionContextName)
	if !ok {
		c.Redirect(http.StatusFound, "/session/login")
		return
	}
	sd := sdobj.(*mgr.SessionData)
	log.Printf("[Auth]session data is: %#v", sd)
	isLoginObj, err := sd.Get("isLogin")
	if err != nil {
		c.Redirect(http.StatusFound, "/session/login")
		return
	}
	isLogin := isLoginObj.(bool)
	if !isLogin {
		c.Redirect(http.StatusFound, "/session/login")
		return
	}

	c.Next()
}
