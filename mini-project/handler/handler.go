package handler

import (
	"mini-project/models"
	"mini-project/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	storage *storage.Storage
}

func NewHandler(strg *storage.Storage) *Handler {
	return &Handler{
		storage: strg,
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var body models.User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	user, err := h.storage.CreateUser(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *Handler) CreateTask(c *gin.Context) {
	var body models.Task
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	task, err := h.storage.CreateTask(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, task)
}
