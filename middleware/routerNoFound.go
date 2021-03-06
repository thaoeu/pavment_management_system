package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/izghua/zgh"
)

func CheckExist() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		s := strings.Contains(path, "/backend/")
		c.Next()
		status := c.Writer.Status()
		if status == 404 {
			zgh.ZLog().Info("message", "routerNoFound CheckExist", "Path", path)
			if s {
				c.Redirect(http.StatusMovedPermanently, "/backend/")
			} else {
				c.Redirect(http.StatusMovedPermanently, "/404")
			}
		}
	}
}
