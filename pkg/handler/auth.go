package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	fmt.Println("Hello")
}

func (h *Handler) SignIn(c *gin.Context) {
	fmt.Println("Hello")
}
