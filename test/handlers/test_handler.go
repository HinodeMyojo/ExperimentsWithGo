package handlers

import (
	"net/http"
	"test/services"

	"github.com/gin-gonic/gin"
)

type TestHandler struct {
	service services.TestServiceImpl
}

func (h *TestHandler) GetTest(c *gin.Context){
	result, err := h.service.GetTests(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return;
	}
	c.JSON(http.StatusOK, result)
}