package controllers

import (
	"github.com/gin-gonic/gin"
	lg "github.com/hiromaily/golibs/log"
	"net/http"
)

// GetIndex is top page
func GetIndexAction(c *gin.Context) {
	lg.Info("GetIndexAction()")
	c.HTML(http.StatusOK, "index", nil)
}
