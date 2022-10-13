package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) GetAll(c *gin.Context) {
	data, err := r.services.Tasks.GetAll(1)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}
